// Wait for the DOM to be fully loaded before running any scripts
document.addEventListener('DOMContentLoaded', function() {
    console.log("Main JavaScript file loaded");
    
    // Add smooth scrolling to all links
    addSmoothScrolling();
    
    // Add active class to current navigation item
    highlightCurrentNavItem();
    
    // If on the blog page, load blog posts
    if (window.location.pathname === '/blog') {
        loadBlogPosts();
    }
    
    // If on the contact page, set up form validation
    if (window.location.pathname === '/contact') {
        setupContactForm();
    }
});

// Function to add smooth scrolling to all links
function addSmoothScrolling() {
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function (e) {
            e.preventDefault();
            document.querySelector(this.getAttribute('href')).scrollIntoView({
                behavior: 'smooth'
            });
        });
    });
}

// Function to highlight the current navigation item
function highlightCurrentNavItem() {
    const currentPage = window.location.pathname;
    document.querySelectorAll('nav a').forEach(link => {
        if (link.getAttribute('href') === currentPage) {
            link.classList.add('active');
        }
    });
}

// Function to load blog posts (simulated)
function loadBlogPosts() {
    // This is a placeholder. In a real application, you might fetch this data from an API
    const posts = [
        { title: "First Blog Post", content: "This is the content of the first blog post." },
        { title: "Second Blog Post", content: "This is the content of the second blog post." },
        { title: "Third Blog Post", content: "This is the content of the third blog post." }
    ];
    
    const blogContainer = document.getElementById('blog-posts');
    if (blogContainer) {
        posts.forEach(post => {
            const postElement = document.createElement('article');
            postElement.innerHTML = `
                <h2>${post.title}</h2>
                <p>${post.content}</p>
            `;
            blogContainer.appendChild(postElement);
        });
    }
}

// Function to set up contact form validation
function setupContactForm() {
    const form = document.getElementById('contact-form');
    if (form) {
        form.addEventListener('submit', function(e) {
            e.preventDefault();
            if (validateForm()) {
                console.log('Form submitted successfully');
                // Here you would typically send the form data to your server
            }
        });
    }
}

// Function to validate the contact form
function validateForm() {
    let isValid = true;
    const name = document.getElementById('name');
    const email = document.getElementById('email');
    const message = document.getElementById('message');
    
    if (!name.value.trim()) {
        isValid = false;
        showError(name, 'Name is required');
    }
    
    if (!email.value.trim()) {
        isValid = false;
        showError(email, 'Email is required');
    } else if (!isValidEmail(email.value)) {
        isValid = false;
        showError(email, 'Please enter a valid email address');
    }
    
    if (!message.value.trim()) {
        isValid = false;
        showError(message, 'Message is required');
    }
    
    return isValid;
}

// Helper function to show error messages
function showError(input, message) {
    const errorElement = input.nextElementSibling;
    if (errorElement && errorElement.classList.contains('error-message')) {
        errorElement.textContent = message;
    } else {
        const error = document.createElement('div');
        error.className = 'error-message';
        error.textContent = message;
        input.parentNode.insertBefore(error, input.nextSibling);
    }
}

// Helper function to validate email format
function isValidEmail(email) {
    const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}
