document.addEventListener("DOMContentLoaded", function () {
    // Icons initialisieren
    if (window.lucide) lucide.createIcons();

    // Jahr im Footer setzen
    const y = document.getElementById("year");
    if (y) y.textContent = new Date().getFullYear();

    // Password Match Logik
    const password = document.getElementById("regPassword");
    const confirmPassword = document.getElementById("regConfirmPassword");
    const icon = document.getElementById("confirmIcon");
    const hint = document.getElementById("pwHint");

    if (!password || !confirmPassword || !icon) return;

    function clearState() {
        icon.classList.remove("match-success", "match-error");
        confirmPassword.classList.remove("match-success", "match-error");
        if (hint) hint.textContent = "";
    }

    function setSuccess() {
        icon.classList.add("match-success");
        icon.classList.remove("match-error");

        confirmPassword.classList.add("match-success");
        confirmPassword.classList.remove("match-error");

        hint.classList.add("match-success");
        hint.classList.remove("match-error");

        hint.textContent = "Passwörter stimmen überein.";
    }

    function setError() {
        icon.classList.add("match-error");
        icon.classList.remove("match-success");

        confirmPassword.classList.add("match-error");
        confirmPassword.classList.remove("match-success");

        hint.classList.add("match-error");
        hint.classList.remove("match-success");

        hint.textContent = "Passwörter stimmen nicht überein.";
    }

    function checkPasswords() {
        if (!confirmPassword.value) {
            clearState();
            return;
        }

        if (password.value === confirmPassword.value) {
            setSuccess();
        } else {
            setError();
        }
    }

    password.addEventListener("input", checkPasswords);
    confirmPassword.addEventListener("input", checkPasswords);

    // Initial check falls Browser autofill benutzt
    checkPasswords();
});
