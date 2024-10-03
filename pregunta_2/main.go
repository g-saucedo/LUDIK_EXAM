package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null"`
	Stock       int     `json:"stock" gorm:"not null"`
	// CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	// UpdatedAt   time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

var db *gorm.DB
var err error

func init() {
	// cadena de conexion
	dsn := "symfony_user:symfony_password@tcp(127.0.0.1:3306)/laravel_db"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the databse: ", err)
	}
	db.AutoMigrate(&Product{})
}

func main() {
	fmt.Println("!Hola, Go¡")
	r := mux.NewRouter()
	r.HandleFunc("/products", CreateProduct).Methods("POST")
	r.HandleFunc("/products", GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	r.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	if product.Price < 0 || product.Stock < 0 {
		http.Error(w, "El precio y la cantidad en stock deben ser numeros positivos", http.StatusBadRequest)
		return
	}
	db.Create(&product)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []Product
	db.Find(&products)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID no valida", http.StatusBadRequest)
		return
	}
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product Product
	if err := db.First(&product, params["id"]).Error; err != nil {
		http.Error(w, "Producto no encontrado", http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&product)

	if product.Price < 0 || product.Stock < 0 {
		http.Error(w, "El precio y la cantidad en stock deben ser numeros positivos", http.StatusBadRequest)
		return
	}

	db.Save(&product)
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product Product
	if err := db.First(&product, params["id"]).Error; err != nil {
		http.Error(w, "Producto no encontrado", http.StatusBadRequest)
		return
	}
	db.Delete(&product)
	w.WriteHeader(http.StatusNoContent)
}
