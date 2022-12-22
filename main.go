package main

import (
	"encoding/json"
	"fmt"
	"go-crud-postgres/dao"
	"go-crud-postgres/dto"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/products/", GetProducts).Methods("GET")
	router.HandleFunc("/products", InsertProducts).Methods("POST")
	router.HandleFunc("/products/{productId}", UpdateProducts).Methods("PUT")
	router.HandleFunc("/products/{productId}", DeleteProducts).Methods("DELETE")

	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
	http.ListenAndServe(":8000", router)

}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []dto.Product
	dbCon := dao.ConnectDB()
	v, err := dbCon.Query("select * from public.products")
	for v.Next() {
		var id int
		var productId string
		var productName string

		err = v.Scan(&id, &productId, &productName)

		if err != nil {
			panic(err)
		}

		products = append(products, dto.Product{ProductId: productId, ProductName: productName})
	}

	var response = dto.JsonResponse{Type: "success", Data: products}

	json.NewEncoder(w).Encode(response)
}

func InsertProducts(w http.ResponseWriter, r *http.Request) {
	var product dto.Product
	db := dao.ConnectDB()
	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&product)
	var productId string = product.ProductId
	var productName string= product.ProductName

	v, err := db.Query("Insert into public.products(productId, productName) values($1,$2);", productId, productName)
	if err != nil {
		panic(err)
	}
	fmt.Println("rows affected : ", v)
	var response = dto.JsonResponse{Type: "success", Message: "inserted the record"}

	json.NewEncoder(w).Encode(response)
}

func UpdateProducts(w http.ResponseWriter, r *http.Request) {
	db := dao.ConnectDB()

	params := mux.Vars(r)
	productId := params["productId"]

	db.Query("update public.products set productname = 'Portugal Jersey' where productid = $1",productId);

	var response = dto.JsonResponse{Type: "success", Message: "Updated the record"}

	json.NewEncoder(w).Encode(response)
}

func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	db := dao.ConnectDB()

	params := mux.Vars(r)
	productId := params["productId"]

	db.Query("delete from public.products where productid = $1",productId);

	var response = dto.JsonResponse{Type: "success", Message: "Deleted  the record"}

	json.NewEncoder(w).Encode(response)
}
