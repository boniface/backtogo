package abstractfactory

type SportMotorbike struct {
}

func (s *SportMotorbike) NumWheels() int {
	return 2
}

func (s *SportMotorbike) NumSeeats() int {
	return 1

}

func (s *SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType

}
