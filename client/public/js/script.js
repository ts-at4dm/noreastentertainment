function updateStates() {
    const countrySelect = document.getElementById('country');
    const stateSelect = document.getElementById('state');
    const selectedCountry = countrySelect.value;

    stateSelect.innerHTML = ''; // Clear existing options

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

        usStates.forEach(state => {
            const option = document.createElement('option');
            option.value = state;
            option.text = state;
            stateSelect.appendChild(option);
        });
    } else if (selectedCountry === 'Canada') {
        const canadianProvinces = [
            "Alberta", "British Columbia", "Manitoba", "New Brunswick", 
            "Newfoundland and Labrador", "Nova Scotia", "Ontario", "Prince Edward Island", 
            "Quebec", "Saskatchewan"
        ];

        canadianProvinces.forEach(province => {
            const option = document.createElement('option');
            option.value = province;
            option.text = province;
            stateSelect.appendChild(option);
        });
    } 
}
function showSidebar(){
    const sidebar = document.querySelector('.sidebar');
    sidebar.style.display = 'flex';
}
function hideSidebar(){
    const sidebar = document.querySelector('.sidebar');
    sidebar.style.display = 'none';
}
function toggleDropdown(event) {
    event.preventDefault(); // Prevent link from redirecting
    const dropdown = event.target.closest('.dropdown');
    dropdown.classList.toggle('active');
}
