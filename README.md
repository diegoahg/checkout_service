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

Install go dependences
    - [Testify] (https://github.com/stretchr/testify)
    - [Gorilla Mux] (https://github.com/gorilla/mux)

```
go get github.com/gorilla/mux
go get github.com/stretchr/testify/assert
```

# Use aplication

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