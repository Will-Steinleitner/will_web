// static/js/base.js
(() => {
    function setupIconsAndYear() {
        if (window.lucide) lucide.createIcons();
        const year = document.getElementById("year");
        if (year) year.textContent = new Date().getFullYear();
    }

    function setupLoginModal() {
        const openBtn = document.getElementById("openLogin");
        const modal = document.getElementById("loginModal");

        // Wenn es auf der Seite kein Modal gibt: sauber raus
        if (!openBtn || !modal) return;

        const closeBtnFront = document.getElementById("closeLogin");
        // kann mehrere geben (du hast data-close-login evtl. mehrfach)
        const closeBtnsBack = modal.querySelectorAll("[data-close-login]");

        const flip = document.getElementById("authFlip");
        const forgotPwBtn = document.getElementById("forgotPwBtn");
        const backToLoginBtn = document.getElementById("backToLoginBtn");

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

            // Flip erst nach Fade-Out resetten (wie bei dir)
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

        // --- Event-Listener (mit Guards)
        openBtn.addEventListener("click", openModal);

        if (closeBtnFront) closeBtnFront.addEventListener("click", closeModal);
        closeBtnsBack.forEach((btn) => btn.addEventListener("click", closeModal));

        // Overlay click -> schließen
        modal.addEventListener("click", (e) => {
            if (e.target === modal) closeModal();
        });

        // ESC -> schließen
        document.addEventListener("keydown", (e) => {
            if (e.key === "Escape" && modal.classList.contains("active")) {
                closeModal();
            }
        });

        if (forgotPwBtn) forgotPwBtn.addEventListener("click", showForgot);
        if (backToLoginBtn) backToLoginBtn.addEventListener("click", showLogin);
    }

    document.addEventListener("DOMContentLoaded", () => {
        setupIconsAndYear();
        setupLoginModal();
    });
})();
