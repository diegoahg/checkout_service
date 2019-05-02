package app

var Checkout []Basket
var IDS = 0

func getID() int {
	return IDS + 1
}

func existID(ID int) bool {
	for _, c := range Checkout {
		if c.ID == ID {
			return true
		}
	}
	return false
}

func CreateBasket() int {
	var b Basket
	b.ID = getID()
	Checkout = append(Checkout, b)
	return b.ID
}

func AddProduct(ID int, code string) string {
	if !existID(ID) {
		return "Basket does not exist"
	}
	var i Item
	i.FillItem(code)
	Checkout[ID-1].Items = append(Checkout[ID-1].Items, i)
	return "Item added to basket" + string(ID)
}

func GetAmount(ID int) float64 {
	return Checkout[ID-1].GetTotal()
}

func RemoveBasket(ID int) string {
	i := ID - 1
	Checkout = append(Checkout[:i], Checkout[i+1:]...)
	return "Basket was removed"
}
