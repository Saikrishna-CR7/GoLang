package main

import (
	"encoding/json"
	"fmt"
	"go-crud-postgres/dto"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main(){
	router := mux.NewRouter()


	router.HandleFunc("/products/", GetProducts).Methods("GET")
	router.HandleFunc("/products", InsertProducts).Methods("POST")
    router.HandleFunc("/products/{productId}", UpdateProducts).Methods("PUT")
    router.HandleFunc("/products/", DeleteProducts).Methods("DELETE")

    // serve the app
    fmt.Println("Server at 8080")
    log.Fatal(http.ListenAndServe(":8000", router))
	http.ListenAndServe(":8000", router)

}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products = []dto.Product{
		{ProductId: "1", ProductName: "SoccerBall"},
		{ProductId: "2", ProductName: "AirPump"},
		{ProductId: "3", ProductName: "Nike Studs"},
	}

	var response = dto.JsonResponse{Type: "success", Data: products}

    json.NewEncoder(w).Encode(response)
} 

func InsertProducts(w http.ResponseWriter, r *http.Request) {
	var products = []dto.Product{
		{ProductId: "1", ProductName: "SoccerBall"},
		{ProductId: "2", ProductName: "Bandages"},
		{ProductId: "3", ProductName: "Nike Studs"},
		{ProductId: r.FormValue("ProductId"), ProductName: r.FormValue("ProductName")},
	}
	//var product dto.Product
	//json.NewEncoder(r.Body).Decode(product)
	fmt.Println(r.Body)

	var response = dto.JsonResponse{Type: "success", Data: products}

    json.NewEncoder(w).Encode(response)
}

func UpdateProducts(w http.ResponseWriter, r *http.Request) {
	var products = []dto.Product{
		{ProductId: "1", ProductName: "SoccerBall"},
		{ProductId: "2", ProductName: r.FormValue("ProductName")},
		{ProductId: "3", ProductName: "Nike Studs"},
	}

	var response = dto.JsonResponse{Type: "success", Data: products}

    json.NewEncoder(w).Encode(response)
}

func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	var products = []dto.Product{
		{ProductId: "1", ProductName: "SoccerBall"},
		//{ProductId: "2", ProductName: "AirPump"},
		{ProductId: "3", ProductName: "Nike Studs"},
	}

	var response = dto.JsonResponse{Type: "success", Data: products}

    json.NewEncoder(w).Encode(response)
}