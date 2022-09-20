const loginButton = document.getElementById("login_button")
const askTime = document.getElementById("ask-time")

loginButton.addEventListener("click",function (){
    askTime.textContent = "button clicked"
});