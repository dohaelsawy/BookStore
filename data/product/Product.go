package data

import (
	"encoding/json"
	"io"
	"regexp"
	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int    `json:"id"`
	NAME        string `json:"name" validate:"required"`
	PublishDate string `json:"publishdate" validate:"required"`
	PRICE       int `json:"price" validate:"required"`
	SKU         string `json:"sku" validate:"required,sku"`
	DESCRIPTION string `json:"description" validate:"required"`
	CREATEDON   string `json:"createdon"`
	UPDATEDON   string `json:"updatedon"`
	AUTHOR      string `json:"author" validate:"required"`
}

type Products []*Product

var productList = []*Product{}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(),-1)
	if len(matches) != 1 {
		return false
	}
	return true
}


func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku",validateSKU)
	return validate.Struct(p)
}

func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)

}
func (p *Product) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)

}
