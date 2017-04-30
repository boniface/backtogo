package factorymethod

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of Type Cash Must Exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error(" The Cash payment Method was not correct ")
	}

	t.Log("LOG:", msg)

}

func TestPaymentMethodDebicard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A payment method of Type 'DebitCard' must exist")
	}

	msg := payment.Pay(22.30)

	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The Debit Card Payment Method Was not Correct")

	}

	t.Log("LOG:", msg)
}
