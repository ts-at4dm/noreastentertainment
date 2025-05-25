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

// Testimonial

class Testimonial3DSlider {
            constructor() {
                this.currentSlide = 0;
                this.slides = document.querySelectorAll('.testimonial-slide');
                this.totalSlides = this.slides.length;
                this.indicators = document.querySelectorAll('.indicator');
                this.prevBtn = document.getElementById('prevBtn');
                this.nextBtn = document.getElementById('nextBtn');
                this.playPauseBtn = document.getElementById('playPauseBtn');
                this.isPlaying = true;
                this.autoPlayInterval = null;
                
                this.init();
            }

            init() {
                this.bindEvents();
                this.startAutoPlay();
                this.updateSlidePositions();
            }

            bindEvents() {
                this.prevBtn.addEventListener('click', () => this.prevSlide());
                this.nextBtn.addEventListener('click', () => this.nextSlide());
                this.playPauseBtn.addEventListener('click', () => this.toggleAutoPlay());
                
                this.indicators.forEach((indicator, index) => {
                    indicator.addEventListener('click', () => this.goToSlide(index));
                });

                // Touch/swipe support
                let startX = 0;
                let endX = 0;
                const slider = document.getElementById('slider3D');
                
                slider.addEventListener('touchstart', (e) => {
                    startX = e.touches[0].clientX;
                });
                
                slider.addEventListener('touchend', (e) => {
                    endX = e.changedTouches[0].clientX;
                    this.handleSwipe();
                });

                // Mouse drag support
                let isDragging = false;
                slider.addEventListener('mousedown', (e) => {
                    isDragging = true;
                    startX = e.clientX;
                });
                
                slider.addEventListener('mousemove', (e) => {
                    if (!isDragging) return;
                    e.preventDefault();
                });
                
                slider.addEventListener('mouseup', (e) => {
                    if (!isDragging) return;
                    isDragging = false;
                    endX = e.clientX;
                    this.handleSwipe();
                });

                // Pause auto-play on hover
                const section = document.querySelector('.testimonial-section');
                section.addEventListener('mouseenter', () => this.pauseAutoPlay());
                section.addEventListener('mouseleave', () => {
                    if (this.isPlaying) this.startAutoPlay();
                });
            }

            handleSwipe() {
                const swipeThreshold = 50;
                const diff = startX - endX;
                
                if (Math.abs(diff) > swipeThreshold) {
                    if (diff > 0) {
                        this.nextSlide();
                    } else {
                        this.prevSlide();
                    }
                }
            }

            updateSlidePositions() {
                this.slides.forEach((slide, index) => {
                    slide.classList.remove('active', 'prev', 'next', 'hidden');
                    
                    if (index === this.currentSlide) {
                        slide.classList.add('active');
                    } else if (index === this.getPrevIndex()) {
                        slide.classList.add('prev');
                    } else if (index === this.getNextIndex()) {
                        slide.classList.add('next');
                    } else {
                        slide.classList.add('hidden');
                    }
                });

                this.updateIndicators();
            }

            updateIndicators() {
                this.indicators.forEach((indicator, index) => {
                    indicator.classList.toggle('active', index === this.currentSlide);
                });
            }

            getPrevIndex() {
                return this.currentSlide === 0 ? this.totalSlides - 1 : this.currentSlide - 1;
            }

            getNextIndex() {
                return this.currentSlide === this.totalSlides - 1 ? 0 : this.currentSlide + 1;
            }

            nextSlide() {
                this.currentSlide = this.getNextIndex();
                this.updateSlidePositions();
            }

            prevSlide() {
                this.currentSlide = this.getPrevIndex();
                this.updateSlidePositions();
            }

            goToSlide(index) {
                this.currentSlide = index;
                this.updateSlidePositions();
            }

            startAutoPlay() {
                this.autoPlayInterval = setInterval(() => {
                    this.nextSlide();
                }, 8000);
            }

            pauseAutoPlay() {
                if (this.autoPlayInterval) {
                    clearInterval(this.autoPlayInterval);
                    this.autoPlayInterval = null;
                }
            }

            toggleAutoPlay() {
                const icon = this.playPauseBtn.querySelector('i');
                
                if (this.isPlaying) {
                    this.pauseAutoPlay();
                    this.isPlaying = false;
                    icon.className = 'bx bx-play';
                    this.playPauseBtn.title = 'Start auto-play';
                } else {
                    this.startAutoPlay();
                    this.isPlaying = true;
                    icon.className = 'bx bx-pause';
                    this.playPauseBtn.title = 'Pause auto-play';
                }
            }
        }

        // Initialize the slider when the page loads
        document.addEventListener('DOMContentLoaded', () => {
            new Testimonial3DSlider();
        });