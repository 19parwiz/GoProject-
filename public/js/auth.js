function getToken() {
    return localStorage.getItem("authToken"); // Use "authToken" instead of "token"
}

function saveToken(token) {
    localStorage.setItem("authToken", token);
}

function removeToken() {
    localStorage.removeItem("authToken");
}

// Checking the token before loading pages
function checkAuth() {
    const token = getToken();
    console.log("Checking auth token:", token); // Log the token for checking
    if (!token) {
        window.location.href = "login.html"; // Redirect to login if no token
    }
}

// Fetching user data and displaying their name in the UI
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
                userNameElement.textContent = `Hello, ${data.name}`; //  Display user's name
            })
            .catch(error => {
                console.error("Failed to load user data:", error);
                userNameElement.textContent = "Guest"; // Display "Guest" if there is an error
            });
    }
}

// Perform the check on page load
document.addEventListener("DOMContentLoaded", function () {
    checkAuth(); // Check if the user is logged in
    loadUserData(); //Load the user data


    // Logout from the account
    const logoutBtn = document.getElementById("logout-btn");
    if (logoutBtn) {
        logoutBtn.addEventListener("click", function () {
            removeToken();
            alert("Logged out successfully!");
            window.location.href = "login.html";
        });
    }
});
