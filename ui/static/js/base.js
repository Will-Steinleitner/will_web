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

        const searchGame = document.getElementById("searchGame")

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

            clearError(); // ✅ Fehler beim Schließen löschen

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

    document.addEventListener("DOMContentLoaded", () => {
        setupIconsAndYear();
        setupLoginModal();
    });
})();
