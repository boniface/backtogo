package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 :=GetInstance()

	if counter1 == nil {
		t.Error("Expected Pointer to Singleton after calling getInstance")

	}

	expectedCounter :=counter1

	currentCount := counter1.AddOne()

	if currentCount != 1 {
		t.Error("After Calling the First Time must be 1 but it is  %d\n",currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter{
		t.Error("Expected Same instance in Counter 2 nbut it got a different instance")

	}


}
