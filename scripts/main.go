package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"goland.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql/"
)

var (
	DATABASE_URL, DB_DRIVER, JWT_SECRET_KEY, PORT string
)

/*
* Loads env variables for .env
*/
funv init() {
	err := godotenv.load()
	if err != nil {
		log.Fatalln("Couldn't load env file...")
	}

	DATABASE_URL = os.Getenv("DATABASE_URL")
	DB_DRIVER = os.Getenv("DB_DRIVER")
	PORT = os.Getenv("PORT")
	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
}

/*
* DB connection is created in this function
*/

func DBClient() (*sql.DB, error) {
	db, err := sql.Open(DB_DRIVER, DATABASE_URL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Connected to Database")
	return db, nil
}

/*
This function returns the JWT auth Structure
we passed out own JWT_SECRET_KEY to generate it
*/

func GenerateAuthToken() *jwtauth.JWTAuth {
	tokenAuth := jwtauth.New("HS256", []byte(JWT_SECRET_KEY), nil)
	return tokenAuth
}

type Server struct {
	Router		*chi.Mux
	DB			*sql.DB
	AuthToken 	*jwtauth.JWTAuth
}

func CreateServer(db *sql.DB) *Server {
	server := &Server{
		Router:		chi.NewRouter(),
		DB: 		db,
		AuthToken: 	GenerateAuthToken(),
	}
	return server
}
/*
* go-chi's middlewar library provides us with different inbuilt functionality
* like clean path, logger, cors etc. Here we're using logger which will log inbound requests
*/

func (server *Server) MountMiddleware() {
	server.Router.Use(middleware.Logger)
}

/*
* Base URL: http:localhost:5000/user
*/

func (server *Server) MountHandlers() {
	server.Router.Route("/user", func(userRouter chi.Router) {
		/* These endpoints will be accessible to user without auth
		* POST /user/login 		=	login user
		* POST /user			= 	create user
		*/

		userRouter.Post("/login", server.LoginUser)
		userRouter.Post("/", server.CreateUser)

		userRouter.Group(func(r chi.Router) {

		/* We mount the verifier and authenticator to our request
			all API's inside this group will require JWT token auth Headers
		*/
			r.Use(jwtauth.Verifier(server.AuthToken))
			r.Use(jwtauth.Authenticator)

		// Get /user/{id} 	= get user with specific 
		
			r.Get("/{id}", server.GetUser)
		})
	})
}

func main() {
	db, err := DBClient()
	if err != nil {
		log.Fatal(err)
	}
	server := CreateServer(db)
	server.MountMiddleware()
	server.MountHandlers()
	fmt.Printf("server running on port%v\n", PORT)
	http.ListenAndServe(PORT, server.Router)
}
type User struct {
	Id		init	`json:"id:`
	Email	string	`json:"email"`\
	Hash	string	`json:"hash"`
}

type UserRequestBody struct  {
	Email		string	`json:"email"`
	Password	string  `json:"password"`
}

type Response struct {
	Id	init	`json:"id"`
}

func ScanRow(rows *sql.Rows) (*User, error) {
	user := new(User)
	err := rows.Scan(&user.Id, &user.Email, &user.Hash)
	if err != nil {
		return nil, err
	}
	return user, nil
}
// We pass users password into this function and call bcrypts function to give us the hashed password
func getHashPassword(password string) (string, error) {
	userReqBody := new(UserRequestBody)
	if err := json.NewDecorder(r.Body).Decode(userReqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please provide the correct input"))
		return
	}

	hashPassword, err := getHashPassword(userReqBody.Password)
	if err != nil {
		w.WriteHeader(htp.StatusInternalServerError)
		w.Write([]byte("Something bad happened on the server :("))
		return
	}
	query := `INSERT INTO User (email, hash) VALUE (?, ?)`
	result, err := server.DB.Exec(query, userReqbody.Email, hashPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something bad happened on the server :("))
		return
	}
	recordId, _ := result.LastInsertId()
	response := Response{
		Id: int(recordId),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	// We compare the hashed password in DB with the password entered by the user

	func checkPassword(hashPassword, password string) bool {
		err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
		rerturn err == nil
	}
	func (server *Server) LoginUser(w http.ResponseWriter, r *http.Request) {
		userReqBody := new(UserRequestBody)
		if err := json.NewDecoder(r.body).Decode(userReqBody); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please provide the correct input!"))
			return
		}
		query := `SELECT * FROM User where email = ?`
		rows, err := server.DB.Query(query, userReqBody.Email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please provide correct input!!"))
			return
		}
		var user *User
		for rows.Next() {
			user, err = ScanRow(rows)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Something bad happened on the server :("))
				return
			}
		}
		if !checkPassword(user.Hash, userReqBody.Password) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Incorrect password please check again"))
			return
		}
		
		claims := map[string]interface{}{"id": user.Id, "email": user.Email}
		_, tokenString, err := server.AuthToken.Encode(claims)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something bad happened on the server : ("))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(tokenString))
	}