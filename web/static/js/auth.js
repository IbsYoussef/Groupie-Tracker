// Auth Page JavaScript

document.addEventListener('DOMContentLoaded', function() {
    // Get current page
    const isRegisterPage = document.getElementById('registerForm') !== null;

    // ===== REGISTER PAGE - Password Strength Indicator Only =====
    if (isRegisterPage) {
        const passwordInput = document.getElementById('password');
        const passwordStrength = document.getElementById('passwordStrength');
        const strengthLabel = document.getElementById('strengthLabel');

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
        if (passwordInput && passwordStrength && strengthLabel) {
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
        }
    }

    // ===== OAUTH BUTTONS =====
    // TODO: Implement OAuth flow when ready
});

// Spotify OAuth button
const spotifyBtn = document.getElementById('spotifyBtn');
if (spotifyBtn) {
    spotifyBtn.addEventListener('click', function () {
        window.location.href = '/auth/spotify';
    })
}