// script.js
var currentQuantity = ubahKuantitas();

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
    // var quantity = parseInt(document.querySelector('.quantity-input').value);

    fetch('/tambah-keranjang', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({
            'namaProduk': namaProduk,
            'kuantitas': currentQuantity,
            'hargaProduk': hargaProduk,
        }),
    })
    .then(response => response.json())
    .then(data => {
        alert(data.message);
    })
    .catch(error => {
        console.error('Error:', error);
    });
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
