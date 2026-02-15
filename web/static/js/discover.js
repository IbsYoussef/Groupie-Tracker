// ============================================
// GROUPIE TRACKER - DISCOVER PAGE JS
// ============================================

document.addEventListener('DOMContentLoaded', () => {

    // === USER DROPDOWN ===
    const userMenuBtn   = document.getElementById('userMenuBtn');
    const userDropdown  = document.getElementById('userDropdown');
    const userContainer = document.getElementById('userMenuContainer');

    if (userMenuBtn && userDropdown) {
        userMenuBtn.addEventListener('click', () => {
            const isOpen = !userDropdown.hidden;
            userDropdown.hidden = isOpen;
            userMenuBtn.setAttribute('aria-expanded', String(!isOpen));
        });

        document.addEventListener('click', (e) => {
            if (!userContainer.contains(e.target)) {
                userDropdown.hidden = true;
                userMenuBtn.setAttribute('aria-expanded', 'false');
            }
        });
    }

    // === SEARCH BAR ===
    const searchInput     = document.getElementById('searchInput');
    const searchDropdown  = document.getElementById('searchDropdown');
    const searchResults   = document.getElementById('searchResults');
    const searchEmpty     = document.getElementById('searchEmpty');
    const searchContainer = document.getElementById('searchContainer');

    let searchTimeout = null;

    // Keyboard shortcut ⌘K / Ctrl+K
    document.addEventListener('keydown', (e) => {
        if ((e.metaKey || e.ctrlKey) && e.key === 'k') {
            e.preventDefault();
            searchInput?.focus();
        }
        if (e.key === 'Escape' && searchDropdown && !searchDropdown.hidden) {
            searchDropdown.hidden = true;
            searchInput.setAttribute('aria-expanded', 'false');
        }
    });

    if (searchInput) {
        searchInput.addEventListener('input', () => {
            clearTimeout(searchTimeout);
            const query = searchInput.value.trim();

            if (query.length === 0) {
                searchDropdown.hidden = true;
                searchInput.setAttribute('aria-expanded', 'false');
                return;
            }

            searchTimeout = setTimeout(() => performSearch(query), 300);
        });

        searchInput.addEventListener('focus', () => {
            const query = searchInput.value.trim();
            if (query.length > 0 && searchResults.children.length > 0) {
                searchDropdown.hidden = false;
                searchInput.setAttribute('aria-expanded', 'true');
            }
        });
    }

    if (searchContainer) {
        document.addEventListener('click', (e) => {
            if (!searchContainer.contains(e.target) && searchDropdown) {
                searchDropdown.hidden = true;
                searchInput?.setAttribute('aria-expanded', 'false');
            }
        });
    }

    async function performSearch(query) {
        // TODO: Replace with real API call to /api/search?q=query
        // For now shows loading state
        showLoading(true);
        showEmptyState(false);
        showNoResults(false);
        showArtistGrid(false);

        // Placeholder - will be wired to Spotify API
        console.log('Searching for:', query);
        showLoading(false);

        // Update results count
        const resultsCount = document.getElementById('resultsCount');
        if (resultsCount) resultsCount.textContent = `Searching for "${query}"...`;
    }

    // === SUGGESTION CHIPS ===
    document.querySelectorAll('.suggestion-chip').forEach(chip => {
        chip.addEventListener('click', () => {
            const query = chip.dataset.query;
            if (searchInput) {
                searchInput.value = query;
                searchInput.focus();
                performSearch(query);
            }
        });
    });

    // === FILTER SIDEBAR (COLLAPSIBLE SECTIONS) ===
    document.querySelectorAll('.filter-section-toggle').forEach(btn => {
        btn.addEventListener('click', () => {
            const targetId = btn.dataset.target;
            const target   = document.getElementById(targetId);
            const isOpen   = btn.getAttribute('aria-expanded') === 'true';

            btn.setAttribute('aria-expanded', String(!isOpen));
            if (target) target.hidden = isOpen;
        });
    });

    // === POPULARITY SLIDER ===
    const popularitySlider = document.getElementById('popularitySlider');
    const popularityValue  = document.getElementById('popularityValue');
    const popularityFill   = document.getElementById('popularityFill');

    if (popularitySlider) {
        popularitySlider.addEventListener('input', () => {
            const val = popularitySlider.value;
            if (popularityValue) popularityValue.textContent = `${val}%`;
            if (popularityFill)  popularityFill.style.width = `${val}%`;
            updateActiveFilters();
        });
    }

    // === GENRE SEARCH FILTER ===
    const genreSearch = document.getElementById('genreSearch');
    if (genreSearch) {
        genreSearch.addEventListener('input', () => {
            const query = genreSearch.value.toLowerCase();
            document.querySelectorAll('#genreCheckboxList .filter-checkbox-item').forEach(item => {
                const text = item.querySelector('.filter-checkbox-text')?.textContent.toLowerCase() || '';
                item.style.display = text.includes(query) ? '' : 'none';
            });
        });
    }

    // === FILTER CHECKBOXES & RADIO ===
    document.querySelectorAll('.filter-checkbox, .filter-radio, .filter-toggle-input').forEach(input => {
        input.addEventListener('change', updateActiveFilters);
    });

    // === ACTIVE FILTER TAGS ===
    function updateActiveFilters() {
        const tagsContainer = document.getElementById('activeFilterTags');
        const filterBadge   = document.getElementById('filterBadge');
        if (!tagsContainer) return;

        const tags = [];

        // Genre checkboxes
        document.querySelectorAll('.filter-checkbox[data-filter="genre"]:checked').forEach(cb => {
            tags.push({ label: cb.value, type: 'genre', value: cb.value });
        });

        // Popularity slider
        const popVal = parseInt(popularitySlider?.value || '0');
        if (popVal > 0) {
            tags.push({ label: `${popVal}%+ popularity`, type: 'popularity', value: popVal });
        }

        // Followers radio
        const followersChecked = document.querySelector('.filter-radio[data-filter="followers"]:checked');
        if (followersChecked && followersChecked.value !== 'any') {
            tags.push({ label: `${followersChecked.value} followers`, type: 'followers', value: followersChecked.value });
        }

        // Concerts toggle
        const concertsChecked = document.getElementById('upcomingConcerts');
        if (concertsChecked?.checked) {
            tags.push({ label: 'Has upcoming concerts', type: 'concerts', value: 'upcoming' });
        }

        // Render tags
        if (tags.length > 0) {
            tagsContainer.hidden = false;
            tagsContainer.innerHTML = tags.map(tag => `
                <span class="filter-tag">
                    ${tag.label}
                    <button type="button" class="filter-tag-remove" data-type="${tag.type}" data-value="${tag.value}" aria-label="Remove ${tag.label} filter">
                        <svg fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="3">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/>
                        </svg>
                    </button>
                </span>
            `).join('');

            // Bind remove buttons
            tagsContainer.querySelectorAll('.filter-tag-remove').forEach(btn => {
                btn.addEventListener('click', () => removeFilter(btn.dataset.type, btn.dataset.value));
            });
        } else {
            tagsContainer.hidden = true;
            tagsContainer.innerHTML = '';
        }

        // Update mobile badge
        if (filterBadge) {
            filterBadge.hidden  = tags.length === 0;
            filterBadge.textContent = tags.length;
        }
    }

    function removeFilter(type, value) {
        if (type === 'genre') {
            const cb = document.querySelector(`.filter-checkbox[data-filter="genre"][value="${value}"]`);
            if (cb) { cb.checked = false; }
        }
        if (type === 'popularity' && popularitySlider) {
            popularitySlider.value = 0;
            if (popularityValue) popularityValue.textContent = '0%';
            if (popularityFill)  popularityFill.style.width = '0%';
        }
        if (type === 'followers') {
            const anyRadio = document.querySelector('.filter-radio[data-filter="followers"][value="any"]');
            if (anyRadio) anyRadio.checked = true;
        }
        if (type === 'concerts') {
            const toggle = document.getElementById('upcomingConcerts');
            if (toggle) toggle.checked = false;
        }
        updateActiveFilters();
    }

    // === CLEAR ALL FILTERS ===
    const clearAllBtn = document.getElementById('clearAllFilters');
    const resetBtn    = document.getElementById('resetFiltersBtn');

    function clearAllFilters() {
        document.querySelectorAll('.filter-checkbox').forEach(cb => cb.checked = false);
        document.querySelectorAll('.filter-radio[value="any"]').forEach(r => r.checked = true);
        if (popularitySlider) { popularitySlider.value = 0; }
        if (popularityValue)  popularityValue.textContent = '0%';
        if (popularityFill)   popularityFill.style.width = '0%';
        const toggle = document.getElementById('upcomingConcerts');
        if (toggle) toggle.checked = false;
        updateActiveFilters();
    }

    clearAllBtn?.addEventListener('click', clearAllFilters);
    resetBtn?.addEventListener('click', clearAllFilters);

    // === MOBILE SIDEBAR TOGGLE ===
    const mobileFilterBtn  = document.getElementById('mobileFilterBtn');
    const filterSidebar    = document.getElementById('filterSidebar');
    const filterOverlay    = document.getElementById('filterOverlay');
    const filterMobileClose = document.getElementById('filterMobileClose');

    function openFilterSidebar() {
        filterSidebar?.classList.add('open');
        if (filterOverlay) filterOverlay.classList.add('active');
    }

    function closeFilterSidebar() {
        filterSidebar?.classList.remove('open');
        if (filterOverlay) filterOverlay.classList.remove('active');
    }

    mobileFilterBtn?.addEventListener('click', openFilterSidebar);
    filterMobileClose?.addEventListener('click', closeFilterSidebar);
    filterOverlay?.addEventListener('click', closeFilterSidebar);

    // === VIEW TOGGLE (GRID / LIST) ===
    const gridViewBtn = document.getElementById('gridViewBtn');
    const listViewBtn = document.getElementById('listViewBtn');
    const artistGrid  = document.getElementById('artistGrid');

    gridViewBtn?.addEventListener('click', () => {
        gridViewBtn.classList.add('active');
        listViewBtn?.classList.remove('active');
        artistGrid?.classList.remove('list-view');
    });

    listViewBtn?.addEventListener('click', () => {
        listViewBtn.classList.add('active');
        gridViewBtn?.classList.remove('active');
        artistGrid?.classList.add('list-view');
    });

    // === STATE HELPERS ===
    function showEmptyState(show) {
        const el = document.getElementById('emptyState');
        if (el) el.hidden = !show;
    }

    function showArtistGrid(show) {
        if (artistGrid) artistGrid.hidden = !show;
    }

    function showNoResults(show) {
        const el = document.getElementById('noResults');
        if (el) el.hidden = !show;
    }

    function showLoading(show) {
        const el = document.getElementById('loadingState');
        if (el) el.hidden = !show;
        if (show) showEmptyState(false);
    }

    // === SORT ===
    const sortSelect = document.getElementById('sortSelect');
    sortSelect?.addEventListener('change', () => {
        console.log('Sort changed to:', sortSelect.value);
        // TODO: Re-fetch results with sort parameter
    });

    // === INITIAL STATE ===
    showEmptyState(true);
    showArtistGrid(false);
    showNoResults(false);
    showLoading(false);

});