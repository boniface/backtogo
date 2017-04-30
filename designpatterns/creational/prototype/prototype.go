package prototype

import "github.com/kataras/go-errors"

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const(
	White=1
	Black=2
	Blue =3
)

func GetShirtsCloner() ShirtCloner{

	return nil
}

type ShirtsCache struct{

}

var whitePrototype *Shirt = &Shirt{
	Price:15.00,
	SKU:"empty",
	Color: White,
}

func (s *ShirtsCache) GetClone(p int) (ItemInfoGetter, error)  {

	return nil, errors.New("Not Implementted")

}
