document.addEventListener("DOMContentLoaded", function () {
    const loginForm = document.getElementById("login-form");

    if (!loginForm) {
        console.error("Ошибка: Форма логина не найдена!");
        return;
    }

    function saveToken(token) {
        localStorage.setItem("authToken", token);
    }

    loginForm.addEventListener("submit", function (e) {
        e.preventDefault();

        const email = document.getElementById("email").value;
        const password = document.getElementById("password").value;

        fetch("http://localhost:8080/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ email, password }),
        })
            .then(response => response.json())
            .then(data => {
                if (data.token) {
                    saveToken(data.token); // Сохраняем токен в localStorage
                    alert("Login successful!");

                    window.location.replace("home.html"); // Перенаправление на главную

                } else {
                    alert("Error: " + (data.error || "Invalid credentials"));
                }
            })
            .catch(error => console.error("Error:", error));
    });
});
