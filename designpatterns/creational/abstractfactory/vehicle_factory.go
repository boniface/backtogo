package abstractfactory

import (
	"github.com/kataras/go-errors"
	"fmt"
)

type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

const(
	CarFactoryType =1
	MotorbikeFactoryType = 2
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognised\n", f))

	}
}

