package main

import (
	"log"
	"net/http"

	"github.com/diegoahg/checkout_service/app"
)

func main() {
	r := app.NewRouter()

	r.HandleFunc("/api/create-basket", app.CreateBasketHandler)
	r.HandleFunc("/api/add-product", app.AddProductHandler)
	r.HandleFunc("/api/get-amount", app.GetAmountHandler)
	r.HandleFunc("/api/remove-basket", app.RemoveBasketHandler).Methods("DELETE")

	log.Println("Initialized api")
	http.ListenAndServe(":8080", r)

}
