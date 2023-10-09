package data

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

type Person struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	EMAIL       string `json:"email" validate:"required"`
	PASSWORD    string `json:"password" validate:"required"`
	PhoneNumber int    `json:"phoneNumber" validate:"required"`
	TOKEN       string `json:"token"`
	ROLE        int    `json:"role"`
	CITY        string `json:"city" validate:"required"`
}

type People []*Person

var PepoleList = []*Person{}

func (person *Person) FromJson(r io.Reader) error{
	d := json.NewDecoder(r)
	return d.Decode(person)
}

func (Person *Person) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(Person)
}

func (person *Person) Validate () error {
	validate := validator.New()
	return validate.Struct(person)
}