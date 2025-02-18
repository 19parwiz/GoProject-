document.getElementById("checkout-btn").addEventListener("click", function () {
    const token = getToken();
    if (!token) {
        alert("You must be logged in to place an order!");
        return;
    }

    fetch("http://localhost:8080/orders", {
        method: "POST",
        headers: { "Content-Type": "application/json", "Authorization": `Bearer ${token}` },
        body: JSON.stringify({ total_price: 99.99 }) // Цена берётся из корзины
    })
        .then(response => response.json())
        .then(data => {
            alert("Order placed successfully!");
            window.location.href = "order.html";
        })
        .catch(error => console.error("Error placing order:", error));
});

document.addEventListener("DOMContentLoaded", function () {
    const checkoutBtn = document.getElementById("checkout-btn");
    if (checkoutBtn) {
        checkoutBtn.addEventListener("click", function () {
            console.log("Proceeding to checkout...");
        });
    } else {
        console.error("Checkout button not found!");
    }
});
