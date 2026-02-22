(() => {
    function setupIconsAndYear() {
        if (window.lucide) lucide.createIcons();
        const year = document.getElementById("year");
        if (year) year.textContent = new Date().getFullYear();
    }
    function setupLoginModal() {
        const openBtn = document.getElementById("openLogin");
        const modal = document.getElementById("loginModal");

        if (!modal) return;

        const closeBtnFront = document.getElementById("closeLogin");
        const closeBtnsBack = modal.querySelectorAll("[data-close-login]");

        const flip = document.getElementById("authFlip");
        const forgotPwBtn = document.getElementById("forgotPwBtn");
        const backToLoginBtn = document.getElementById("backToLoginBtn");

        const errorBox = modal.querySelector(".errorBox");

        function clearError() {
            if (errorBox) {
                errorBox.style.display = "none";
                errorBox.innerHTML = "";
            }
        }

        function resetFlipToLogin() {
            if (flip) flip.classList.remove("is-flipped");
        }

        function openModal() {
            modal.classList.add("active");
            modal.setAttribute("aria-hidden", "false");
            document.body.style.overflow = "hidden";
        }

        function closeModal() {
            modal.classList.remove("active");
            modal.setAttribute("aria-hidden", "true");
            document.body.style.overflow = "";

            clearError();

            const handler = (e) => {
                if (e.propertyName === "opacity") {
                    resetFlipToLogin();
                    modal.removeEventListener("transitionend", handler);
                }
            };
            modal.addEventListener("transitionend", handler);
        }

        function showForgot() {
            if (!flip) return;
            flip.classList.add("is-flipped");
            const reset = document.getElementById("resetEmail");
            if (reset) setTimeout(() => reset.focus(), 200);
        }

        function showLogin() {
            if (!flip) return;
            flip.classList.remove("is-flipped");
            const email = document.getElementById("email");
            if (email) setTimeout(() => email.focus(), 200);
        }

        if (openBtn) openBtn.addEventListener("click", openModal);
        if (closeBtnFront) closeBtnFront.addEventListener("click", closeModal);
        closeBtnsBack.forEach((btn) => btn.addEventListener("click", closeModal));

        // Overlay click -> close
        modal.addEventListener("click", (e) => {
            if (e.target === modal) closeModal();
        });

        // ESC -> close
        document.addEventListener("keydown", (e) => {
            if (e.key === "Escape" && modal.classList.contains("active")) {
                closeModal();
            }
        });

        if (forgotPwBtn) forgotPwBtn.addEventListener("click", showForgot);
        if (backToLoginBtn) backToLoginBtn.addEventListener("click", showLogin);

        // Serverseitig: Modal automatisch oeffen
        if (modal.dataset.autoOpen === "1") {
            openModal();
        }
    }
    function showToast(type, duration = 1000) {

        const toast = document.querySelector(`.toast-${type}`);
        if (!toast) return;

        toast.classList.add("show");

        setTimeout(() => {
            toast.classList.remove("show");
        }, duration);
    }
    function setupGameAddIcons() {
        const container = document.querySelector('.game-list-div[type="search-game"]');
        const searchInput = document.querySelector(".search-input");
        if (!container) return;

        const isLoggedInNow = () => document.body?.dataset?.loggedIn === "1";

        document.querySelectorAll('[data-lucide="diamond-plus"]')
            .forEach(icon => {

                icon.addEventListener("click", (e) => {
                    e.preventDefault();
                    e.stopPropagation();
                    e.stopImmediatePropagation();

                    if (!isLoggedInNow()) {
                        window.requireLogin?.(e);
                        return;
                    }

                    const item = icon.closest("a");
                    if (!item) return;

                    const gameClass = item.querySelector(".navImg")?.classList?.[1];
                    if (!gameClass) return;

                    const navbarItem = document.querySelector(`.navlinks .${gameClass}`);
                    if (!navbarItem) return;

                    const isAlreadyActive = navbarItem.classList.contains("is-active");
                    const activeCount = document.querySelectorAll('.navlinks .navlinkWithImg.is-active').length;

                    if (!isAlreadyActive && activeCount >= 4) {
                        showToast("error");
                        return;
                    }

                    icon.classList.toggle("active");
                    navbarItem.classList.toggle("is-active");

                    if (!isAlreadyActive) {
                        container.prepend(item);
                        showToast("success");
                    } else {
                        const activeItems = Array.from(container.querySelectorAll("a"))
                            .filter(a => a.querySelector('[data-lucide="diamond-plus"].active'));

                        if (activeItems.length === 0) {
                            container.appendChild(item);
                        } else {
                            activeItems[activeItems.length - 1].after(item);
                        }

                        if (searchInput) searchInput.focus();
                    }
                }, true);
            });
    }
    function setupGameSearchFilter() {
        const input = document.getElementById("searchInput");
        const list = document.querySelector('.game-list-div[type="search-game"]');
        if (!input || !list) return;

        const items = Array.from(list.querySelectorAll("a.search-item"));

        const filter = () => {
            const q = input.value.trim().toLowerCase();

            items.forEach(a => {
                const text = (a.querySelector("span")?.textContent || "").toLowerCase();
                a.style.display = text.includes(q) ? "" : "none";
            });
        };

        input.addEventListener("input", filter);
        filter();
    }

    window.requireLogin = function (e) {
        if (e) e.preventDefault();

        const openBtn = document.getElementById("openLogin");
        if (openBtn) {
            openBtn.click();
            return false;
        }

        const modal = document.getElementById("loginModal");
        if (modal) {
            modal.setAttribute("aria-hidden", "false");
            modal.classList.add("open");
        }
        return false;
    };
    document.addEventListener("DOMContentLoaded", () => {
        setupIconsAndYear();
        setupLoginModal();
        setupGameAddIcons();
        setupGameSearchFilter();
    });
})();
