// ==== login ====
const openBtn = document.getElementById("openLogin");
const modal = document.getElementById("loginModal");

const closeBtnFront = document.getElementById("closeLogin");
const closeBtnBack = document.querySelector("[data-close-login]");

function openModal() {
    modal.classList.add("active");
    modal.setAttribute("aria-hidden", "false");
    document.body.style.overflow = "hidden";
}

function closeModal() {
    modal.classList.remove("active");
    modal.setAttribute("aria-hidden", "true");
    document.body.style.overflow = "";

    modal.addEventListener("transitionend", function handler(e) {
        if (e.propertyName === "opacity") {
            resetFlipToLogin();
            modal.removeEventListener("transitionend", handler);
        }
    });
}

openBtn.addEventListener("click", openModal);
closeBtnFront.addEventListener("click", closeModal);
closeBtnBack.addEventListener("click", closeModal);

// Klick auf Overlay schließt
modal.addEventListener("click", (e) => {
    if (e.target === modal) closeModal();
});

// ESC schließt
document.addEventListener("keydown", (e) => {
    if (e.key === "Escape" && modal.classList.contains("active")) {
        closeModal();
    }
});

// ==== flip card ====
const flip = document.getElementById("authFlip");
const forgotPwBtn = document.getElementById("forgotPwBtn");
const backToLoginBtn = document.getElementById("backToLoginBtn");

function showForgot() {
    flip.classList.add("is-flipped");
    const reset = document.getElementById("resetEmail");
    if (reset) setTimeout(() => reset.focus(), 200);
}

function showLogin() {
    flip.classList.remove("is-flipped");
    const email = document.getElementById("email");
    if (email) setTimeout(() => email.focus(), 200);
}

forgotPwBtn.addEventListener("click", showForgot);
backToLoginBtn.addEventListener("click", showLogin);

function resetFlipToLogin() {
    if (flip) flip.classList.remove("is-flipped");
}


// ==== Icons ====
document.addEventListener("DOMContentLoaded", function () {
    lucide.createIcons();
    document.getElementById("year").textContent = new Date().getFullYear();
});



