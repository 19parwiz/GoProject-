function getToken() {
    return localStorage.getItem("authToken"); // Используем "authToken" вместо "token"
}

function saveToken(token) {
    localStorage.setItem("authToken", token);
}

function removeToken() {
    localStorage.removeItem("authToken");
}

// Проверка токена перед загрузкой страниц
function checkAuth() {
    const token = getToken();
    console.log("Checking auth token:", token); // Логируем токен для проверки
    if (!token) {
        window.location.href = "login.html"; // Перенаправление на логин
    }
}

// Получение данных пользователя и отображение имени в UI
function loadUserData() {
    const token = getToken();
    const userNameElement = document.getElementById("user-name");

    if (token && userNameElement) {
        fetch("http://localhost:8080/api/user", {
            method: "GET",
            headers: { "Authorization": `Bearer ${token}` }
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error("Failed to fetch user data");
                }
                return response.json();
            })
            .then(data => {
                userNameElement.textContent = `Hello, ${data.name}`; // Отображаем имя пользователя
            })
            .catch(error => {
                console.error("Failed to load user data:", error);
                userNameElement.textContent = "Guest"; // Показываем "Guest", если ошибка
            });
    }
}

// Выполняем проверку при загрузке страницы
document.addEventListener("DOMContentLoaded", function () {
    checkAuth(); // Проверяем, залогинен ли пользователь
    loadUserData(); // Загружаем данные пользователя

    // Выход из аккаунта
    const logoutBtn = document.getElementById("logout-btn");
    if (logoutBtn) {
        logoutBtn.addEventListener("click", function () {
            removeToken();
            alert("Logged out successfully!");
            window.location.href = "login.html";
        });
    }
});
