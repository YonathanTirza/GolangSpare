// keranjang.js

document.addEventListener("DOMContentLoaded", function () {
    // Dummy data untuk item keranjang
    const cartData = [
        { name: "Ban Mobil", price: 20000, quantity: 2 },
        { name: "Aki Mobil", price: 50000, quantity: 1 },
        // Tambahkan item lain sesuai kebutuhan
    ];

    const cartContainer = document.getElementById("cart-items");

    // Tambahkan item keranjang ke halaman
    cartData.forEach(item => {
        const cartItem = createCartItem(item);
        cartContainer.appendChild(cartItem);
    });

    // Fungsi untuk membuat elemen item keranjang
    function createCartItem(item) {
        const cartItem = document.createElement("div");
        cartItem.classList.add("cart-item");

        const image = document.createElement("img");
        image.src = `/static/material/${item.name.toLowerCase().replace(/\s+/g, '-')}.jpg`; // Sesuaikan dengan struktur file gambar Anda
        image.alt = item.name;

        const cartItemContent = document.createElement("div");
        cartItemContent.classList.add("cart-item-content");

        const itemName = document.createElement("h3");
        itemName.textContent = item.name;

        const itemPrice = document.createElement("p");
        itemPrice.textContent = `Harga: Rp${item.price}`;

        const quantitySection = document.createElement("div");
        quantitySection.classList.add("quantity-section");

        const quantityInput = document.createElement("input");
        quantityInput.type = "text";
        quantityInput.value = item.quantity;
        quantityInput.readOnly = true;

        const incrementBtn = document.createElement("button");
        incrementBtn.classList.add("quantity-btn");
        incrementBtn.textContent = "+";

        const decrementBtn = document.createElement("button");
        decrementBtn.classList.add("quantity-btn");
        decrementBtn.textContent = "-";

        // Event listener untuk tombol tambah dan kurang
        incrementBtn.addEventListener("click", function () {
            // Logika penambahan jumlah
            // Tambahkan implementasi logika penambahan sesuai kebutuhan
        });

        decrementBtn.addEventListener("click", function () {
            // Logika pengurangan jumlah
            // Tambahkan implementasi logika pengurangan sesuai kebutuhan
        });

        // Susun elemen-elemen
        quantitySection.appendChild(incrementBtn);
        quantitySection.appendChild(quantityInput);
        quantitySection.appendChild(decrementBtn);

        cartItemContent.appendChild(itemName);
        cartItemContent.appendChild(itemPrice);

        cartItem.appendChild(image);
    }},)
