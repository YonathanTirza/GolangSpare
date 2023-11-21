// model.go

package main

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	nama     string
	username string
	password string
}

type Product struct {
	gorm.Model
	Name     string
	Price    int
	Quantity int
}
