package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	// Ganti dengan informasi koneksi database Anda
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/sparepartGo") // ini coba aja dulu, kalo ga bisa baru ganti
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}

	fmt.Println("Connected to the database!")
}

func main() {
	router := gin.Default()

	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	router.GET("/menu", func(c *gin.Context) {
		c.HTML(http.StatusOK, "menu.html", nil)
	})

	router.GET("/signin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signin.html", nil)
	})

	router.GET("/keranjang", func(c *gin.Context) {
		c.HTML(http.StatusOK, "keranjang.html", nil)
	})

	router.POST("/", gin.WrapF(loginHandler))
	router.POST("/signin", signInHandler)

	router.Run(":8080")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse nilai formulir dari permintaan
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Gagal mengurai formulir", http.StatusBadRequest)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	// Authentikasi pengguna
	if authenticateUser(username, password) {
		// Redirect ke halaman menu saat login berhasil
		http.Redirect(w, r, "/menu", http.StatusSeeOther)
	} else {
		// Tampilkan pesan kesalahan atau redirect kembali ke halaman login dengan pesan kesalahan
		http.Redirect(w, r, "/login?error=1", http.StatusSeeOther)
	}
}

func signInHandler(c *gin.Context) {
	// Ambil data dari form
	nama := c.PostForm("nama")
	username := c.PostForm("username")
	password := c.PostForm("password")


	// Simpan data ke database
	_, err := db.Exec("INSERT INTO users (nama, username, password) VALUES (?, ?, ?)", nama, username, password)
	if err != nil {
		fmt.Println("Error inserting user into database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign in"})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"message": "Sign in successful!"}) Di ganit jadi page success.html
	c.HTML(http.StatusOK, "success.html", gin.H{"Nama": nama})
}

func authenticateUser(username, password string) bool {
	// Kueri database untuk memeriksa apakah username dan password cocok
	kueri := "SELECT COUNT(*) FROM users WHERE username = ? AND password = ?"
	var jumlah int
	err := db.QueryRow(kueri, username, password).Scan(&jumlah)
	if err != nil {
		panic(err.Error())
	}

	return jumlah > 0
}
