document.addEventListener("DOMContentLoaded", function () {
    const token = getToken();
    if (!token) {
        window.location.href = "login.html";
        return;
    }

    fetch("http://localhost:8080/orders", {
        method: "GET",
        headers: { "Content-Type": "application/json", "Authorization": `Bearer ${token}` }
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            const ordersList = document.getElementById("ordersList");
            if (!ordersList) {
                console.error("Element #ordersList not found!");
                return;
            }

            // Если статус отсутствует, заменяем на "unknown"
            ordersList.innerHTML = data.map(order => `
                <li class="list-group-item">Order #${order.id} - ${order.status || "unknown"}</li>
            `).join("");
        })
        .catch(error => console.error("Error fetching orders:", error));
});

function checkOrderStatus() {
    const orderNumber = document.getElementById("orderNumber").value;
    const statusDisplay = document.getElementById("orderStatus");

    if (!orderNumber) {
        statusDisplay.textContent = "Please enter a valid order number.";
        return;
    }

    // Симуляция ответа от сервера
    const statuses = ["Not Paid", "Paid", "Issued", "Cancelled"];
    const randomStatus = statuses[Math.floor(Math.random() * statuses.length)];

    statusDisplay.textContent = `Order #${orderNumber}: ${randomStatus}`;
}
