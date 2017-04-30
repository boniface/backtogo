package prototype

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor  byte

type Shirt struct {
	Price float32
	SKU string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return ""
}

func (i* Shirt) GetPrice()  float32 {
	return i.Price
}