package app

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBasketHandlerOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/create-basket", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBasketHandler)
	handler.ServeHTTP(rr, req)

	e := "{\n  \"ID\": 1,\n  \"Items\": null,\n  \"Total\": 0\n}"

	assert.Equal(t, 201, rr.Code)
	assert.Equal(t, e, rr.Body.String())
	assert.Nil(t, err)
}

func TestAddProductHandlerOK(t *testing.T) {
	CreateBasket()
	var jsonStr = []byte(`{"busket_id": 1,"code": "VOUCHER"}`)
	req, err := http.NewRequest("POST", "/api/add-product", bytes.NewBuffer(jsonStr))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddProductHandler)
	handler.ServeHTTP(rr, req)

	e := "{\n  \"message\": \"Product successfully added\"\n}"

	assert.Equal(t, 201, rr.Code)
	assert.Equal(t, e, rr.Body.String())
	assert.Nil(t, err)
}

func TestAddProductHandlerErrorRequest(t *testing.T) {
	var jsonStr = []byte(``)
	req, err := http.NewRequest("POST", "/api/add-product", bytes.NewBuffer(jsonStr))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddProductHandler)
	handler.ServeHTTP(rr, req)
	e := "{}"
	assert.Equal(t, 400, rr.Code)
	assert.Nil(t, err)
	assert.Equal(t, e, rr.Body.String())
}

func TestAddProductHandlerErrorResponse(t *testing.T) {
	var jsonStr = []byte(`{"busket_id": 1000,"code": "VOUCHER"}`)
	req, err := http.NewRequest("POST", "/api/add-product", bytes.NewBuffer(jsonStr))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddProductHandler)
	handler.ServeHTTP(rr, req)
	e := "{\n  \"message\": \"Basket does not exist\"\n}"
	assert.Equal(t, 400, rr.Code)
	assert.Nil(t, err)
	assert.Equal(t, e, rr.Body.String())
}

func TestGetAmountHandlerOK(t *testing.T) {
	CreateBasket()
	var jsonStr = []byte(`{"busket_id": 1}`)
	req, err := http.NewRequest("POST", "/api/get-amount", bytes.NewBuffer(jsonStr))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAmountHandler)
	handler.ServeHTTP(rr, req)

	e := "{\n  \"amont\": 5\n}"

	assert.Equal(t, 201, rr.Code)
	assert.Equal(t, e, rr.Body.String())
	assert.Nil(t, err)
}

func TestGetAmountHandlerErrorRequest(t *testing.T) {
	var jsonStr = []byte(``)
	req, err := http.NewRequest("POST", "/api/get-amount", bytes.NewBuffer(jsonStr))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAmountHandler)
	handler.ServeHTTP(rr, req)
	e := "{}"
	assert.Equal(t, 400, rr.Code)
	assert.Nil(t, err)
	assert.Equal(t, e, rr.Body.String())
}

func TestGetAmountHandlerErrorResponse(t *testing.T) {
	var jsonStr = []byte(`{"busket_id": 1000,"code": "VOUCHER"}`)
	req, err := http.NewRequest("POST", "/api/add-product", bytes.NewBuffer(jsonStr))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAmountHandler)
	handler.ServeHTTP(rr, req)
	e := "{\n  \"message\": \"Basket does not exist\"\n}"
	assert.Equal(t, 400, rr.Code)
	assert.Nil(t, err)
	assert.Equal(t, e, rr.Body.String())
}

func TestRemoveBasketHandlerOK(t *testing.T) {
	CreateBasket()
	var jsonStr = []byte(`{"busket_id": 2}`)
	req, err := http.NewRequest("DELETE", "/api/get-amount", bytes.NewBuffer(jsonStr))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RemoveBasketHandler)
	handler.ServeHTTP(rr, req)

	e := "{\n  \"message\": \"Basket successfully removed\"\n}"

	assert.Equal(t, 201, rr.Code)
	assert.Equal(t, e, rr.Body.String())
	assert.Nil(t, err)
}

func TestRemoveBasketHandlerrErrorRequest(t *testing.T) {
	var jsonStr = []byte(``)
	req, err := http.NewRequest("POST", "/api/add-product", bytes.NewBuffer(jsonStr))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RemoveBasketHandler)
	handler.ServeHTTP(rr, req)
	e := "{}"
	assert.Equal(t, 400, rr.Code)
	assert.Nil(t, err)
	assert.Equal(t, e, rr.Body.String())
}

func TestRemoveBasketHandlerErrorResponse(t *testing.T) {
	var jsonStr = []byte(`{"busket_id": 1000,"code": "VOUCHER"}`)
	req, err := http.NewRequest("POST", "/api/add-product", bytes.NewBuffer(jsonStr))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RemoveBasketHandler)
	handler.ServeHTTP(rr, req)
	e := "{\n  \"message\": \"Basket does not exist\"\n}"
	assert.Equal(t, 400, rr.Code)
	assert.Nil(t, err)
	assert.Equal(t, e, rr.Body.String())
}
