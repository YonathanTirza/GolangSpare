var currentQuantity = 0;

function ubahKuantitas(operator, button) {
    var card = button.closest(".card");
    var quantityText = card.querySelector('.quantity-input');

    if (!quantityText) {
        console.error('Elemen dengan ID "quantity" tidak ditemukan.');
        return;
    }

    currentQuantity = parseInt(quantityText.value);

    if (operator === 'increment') {
        currentQuantity += 1;
    } else if (operator === 'decrement' && currentQuantity > 0) {
        currentQuantity -= 1;
    }

    console.log('Current Quantity:', currentQuantity);

    quantityText.value = currentQuantity;
}

function tambahKeKeranjang(namaProduk, hargaProduk) {
    let totalHarga = currentQuantity * hargaProduk;
    alert(`Produk ${namaProduk} dengan harga ${hargaProduk} berjumlah ${currentQuantity} berhasil dibeli, total harga sekarang adalah ${totalHarga}`);
    // Logika tambahan untuk menambahkan produk ke keranjang belanja
}

function toggleSidebar() {
    var sidebar = document.getElementById("sidebar");
    var content = document.querySelector(".content");
    
    if (sidebar.style.width === "250px") {
        sidebar.style.width = "0";
        content.style.marginLeft = "0";
    } else {
        sidebar.style.width = "250px";
        content.style.marginLeft = "250px";
    }
}