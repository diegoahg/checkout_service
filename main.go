package main

import (
	"fmt"

	"github.com/diegoahg/checkout_service/app"
)

func main() {
	id := app.CreateBasket()
	app.AddProduct(id, "VOUCHER")
	app.AddProduct(id, "VOUCHER")
	app.AddProduct(id, "VOUCHER")
	app.AddProduct(id, "VOUCHER")
	app.AddProduct(id, "TSHIRT")
	app.AddProduct(id, "TSHIRT")
	app.AddProduct(id, "TSHIRT")
	app.AddProduct(id, "TSHIRT")
	app.AddProduct(id, "MUG")
	amount := app.GetAmount(id)
	samount := fmt.Sprintf("%.2f", amount)
	println("output: " + string(samount))
}
