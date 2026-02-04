// Auth Page JavaScript

document.addEventListener('DOMContentLoaded', function() {
    // Get current page
    const isLoginPage = document.getElementById('loginForm') !== null;
    const isRegisterPage = document.getElementById('registerForm') !== null;

    // ===== LOGIN PAGE =====
    if (isLoginPage) {
        const loginForm = document.getElementById('loginForm');
        const loginBtn = document.getElementById('loginBtn');
        const errorMessage = document.getElementById('errorMessage');
        const errorText = document.getElementById('errorText');

        loginForm.addEventListener('submit', async function(e) {
            e.preventDefault();
            
            const email = document.getElementById('email').value;
            const password = document.getElementById('password').value;

            // Hide previous errors
            errorMessage.style.display = 'none';

            // Basic validation
            if (!email || !password) {
                showError('Please fill in all fields');
                return;
            }

            // Show loading state
            loginBtn.classList.add('loading');
            loginBtn.disabled = true;

            // Simulate API call
            await new Promise(resolve => setTimeout(resolve, 1000));

            // For demo, redirect to discover page
            window.location.href = '/discover';
        });

        function showError(message) {
            errorText.textContent = message;
            errorMessage.style.display = 'block';
        }
    }

    // ===== REGISTER PAGE =====
    if (isRegisterPage) {
        const registerForm = document.getElementById('registerForm');
        const registerBtn = document.getElementById('registerBtn');
        const passwordInput = document.getElementById('password');
        const confirmPasswordInput = document.getElementById('confirmPassword');
        const passwordStrength = document.getElementById('passwordStrength');
        const strengthLabel = document.getElementById('strengthLabel');
        const errorMessage = document.getElementById('errorMessage');
        const errorText = document.getElementById('errorText');

        // Password strength calculation
        function getPasswordStrength(password) {
            let score = 0;

            if (password.length >= 8) score++;
            if (password.length >= 12) score++;
            if (/[a-z]/.test(password) && /[A-Z]/.test(password)) score++;
            if (/\d/.test(password)) score++;
            if (/[^a-zA-Z0-9]/.test(password)) score++;

            if (score <= 1) return { score: 1, label: 'Weak', color: 'weak' };
            if (score === 2) return { score: 2, label: 'Fair', color: 'fair' };
            if (score === 3) return { score: 3, label: 'Good', color: 'good' };
            if (score === 4) return { score: 4, label: 'Strong', color: 'strong' };
            return { score: 5, label: 'Very Strong', color: 'very-strong' };
        }

        // Update password strength indicator
        passwordInput.addEventListener('input', function() {
            const password = passwordInput.value;

            if (password.length === 0) {
                passwordStrength.style.display = 'none';
                return;
            }

            passwordStrength.style.display = 'block';
            const strength = getPasswordStrength(password);

            // Update bars
            const bars = document.querySelectorAll('.strength-bar');
            bars.forEach((bar, index) => {
                // Reset classes
                bar.className = 'strength-bar';
                
                // Add active class based on score
                if (index < strength.score) {
                    bar.classList.add(`active-${strength.color}`);
                }
            });

            // Update label
            strengthLabel.textContent = strength.label;
            strengthLabel.className = `strength-label ${strength.color}`;
        });

        // Form submission
        registerForm.addEventListener('submit', async function(e) {
            e.preventDefault();

            const username = document.getElementById('username').value;
            const email = document.getElementById('email').value;
            const password = passwordInput.value;
            const confirmPassword = confirmPasswordInput.value;

            // Hide previous errors
            errorMessage.style.display = 'none';

            // Validation
            if (!username || !email || !password || !confirmPassword) {
                showError('Please fill in all fields');
                return;
            }

            if (password !== confirmPassword) {
                showError('Passwords do not match');
                return;
            }

            if (password.length < 8) {
                showError('Password must be at least 8 characters');
                return;
            }

            // Show loading state
            registerBtn.classList.add('loading');
            registerBtn.disabled = true;

            // Simulate API call
            await new Promise(resolve => setTimeout(resolve, 1000));

            // For demo, redirect to discover page
            window.location.href = '/discover';
        });

        function showError(message) {
            errorText.textContent = message;
            errorMessage.style.display = 'block';
        }
    }

    // ===== OAUTH BUTTONS (BOTH PAGES) =====
    const spotifyBtn = document.getElementById('spotifyBtn');
    const googleBtn = document.getElementById('googleBtn');
    const appleBtn = document.getElementById('appleBtn');

    if (spotifyBtn) {
        spotifyBtn.addEventListener('click', async function() {
            await handleOAuthLogin(spotifyBtn, 'spotify');
        });
    }

    if (googleBtn) {
        googleBtn.addEventListener('click', async function() {
            await handleOAuthLogin(googleBtn, 'google');
        });
    }

    if (appleBtn) {
        appleBtn.addEventListener('click', async function() {
            await handleOAuthLogin(appleBtn, 'apple');
        });
    }

    async function handleOAuthLogin(button, provider) {
        // Show loading state
        button.classList.add('loading');
        button.disabled = true;

        // Disable all buttons
        const allButtons = document.querySelectorAll('.btn');
        allButtons.forEach(btn => btn.disabled = true);

        // Simulate OAuth flow
        await new Promise(resolve => setTimeout(resolve, 1000));

        // For demo, redirect to discover page
        window.location.href = '/discover';
    }

    // ===== KEYBOARD SHORTCUTS =====
    document.addEventListener('keydown', function(e) {
        // Escape to clear error
        if (e.key === 'Escape') {
            const errorMessage = document.getElementById('errorMessage');
            if (errorMessage) {
                errorMessage.style.display = 'none';
            }
        }
    });
});