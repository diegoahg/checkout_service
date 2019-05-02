package app

var codeName = map[string]string{
	"VOUCHER": "Cabify Voucher",
	"TSHIRT":  "Cabify T-Shirt",
	"MUG":     "Cafify Coffee Mug",
}

var codePrice = map[string]float64{
	"VOUCHER": 5.00,
	"TSHIRT":  20.00,
	"MUG":     7.50,
}

// Item represents item of checkout system
type Item struct {
	Code  string
	Name  string
	Price float64
}

// Basket represents the container of items
type Basket struct {
	ID    int
	Items []Item
	Total float64
}

func (i *Item) FillItem(code string) {
	i.Code = code
	i.Name = codeName[code]
	i.Price = codePrice[code]
}

// AddItem add other item into de checkout
func (b *Basket) AddItem(item Item) {
	b.Items = append(b.Items, item)
	b.Total = b.GetTotal()
}

func (b *Basket) GetTotal() float64 {
	total := 0.0
	total += b.VoucherPromotion()
	total += b.TshirtPromotion()
	total += b.WithoutPromotion()
	return total
}

func (b *Basket) VoucherPromotion() float64 {
	count := 0
	for _, element := range b.Items {
		if element.Name == "VOUCHER" {
			count++
		}
	}
	hasPromotion := int(count / 2)
	hasNot := count - hasPromotion
	totalVocher := (float64(hasPromotion) * codePrice["VOUCHER"]) + (float64(hasNot) * codePrice["VOUCHER"])
	return totalVocher
}

func (b *Basket) TshirtPromotion() float64 {
	count := 0
	for _, element := range b.Items {
		if element.Name == "TSHIRT" {
			count++
		}
	}
	price := codePrice["TSHIRT"]
	if count > 3 {
		price = 19.00
	}
	totalTshirt := (float64(count) * price)
	return totalTshirt
}

func (b *Basket) WithoutPromotion() float64 {
	totalWithoutPromotion := 0.0
	for _, element := range b.Items {
		if element.Name != "VOUCHER" && element.Name != "TSHIRT" {
			totalWithoutPromotion += element.Price
		}
	}
	return totalWithoutPromotion
}
