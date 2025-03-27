// Function to update the state/province dropdown based on the selected country
function updateStates() {
    const countrySelect = document.getElementById('country'); // Get the country dropdown
    const stateSelect = document.getElementById('state'); // Get the state/province dropdown
    const selectedCountry = countrySelect.value; // Get the selected country

    stateSelect.innerHTML = ''; // Clear existing options in the state dropdown

    // If the selected country is the US, populate the dropdown with US states
    if (selectedCountry === 'US') {
        const usStates = [
            "Alabama", "Alaska", "Arizona", "Arkansas", "California", 
            "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", 
            "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", 
            "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", 
            "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", 
            "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", 
            "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", 
            "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", 
            "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", 
            "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"
        ];

        // Loop through US states and create option elements for the dropdown
        usStates.forEach(state => {
            const option = document.createElement('option'); // Create a new option element
            option.value = state; // Set the option value
            option.text = state; // Set the display text
            stateSelect.appendChild(option); // Add the option to the dropdown
        });
    } 
    // If the selected country is Canada, populate the dropdown with Canadian provinces
    else if (selectedCountry === 'Canada') {
        const canadianProvinces = [
            "Alberta", "British Columbia", "Manitoba", "New Brunswick", 
            "Newfoundland and Labrador", "Nova Scotia", "Ontario", "Prince Edward Island", 
            "Quebec", "Saskatchewan"
        ];

        // Loop through Canadian provinces and create option elements for the dropdown
        canadianProvinces.forEach(province => {
            const option = document.createElement('option'); // Create a new option element
            option.value = province; // Set the option value
            option.text = province; // Set the display text
            stateSelect.appendChild(option); // Add the option to the dropdown
        });
    } 
}

// Function to show the sidebar
function showSidebar() {
    const sidebar = document.querySelector('.sidebar'); // Get the sidebar element
    sidebar.style.display = 'flex'; // Set display to flex (make it visible)
}

// Function to hide the sidebar
function hideSidebar() {
    const sidebar = document.querySelector('.sidebar'); // Get the sidebar element
    sidebar.style.display = 'none'; // Hide the sidebar
}

// Wait until the DOM is fully loaded before executing scripts
document.addEventListener("DOMContentLoaded", function() {
    // Define a function to toggle the dropdown menu visibility
    window.toggleDropdown = function() {
        document.getElementById("myDropdown").classList.toggle("show"); // Toggle the 'show' class
    };

    // Close the dropdown if the user clicks outside of it
    window.onclick = function(event) {
        if (!event.target.matches('.dropbtn')) { // Check if the clicked element is NOT the dropdown button
            var dropdowns = document.getElementsByClassName("dropdown-content"); // Get all dropdown elements
            for (var i = 0; i < dropdowns.length; i++) { // Loop through all dropdowns
                var openDropdown = dropdowns[i]; 
                if (openDropdown.classList.contains('show')) { // If the dropdown is open, close it
                    openDropdown.classList.remove('show');
                }
            }
        }
    };
});

