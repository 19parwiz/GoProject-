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

    // === ДОБАВЛЕНИЕ КОРЗИНЫ НА СТРАНИЦУ ===
    const cartItemsContainer = document.querySelector(".list-group");
    const totalPriceElement = document.querySelector(".col-md-4 h1");

    let cart = JSON.parse(localStorage.getItem("cart")) || [];

    cartItemsContainer.innerHTML = ""; // Очищаем список перед рендерингом

    let totalPrice = 0;

    cart.forEach((item, index) => {
        const listItem = document.createElement("li");
        listItem.classList.add("list-group-item", "d-flex", "justify-content-between", "align-items-center");
        listItem.innerHTML = `
            <div>
                <img src="${item.imgSrc}" alt="${item.title}" style="width: 50px; height: 50px; object-fit: cover; margin-right: 10px;">
                ${item.title} - $${item.price}
            </div>
            <i class="fas fa-times text-danger remove-item" data-index="${index}" style="cursor: pointer;"></i>
        `;
        cartItemsContainer.appendChild(listItem);
        totalPrice += parseFloat(item.price);
    });

    totalPriceElement.textContent = `$${totalPrice.toFixed(2)}`;

    // Удаление товаров из корзины
    document.querySelectorAll(".remove-item").forEach(button => {
        button.addEventListener("click", function () {
            const index = this.dataset.index;
            cart.splice(index, 1);
            localStorage.setItem("cart", JSON.stringify(cart));
            location.reload(); // Перезагрузка страницы для обновления списка
        });
    });
});

document.addEventListener("DOMContentLoaded", function () {
    const continueShoppingBtn = document.getElementById("continue-shopping");
    if (continueShoppingBtn) {
        continueShoppingBtn.addEventListener("click", function () {
            window.location.href = "catalog.html"; // Измени, если у каталога другой путь
        });
    }
});

document.getElementById("checkout-btn").addEventListener("click", function () {
    const token = getToken();
    if (!token) {
        alert("You must be logged in to place an order!");
        return;
    }

    const cartItems = JSON.parse(localStorage.getItem("cart")) || [];
    if (cartItems.length === 0) {
        alert("Your cart is empty!");
        return;
    }

    localStorage.setItem("order", JSON.stringify(cartItems)); // Сохраняем заказ

    fetch("http://localhost:8080/orders", {
        method: "POST",
        headers: { "Content-Type": "application/json", "Authorization": `Bearer ${token}` },
        body: JSON.stringify({ total_price: 99.99 }) // Цена берётся из корзины
    })
    .then(response => response.json())
    .then(data => {
        alert("Order placed successfully!");
        localStorage.removeItem("cart"); // Очищаем корзину после оформления заказа
        window.location.href = "order.html";
    })
    .catch(error => console.error("Error placing order:", error));
});
