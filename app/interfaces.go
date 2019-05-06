package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ProductInput takes incoming JSON payload for writing heart rate
type ProductInput struct {
	BasketID int    `json:"busket_id"`
	Code     string `json:"code"`
}

// BasketInput takes incoming JSON payload for writing heart rate
type BasketInput struct {
	BasketID int `json:"busket_id"`
}

// ResponseMessage takes incoming JSON payload for writing heart rate
type ResponseMessage struct {
	Message string `json:"message"`
}

// ResponseAmount takes incoming JSON payload for writing heart rate
type ResponseAmount struct {
	Amount float64 `json:"amont"`
}

// toJSONResponse convert a struct to Json Response
func toJSONResponse(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

// CreateBasketHandler is in charge to trigger a new Basket
func CreateBasketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateBasketHandler actived")
	b := CreateBasket()
	log.Println(fmt.Sprintf("Basket with ID: %d created", b.ID))
	toJSONResponse(w, r, http.StatusCreated, b)
}

// AddProductHandler is in charge to add a new Product in a Basket
func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AddProductHandler actived")
	var input ProductInput

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		toJSONResponse(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()
	var response ResponseMessage

	if err := AddProduct(input.BasketID, input.Code); err != nil {
		response.Message = fmt.Sprint(err)
		toJSONResponse(w, r, http.StatusBadRequest, response)
		return
	}

	response.Message = "Product successfully added"
	toJSONResponse(w, r, http.StatusCreated, response)
}

// GetAmountHandler is in charge to return the total of the Basket
func GetAmountHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAmountHandler actived")
	var input BasketInput

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		toJSONResponse(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	amount, err := GetAmount(input.BasketID)
	if err != nil {
		var response ResponseMessage
		response.Message = fmt.Sprint(err)
		toJSONResponse(w, r, http.StatusBadRequest, response)
		return
	}
	var response ResponseAmount
	response.Amount = amount
	toJSONResponse(w, r, http.StatusCreated, response)
}

// RemoveBasketHandler is in charge to remove a Basket
func RemoveBasketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("RemoveBasketHandler actived")
	var input BasketInput

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		toJSONResponse(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()
	var response ResponseMessage

	if err := RemoveBasket(input.BasketID); err != nil {
		response.Message = fmt.Sprint(err)
		toJSONResponse(w, r, http.StatusBadRequest, response)
		return
	}

	response.Message = "Basket successfully removed"
	toJSONResponse(w, r, http.StatusCreated, response)
}
