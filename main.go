package main

import (
	"CRUD-Golang/config"
	"CRUD-Golang/controllers/categorycontroller"
	"CRUD-Golang/controllers/homecontroller"
	"CRUD-Golang/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. Home page
	http.HandleFunc("/", homecontroller.Welcome)

	// Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}