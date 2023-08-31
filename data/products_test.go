package data

import "testing"

func  CheckingValidationProduct (t *testing.T){
	p := &Product{
		AUTHOR:      "Suzette Bruggeman",
		PublishDate: "March 7, 2023",
		SKU: "abc",
		PRICE:       "$15.99",
		DESCRIPTION: "This is a mesmerizing story that has an irresistible appeal… Desert Phoenix—Inspired by a True Story features clever plotting, focused scenes, and superior storytelling",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}