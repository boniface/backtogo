package abstractfactory

import (
	"github.com/kataras/go-errors"
	"fmt"
)

const (
	LuxuryCarType =1
	FamilyCarType =2
)

type CarFactory struct {

}

func (c *CarFactory) Build(v int) (Vehicle, error)  {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognised \n", v))

	}

}