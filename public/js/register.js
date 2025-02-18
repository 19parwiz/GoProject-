document.addEventListener("DOMContentLoaded", function () {
    const registerForm = document.getElementById("register-form");

    if (!registerForm) {
        console.error("Ошибка: Форма регистрации не найдена!");
        return;
    }

    registerForm.addEventListener("submit", function (e) {
        e.preventDefault();

        const name = document.getElementById("name").value;
        const email = document.getElementById("email").value;
        const password = document.getElementById("password").value;

        fetch("http://localhost:8080/register", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ name, email, password }),
        })
            .then(response => response.json())
            .then(data => {
                if (data.message) {
                    alert("Registration successful!");
                    window.location.href = "login.html";
                } else {
                    alert("Error: " + (data.error || "Unknown error"));
                }
            })
            .catch(error => console.error("Error:", error));
    });
});
