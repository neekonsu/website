// Add this to the top of main.js
const routes = {
    '#/': {
        title: 'Home',
        render: () => {
            return `
                <section class="hero-section">
                    <div class="hero-content">
                        <h1>Welcome to our Website</h1>
                        <p class="hero-subtitle">Discover amazing content</p>
                    </div>
                </section>
                <section class="featured-posts">
                    <h2>Featured Posts</h2>
                    <div class="post-grid" id="featured-posts"></div>
                </section>
            `;
        }
    },
    '#/blog': {
        title: 'Blog',
        render: () => {
            return `
                <div class="blog-container">
                    <h1>Blog Posts</h1>
                    <div id="blog-posts" class="blog-posts"></div>
                </div>
            `;
        },
        afterRender: () => loadBlogPosts()
    },
    '#/contact': {
        title: 'Contact',
        render: () => {
            return `
                <div class="contact-container">
                    <h1>Contact Us</h1>
                    <form id="contact-form" class="contact-form">
                        <div class="form-group">
                            <label for="name">Name:</label>
                            <input type="text" id="name" name="name" required>
                        </div>
                        <div class="form-group">
                            <label for="email">Email:</label>
                            <input type="email" id="email" name="email" required>
                        </div>
                        <div class="form-group">
                            <label for="message">Message:</label>
                            <textarea id="message" name="message" required></textarea>
                        </div>
                        <button type="submit">Send Message</button>
                    </form>
                </div>
            `;
        },
        afterRender: () => setupContactForm()
    },
    '#/about': {
        title: 'About',
        render: () => {
            return `
                <div class="about-container">
                    <h1>About Us</h1>
                    <p>Learn more about our story and mission.</p>
                </div>
            `;
        }
    }
};

// Router function
function router() {
    const hash = window.location.hash || '#/';
    const route = routes[hash] || {
        title: '404 - Not Found',
        render: () => `
            <div class="error-page">
                <h1>404 - Page Not Found</h1>
                <p>The page you're looking for doesn't exist.</p>
                <a href="#/" class="button">Go Home</a>
            </div>
        `
    };

    // Update title
    document.title = `${route.title} - My Website`;

    // Render content
    const main = document.querySelector('main');
    main.innerHTML = route.render();

    // Run after-render functions if they exist
    if (route.afterRender) {
        route.afterRender();
    }

    // Update active nav item
    highlightCurrentNavItem();
}

// Handle navigation
function handleNavigation(e) {
    if (e.target.matches('nav a')) {
        const href = e.target.getAttribute('href');
        if (href.startsWith('#/')) {
            e.preventDefault();
            router();
        }
    }
}

// Initialize router
function initializeRouter() {
    // Handle initial page load
    if (!window.location.hash) {
        window.location.hash = '#/';
    }
    router();

    // Handle hash changes
    window.addEventListener('hashchange', router);

    // Handle navigation clicks
    document.addEventListener('click', handleNavigation);
}

// Update your initializeApp function
function initializeApp() {
    try {
        console.log("Initializing application...");
        
        // Initialize router
        initializeRouter();
        
        // Initialize core functionality
        addSmoothScrolling();
        
        console.log("Application initialized successfully");
    } catch (error) {
        console.error("Failed to initialize application:", error);
    }
}

// Global application state
const APP = {
    config: {
        apiBaseUrl: '/api',
        debugMode: true,
        routes: {
            home: '/',
            blog: '/blog',
            contact: '/contact',
            about: '/about'
        }
    },
    state: {
        isLoading: false,
        currentUser: null,
        darkMode: localStorage.getItem('darkMode') === 'true'
    }
};

// Application initialization
document.addEventListener('DOMContentLoaded', function() {
    initializeApp();
});

// Core initialization function
function initializeApp() {
    try {
        logMessage('Initializing application...');
        
        // Initialize core features
        initializeTheme();
        initializeNavigation();
        addEventListeners();
        
        // Route-specific initialization
        handleRouteSpecificFeatures();
        
        // Initialize common components
        initializeCommonComponents();
        
        logMessage('Application initialized successfully');
    } catch (error) {
        logError('Failed to initialize application:', error);
    }
}

// Theme initialization and management
function initializeTheme() {
    const prefersDark = window.matchMedia('(prefers-color-scheme: dark)');
    document.documentElement.classList.toggle('dark-mode', APP.state.darkMode);
    
    prefersDark.addEventListener('change', (e) => {
        if (APP.state.darkMode === null) {
            document.documentElement.classList.toggle('dark-mode', e.matches);
        }
    });
}

function toggleDarkMode() {
    APP.state.darkMode = !APP.state.darkMode;
    document.documentElement.classList.toggle('dark-mode', APP.state.darkMode);
    localStorage.setItem('darkMode', APP.state.darkMode);
}

// Navigation and routing
function initializeNavigation() {
    addSmoothScrolling();
    highlightCurrentNavItem();
    handlePopState();
}

function handlePopState() {
    window.addEventListener('popstate', (event) => {
        handleRouteSpecificFeatures();
    });
}

function navigateTo(path, title = '') {
    window.history.pushState({}, title, path);
    handleRouteSpecificFeatures();
}

// Route-specific feature handling
function handleRouteSpecificFeatures() {
    const currentPath = window.location.pathname;
    const routes = {
        '/': initializeHomePage,
        '/blog': initializeBlogPage,
        '/contact': initializeContactPage,
        '/about': initializeAboutPage
    };

    // Clear existing content
    const mainContent = document.querySelector('main');
    if (mainContent) {
        mainContent.innerHTML = '';
        showLoader(mainContent);
    }

    const initFunction = routes[currentPath];
    if (initFunction) {
        try {
            initFunction();
        } catch (error) {
            logError(`Failed to initialize route ${currentPath}:`, error);
            showErrorMessage('Failed to load page content');
        }
    } else {
        handle404();
    }
}

// Page-specific initializations
function initializeHomePage() {
    logMessage('Initializing home page...');
    // Add home page specific initialization
}

function initializeBlogPage() {
    logMessage('Initializing blog page...');
    loadBlogPosts();
}

function initializeContactPage() {
    logMessage('Initializing contact page...');
    setupContactForm();
}

function initializeAboutPage() {
    logMessage('Initializing about page...');
    // Add about page specific initialization
}

// Common components initialization
function initializeCommonComponents() {
    initializeModals();
    initializeTooltips();
    initializeDropdowns();
}

// Event listeners
function addEventListeners() {
    // Global event listeners
    document.addEventListener('click', handleGlobalClick);
    window.addEventListener('resize', debounce(handleResize, 250));
    window.addEventListener('scroll', debounce(handleScroll, 100));
    
    // Theme toggle
    const themeToggle = document.getElementById('theme-toggle');
    if (themeToggle) {
        themeToggle.addEventListener('click', toggleDarkMode);
    }
}

// UI Components
function initializeModals() {
    document.querySelectorAll('[data-modal-trigger]').forEach(trigger => {
        trigger.addEventListener('click', (e) => {
            const modalId = trigger.dataset.modalTrigger;
            toggleModal(modalId);
        });
    });
}

function initializeTooltips() {
    document.querySelectorAll('[data-tooltip]').forEach(element => {
        element.addEventListener('mouseenter', showTooltip);
        element.addEventListener('mouseleave', hideTooltip);
    });
}

function initializeDropdowns() {
    document.querySelectorAll('.dropdown-trigger').forEach(trigger => {
        trigger.addEventListener('click', toggleDropdown);
    });
}

// Utility functions
function debounce(func, wait) {
    let timeout;
    return function executedFunction(...args) {
        const later = () => {
            clearTimeout(timeout);
            func(...args);
        };
        clearTimeout(timeout);
        timeout = setTimeout(later, wait);
    };
}

function logMessage(message) {
    if (APP.config.debugMode) {
        console.log(`[${new Date().toISOString()}] ${message}`);
    }
}

function logError(message, error) {
    console.error(`[${new Date().toISOString()}] ${message}`, error);
}

function showLoader(parent) {
    const loader = document.createElement('div');
    loader.className = 'loader';
    loader.innerHTML = '<span class="spinner"></span>';
    parent.appendChild(loader);
}

function showErrorMessage(message) {
    const errorContainer = document.createElement('div');
    errorContainer.className = 'error-message';
    errorContainer.textContent = message;
    document.body.appendChild(errorContainer);
    
    setTimeout(() => {
        errorContainer.remove();
    }, 5000);
}

function handle404() {
    const mainContent = document.querySelector('main');
    if (mainContent) {
        mainContent.innerHTML = `
            <div class="error-page">
                <h1>404 - Page Not Found</h1>
                <p>The page you're looking for doesn't exist.</p>
                <a href="/" class="button">Go Home</a>
            </div>
        `;
    }
}

// Event handlers
function handleGlobalClick(event) {
    // Close dropdowns when clicking outside
    if (!event.target.closest('.dropdown')) {
        document.querySelectorAll('.dropdown.active').forEach(dropdown => {
            dropdown.classList.remove('active');
        });
    }
}

function handleResize() {
    // Handle responsive behavior
    logMessage('Window resized');
}

function handleScroll() {
    // Handle scroll-based behaviors
    const scrollTop = window.pageYOffset || document.documentElement.scrollTop;
    const header = document.querySelector('header');
    
    if (header) {
        header.classList.toggle('scrolled', scrollTop > 0);
    }
}

// Navigation utilities
function addSmoothScrolling() {
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function (e) {
            e.preventDefault();
            const target = document.querySelector(this.getAttribute('href'));
            if (target) {
                target.scrollIntoView({
                    behavior: 'smooth'
                });
            }
        });
    });
}

function highlightCurrentNavItem() {
    const currentPath = window.location.pathname;
    document.querySelectorAll('nav a').forEach(link => {
        if (link.getAttribute('href') === currentPath) {
            link.classList.add('active');
        } else {
            link.classList.remove('active');
        }
    });
}

// Modal functionality
function toggleModal(modalId) {
    const modal = document.getElementById(modalId);
    if (modal) {
        const isVisible = modal.classList.contains('active');
        modal.classList.toggle('active');
        
        if (!isVisible) {
            // Add overlay
            const overlay = document.createElement('div');
            overlay.className = 'modal-overlay';
            overlay.addEventListener('click', () => toggleModal(modalId));
            document.body.appendChild(overlay);
        } else {
            // Remove overlay
            document.querySelector('.modal-overlay')?.remove();
        }
    }
}

// Tooltip functionality
function showTooltip(event) {
    const tooltip = document.createElement('div');
    tooltip.className = 'tooltip';
    tooltip.textContent = event.target.dataset.tooltip;
    document.body.appendChild(tooltip);

    const rect = event.target.getBoundingClientRect();
    tooltip.style.top = `${rect.top - tooltip.offsetHeight - 5}px`;
    tooltip.style.left = `${rect.left + (rect.width - tooltip.offsetWidth) / 2}px`;
}

function hideTooltip() {
    document.querySelector('.tooltip')?.remove();
}

// Dropdown functionality
function toggleDropdown(event) {
    const dropdown = event.target.closest('.dropdown');
    dropdown.classList.toggle('active');
}

// Blog functionality
async function loadBlogPosts() {
    const mainContent = document.querySelector('main');
    try {
        const response = await fetch(`${APP.config.apiBaseUrl}/posts`);
        if (!response.ok) throw new Error('Failed to fetch blog posts');
        
        const posts = await response.json();
        mainContent.innerHTML = `
            <div class="blog-posts">
                ${posts.map(post => `
                    <article class="blog-post">
                        <h2>${post.title}</h2>
                        <div class="post-meta">
                            <span class="date">${new Date(post.date).toLocaleDateString()}</span>
                            <span class="author">${post.author}</span>
                        </div>
                        <p>${post.excerpt}</p>
                        <a href="/blog/${post.id}" class="read-more">Read More</a>
                    </article>
                `).join('')}
            </div>
        `;
    } catch (error) {
        logError('Failed to load blog posts:', error);
        showErrorMessage('Failed to load blog posts');
    } finally {
        document.querySelector('.loader')?.remove();
    }
}

// Contact form functionality
function setupContactForm() {
    const mainContent = document.querySelector('main');
    mainContent.innerHTML = `
        <form id="contact-form" class="contact-form">
            <div class="form-group">
                <label for="name">Name:</label>
                <input type="text" id="name" name="name" required>
            </div>
            <div class="form-group">
                <label for="email">Email:</label>
                <input type="email" id="email" name="email" required>
            </div>
            <div class="form-group">
                <label for="message">Message:</label>
                <textarea id="message" name="message" required></textarea>
            </div>
            <button type="submit">Send Message</button>
        </form>
    `;

    document.getElementById('contact-form').addEventListener('submit', async (e) => {
        e.preventDefault();
        try {
            const formData = new FormData(e.target);
            const response = await fetch(`${APP.config.apiBaseUrl}/contact`, {
                method: 'POST',
                body: JSON.stringify(Object.fromEntries(formData)),
                headers: {
                    'Content-Type': 'application/json'
                }
            });
            
            if (!response.ok) throw new Error('Failed to send message');
            
            showErrorMessage('Message sent successfully!');
            e.target.reset();
        } catch (error) {
            logError('Failed to send message:', error);
            showErrorMessage('Failed to send message');
        }
    });
}