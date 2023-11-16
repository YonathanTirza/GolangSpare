package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/SpareTrial")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to the database successfully!")

	router := gin.Default()

	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")

	router.GET("/menu", func(c *gin.Context) {
		c.HTML(http.StatusOK, "menu.html", nil)
	})

	router.POST("/tambah-keranjang", func(c *gin.Context) {
		// Get data from the request
		namaProduk := c.PostForm("namaProduk")
		kuantitas := c.PostForm("kuantitas")
		hargaProduk := c.PostForm("hargaProduk")

		// Convert kuantitas and hargaProduk to appropriate types (e.g., int, float64)
		// ...
		// fmt.Print(namaProduk)
		// fmt.Print(kuantitas)
		// fmt.Print(hargaProduk)
		// Insert data into the database
		insertDataIntoDatabase(namaProduk, kuantitas, hargaProduk)

		c.JSON(http.StatusOK, gin.H{
			"message": "pesanan masuk ke keranjang (sorry berubah )",
		})
	})

	router.Run(":8080")
}

func insertDataIntoDatabase(namaProduk string, kuantitas string, hargaProduk string) {
	_, err := db.Exec("INSERT INTO keranjang (nama_barang, kuantitas, harga) VALUES (?, ?, ?)", namaProduk, kuantitas, hargaProduk)
	if err != nil {
		fmt.Println("Error inserting data into the database:", err)
		// Handle the error accordingly
	}
}
