package models

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name: "Coffee",
		Price: 1.00,
		SKU: "abs-abc-def",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}