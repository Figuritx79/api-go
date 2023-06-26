package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Estructura de un producto
type Product struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Price int    `json:"Price"`
	Brand string `json:"Brand"`
}

// Arreglo de todos los productos
type AllProducts []Product

// Varible con todos los productos
var products = AllProducts{
	{
		ID:    1,
		Name:  "Laptop",
		Price: 20000,
		Brand: "HP",
	},
}

// Handler para la ruta principal
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Shop</h1")
}

// Handler para mostrar los productos
func getProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

// Handler para crear un producto
func createProduct(w http.ResponseWriter, r *http.Request) {

	// Almcenamos el cuerpo de la peticion en una variable
	reqBody, err := ioutil.ReadAll(r.Body)

	//Comprobamos si hay un error
	if err != nil {
		log.Fatal(err)
	}

	// Creamos una varible de tipo procucto
	var newProduct Product

	// Hacemos que el ID sea autoincrementable
	newProduct.ID = len(products) + 1

	//Convertimos el cuerpo de la peticion en un objeto y lo almacenamos en la variable newProduct
	json.Unmarshal(reqBody, &newProduct)

	// Agregamos el nuevo producto al arreglo de productos
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newProduct)
}

// Handler para buscar un producto
func searchProduct(w http.ResponseWriter, r *http.Request) {

	// Obtenemos las variables de la ruta
	vars := mux.Vars(r)

	productId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "ID INVALID")
	}

	for _, p := range products {
		if p.ID == productId {
			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(p)
		}

	}
}

// Handler para eliminar un producto
func delteProduct(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler)
	router.HandleFunc("/products", getProducts).Methods("GET")

	router.HandleFunc("/products", createProduct).Methods("POST")

	router.HandleFunc("/products/{id}", searchProduct).Methods("GET")

	router.HandleFunc("/products/{id}", delteProduct).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
