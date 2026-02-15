// Theme toggle
const toggle = document.getElementById("themeToggle");
toggle.addEventListener("change", () => {
    document.body.classList.toggle("light", toggle.checked);
});

// Tabs: Login/Register
const tabLogin = document.getElementById("tabLogin");
const tabRegister = document.getElementById("tabRegister");
const loginForm = document.getElementById("loginForm");
const registerForm = document.getElementById("registerForm");

const toRegisterBottom = document.getElementById("toRegisterBottom");
const toLoginBottom = document.getElementById("toLoginBottom");

function showLogin(){
    if(!tabLogin) return;
    tabLogin.classList.add("active");
    tabRegister.classList.remove("active");
    tabLogin.setAttribute("aria-selected", "true");
    tabRegister.setAttribute("aria-selected", "false");

    loginForm.classList.add("active");
    registerForm.classList.remove("active");
}

function showRegister(){
    if(!tabRegister) return;
    tabRegister.classList.add("active");
    tabLogin.classList.remove("active");
    tabRegister.setAttribute("aria-selected", "true");
    tabLogin.setAttribute("aria-selected", "false");

    registerForm.classList.add("active");
    loginForm.classList.remove("active");
}

// Wenn wir im LoggedIn-View sind, existieren die Tabs nicht -> null checks sind wichtig
if(tabLogin && tabRegister && loginForm && registerForm){
    tabLogin.addEventListener("click", showLogin);
    tabRegister.addEventListener("click", showRegister);
    if(toRegisterBottom) toRegisterBottom.addEventListener("click", showRegister);
    if(toLoginBottom) toLoginBottom.addEventListener("click", showLogin);
}
