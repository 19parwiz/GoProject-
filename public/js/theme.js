document.addEventListener("DOMContentLoaded", function () {
    const darkModeToggle = document.getElementById("darkMode");

    // Проверяем, включена ли тёмная тема в localStorage
    if (localStorage.getItem("darkTheme") === "enabled") {
        document.body.classList.add("dark-theme");
        if (darkModeToggle) darkModeToggle.checked = true; // Ставим галочку
    }

    // Обработчик переключения темы
    if (darkModeToggle) {
        darkModeToggle.addEventListener("change", function () {
            if (darkModeToggle.checked) {
                document.body.classList.add("dark-theme");
                localStorage.setItem("darkTheme", "enabled"); // Сохраняем в localStorage
            } else {
                document.body.classList.remove("dark-theme");
                localStorage.setItem("darkTheme", "disabled"); // Убираем из localStorage
            }
        });
    }
});