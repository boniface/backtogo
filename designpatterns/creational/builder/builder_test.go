package builder

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	// Create a Car Builder
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n",
			car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s \n",
			car.Structure)
	}
	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %s \n",

			car.Seats)
	}


	// Create a Bike
	bikeBuilder := &BikeBuilder{}

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1

	if motorbike.Wheels != 2{
		t.Errorf("Wheels on a Bike Must be 2 and they are %d\n", motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike"{
		t.Errorf("Structure on the Motorbike must be 'Motorbike ' and was %d\n", motorbike.Structure)
	}

	// Create a Bus

	busBuilder := &BusBuilder{}
	manufacturingComplex.SetBuilder(busBuilder)
	manufacturingComplex.Construct()

	bus := busBuilder.GetVehicle()

	if bus.Wheels!=8 {
		t.Errorf("The Buss has less Where and son cand be %d\n", bus.Wheels)
	}





}
