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

// FillItem fills a new item
func (i *Item) FillItem(code string) {
	i.Code = code
	i.Name = codeName[code]
	i.Price = codePrice[code]
}

// AddItem adds other item into de checkout
func (b *Basket) AddItem(item Item) {
	b.Items = append(b.Items, item)
	b.Total = b.GetTotal()
}

// GetTotal returns the total with promotions
func (b *Basket) GetTotal() float64 {
	total := 0.0
	total += b.VoucherPromotion()
	total += b.TshirtPromotion()
	total += b.WithoutPromotion()
	b.Total = total
	return total
}

// VoucherPromotion returns promotion value about VOUCHER
func (b *Basket) VoucherPromotion() float64 {
	count := 0
	for _, element := range b.Items {
		if element.Code == "VOUCHER" {
			count++
		}
	}
	hasPromotion := int(count / 2)
	hasNot := count - (hasPromotion * 2)
	totalVocher := (float64(hasPromotion) * codePrice["VOUCHER"]) + (float64(hasNot) * codePrice["VOUCHER"])
	return totalVocher
}

// TshirtPromotion returns promotion value about TSHIRT
func (b *Basket) TshirtPromotion() float64 {
	count := 0
	for _, element := range b.Items {
		if element.Code == "TSHIRT" {
			count++
		}
	}
	price := codePrice["TSHIRT"]
	if count >= 3 {
		price = 19.00
	}
	totalTshirt := (float64(count) * price)
	return totalTshirt
}

// WithoutPromotion returns value of product witout promotion
func (b *Basket) WithoutPromotion() float64 {
	totalWithoutPromotion := 0.0
	for _, element := range b.Items {
		if element.Code != "VOUCHER" && element.Code != "TSHIRT" {
			totalWithoutPromotion += element.Price
		}
	}
	return totalWithoutPromotion
}
