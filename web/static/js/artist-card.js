// ============================================
// ARTIST CARD - Play/Pause Track Preview
// ============================================

// Track the currently playing audio globally so only one plays at a time
let currentAudio = null;
let currentBtn = null;

/**
 * toggleTrackPlay - called from onclick on each play button
 * Handles: play, pause, switching between tracks, preview URL fallback
 */
function toggleTrackPlay(btn) {
    const trackEl = btn.closest('.artist-card-track');
    const previewURL = trackEl.dataset.preview;
    const iconPlay = btn.querySelector('.icon-play');
    const iconPause = btn.querySelector('.icon-pause');

    // If this button is already playing - pause it
    if (btn === currentBtn) {
        if (currentAudio) {
            currentAudio.pause();
            currentAudio = null;
        }
        setButtonState(btn, false);
        currentBtn = null;
        return;
    }

    // Stop whatever was playing before
    if (currentAudio) {
        currentAudio.pause();
        currentAudio = null;
    }
    if (currentBtn) {
        setButtonState(currentBtn, false);
    }

    // No preview URL available - show brief visual feedback only
    if (!previewURL) {
        setButtonState(btn, true);
        currentBtn = btn;
        setTimeout(() => {
            setButtonState(btn, false);
            currentBtn = null;
        }, 1500);
        return;
    }

    // Play the preview
    const audio = new Audio(previewURL);
    audio.volume = 0.7;

    audio.addEventListener('ended', () => {
        setButtonState(btn, false);
        currentBtn = null;
        currentAudio = null;
    });

    audio.addEventListener('error', () => {
        setButtonState(btn, false);
        currentBtn = null;
        currentAudio = null;
    });

    audio.play().then(() => {
        setButtonState(btn, true);
        currentBtn = btn;
        currentAudio = audio;
    }).catch(() => {
        // Autoplay blocked by browser - reset state
        setButtonState(btn, false);
        currentBtn = null;
    });
}

/**
 * setButtonState - toggles play/pause icon and .playing class
 */
function setButtonState(btn, playing) {
    const iconPlay = btn.querySelector('.icon-play');
    const iconPause = btn.querySelector('.icon-pause');

    if (playing) {
        btn.classList.add('playing');
        iconPlay.hidden = true;
        iconPause.hidden = false;
        btn.setAttribute('aria-label', btn.getAttribute('aria-label').replace('Play', 'Pause'));
    } else {
        btn.classList.remove('playing');
        iconPlay.hidden = false;
        iconPause.hidden = true;
        btn.setAttribute('aria-label', btn.getAttribute('aria-label').replace('Pause', 'Play'));
    }
}

// Stop audio when user navigates away
window.addEventListener('beforeunload', () => {
    if (currentAudio) currentAudio.pause();
});