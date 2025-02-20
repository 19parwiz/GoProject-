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

document.addEventListener("DOMContentLoaded", function () {
    const ordersList = document.getElementById("ordersList");
    const orderItems = JSON.parse(localStorage.getItem("order")) || [];

    if (orderItems.length === 0) {
        ordersList.innerHTML = "<li class='list-group-item'>No items in your order.</li>";
    } else {
        ordersList.innerHTML = ""; // Очищаем перед добавлением новых элементов
        orderItems.forEach(item => {
            const li = document.createElement("li");
            li.className = "list-group-item d-flex justify-content-between align-items-center";
            li.innerHTML = `${item.name} - $${item.price} <span class="badge bg-primary">${item.quantity}</span>`;
            ordersList.appendChild(li);
        });
    }
});

document.addEventListener("DOMContentLoaded", function () {
    const placeOrderBtn = document.querySelector(".btn-primary.w-100");

    placeOrderBtn.addEventListener("click", function () {
        const token = getToken();
        if (!token) {
            alert("You must be logged in to place an order.");
            return;
        }

        const order = {
            total_price: 100.0,  // Нужно заменить на реальную сумму
        };

        fetch("http://localhost:8080/orders", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(order)
        })
        .then(response => response.json())
        .then(data => {
            alert(`Order placed! Order Number: ${data.order_number}`);
            location.reload();
        })
        .catch(error => console.error("Error placing order:", error));
    });
});

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
    .then(response => response.json())
    .then(data => {
        const ordersList = document.getElementById("ordersList");

        if (data.length === 0) {
            ordersList.innerHTML = "<li class='list-group-item'>No orders found.</li>";
            return;
        }

        ordersList.innerHTML = data.map(order => `
            <li class="list-group-item">
                Order #${order.order_number} - ${order.status} - $${order.total_price}
            </li>
        `).join("");
    })
    .catch(error => console.error("Error fetching orders:", error));
});
