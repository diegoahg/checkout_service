## Checkout Service

This repository contain checkout system of store made in golang.

The store contain the next products:

``` 
Code         | Name                |  Price
-------------------------------------------------
VOUCHER      | Cabify Voucher      |   5.00€
TSHIRT       | Cabify T-Shirt      |  20.00€
MUG          | Cafify Coffee Mug   |   7.50€
```

Those products have the next discounts:

 * `VOUCHER` 2x1.

 * If you buy 3 or more `TSHIRT` items, the price per unit should be 19.00€.

In this Api you can:
- Create a new checkout basket
- Add a product to a basket
- Get the total amount in a basket
- Remove the basket

# Before to start

Install go dependences:
- [Testify] (https://github.com/stretchr/testify)
- [Gorilla Mux] (https://github.com/gorilla/mux)

```
go get github.com/gorilla/mux
go get github.com/stretchr/testify/assert
```

ENV VARS:
HOST: localhost
PORT: 8080

Note: you can modify host and port in main.go

# Use aplication
Run
```
 go run main.go
```

## GET /api/create-basket
Create a new Basket

Without Request

Response example:
json
```
{
  "ID": 1,
  "Items": null,
  "Total": 0
}
```

## POST /api/add-product
Add product to a basket

Request example:
json
```
{
	"busket_id": 1,
	"code": "VOUCHER"
}
```

Response example:
json
```
{
  "message": "Product successfully added"
}
```

## POST /api/get-amount
Get amount of basket

Request example:
json
```
{
	"busket_id": 1
}
```

Response example:
json
```
{
  "amont": 5
}
```

## POST /api/get-amount
Remove a basket

Request example:
json
```
{
	"busket_id": 1
}
```

Response example:
json
```
{
  "message": "Basket successfully removed"
}
```

# Run Tests
In the "app" folder run
```
go test -cover -v
```

# Solution contain
- Race Condition in the ID's memory  (domain.go)
- Test and Coverage
- Api Rest
- Clean Architecture