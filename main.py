from flask import Flask, render_template, request, redirect, url_for, flash
from flask_mail import Mail, Message
from datetime import datetime

app = Flask(__name__) # Init a flask web application instance

# Implementes mail api
app.config['MAIL_SERVER']='NA'
app.config['MAIL_PORT'] = "NA"
app.config['MAIL_USERNAME'] = 'NA'
app.config['MAIL_PASSWORD'] = 'NA'
app.config['MAIL_USE_TLS'] = "NA"
app.config['MAIL_USE_SSL'] = "NA"
app.config['SECRET_KEY'] = "NA"

mail = Mail(app)

@app.route('/')
def index():
    return render_template('contact.html')

# inits the form from the HTML
@app.route('/send_email', methods=['POST'])
def send_email():
    if request.method == 'POST':
        fname = request.form['First Name']
        lname = request.form['Last Name']
        email = request.form['Email']
        phone = request.form['Phone']
        date_time = request.form["D_T"]
        event_type = request.form["Event Type"]
        services = request.form.getlist('Services[]')
        message = request.form['Message']
        
        formatted_date_time = datetime.strptime(date_time, '%Y-%m-%dT%H:%M')  # Convert to datetime object
        formatted_date_time = formatted_date_time.strftime('%m-%d-%Y %H:%M')  # Format it to MM-DD-YYYY HH:MM
        
        msg = Message('Contact Form Submission', sender='noreast@info.com',
                      recipients=['noreastmaine@gmail.com'])
        # Prints message to email
        msg.body = f"Name: {fname} {lname}\nEmail: {email}\nPhone: {phone}\nDate & Time: {formatted_date_time}\nEvent: {event_type}\nServices: {', '.join(services)}\nMessage: {message}"
        mail.send(msg)
        flash('Message Sent Successfully', 'success')
        
    return redirect(url_for('response'))

# Forwards the user to the response page. 
@app.route('/response')
def response():
    return render_template('response.html')

app.run(debug=True)