package app

import (
	"fmt"
	"sync"
)

var Checkout []Basket
var IDS = 0
var wg sync.WaitGroup

func CreateBasket() Basket {
	var b Basket
	wg.Add(1)
	go func() {
		defer wg.Done()
		IDS++
	}()
	wg.Wait()
	b.ID = IDS
	Checkout = append(Checkout, b)
	return b
}

func AddProduct(ID int, code string) error {
	for i, element := range Checkout {
		var item Item
		item.FillItem(code)
		if element.ID == ID {
			Checkout[i].Items = append(Checkout[i].Items, item)
			return nil
		}
	}
	return fmt.Errorf("Basket does not exist")
}

func GetAmount(ID int) (float64, error) {
	for _, element := range Checkout {
		if element.ID == ID {
			return element.GetTotal(), nil
		}
	}
	GetBaskets()
	return 0.0, fmt.Errorf("Basket does not exist")
}

func RemoveBasket(ID int) error {
	for i, element := range Checkout {
		if element.ID == ID {
			Checkout = append(Checkout[:i], Checkout[i+1:]...)
			return nil
		}
	}
	GetBaskets()
	return fmt.Errorf("Basket does not exist")
}

func GetBaskets() {
	for _, c := range Checkout {
		fmt.Printf("%#v\n", c)
	}
}
