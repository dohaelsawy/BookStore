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

// func GetProducts() Products {
// 	return productList
// }

// func AddProduct(p *Product) {
// 	p.ID = getNextID()
// 	productList = append(productList, p)
// }

// func getNextID() int {
// 	return (productList[len(productList)-1].ID + 1)
// }

// func UpdateProduct(id int, up *Product) error {
// 	ip, err := findIDOfProduct(id)

// 	if err != nil {
// 		return err
// 	}
// 	up.ID = id
// 	productList[ip] = up
// 	return nil
// }

// var ErrProductNotFound = fmt.Errorf("product not found")

// func findIDOfProduct(id int) (int, error) {
// 	for i, p := range productList {
// 		if p.ID == id {
// 			return i, nil
// 		}
// 	}
// 	return -1, ErrProductNotFound
// }


// func DeleteProduct(id int) error {
// 	idp , err := findIDOfProduct(id) 
// 	if err != nil {
// 		return ErrProductNotFound
// 	}
// 	productList = append(productList[:idp],productList[id+1:]... )
// 	return nil
// }

// func GetOneProduct(id int) (Product , error) {
// 	var emptyStruct Product
// 	idProd , err := findIDOfProduct(id) ;
// 	if err != nil {
// 		return emptyStruct , ErrProductNotFound
// 	} 
// 	return *productList[idProd] , nil
// }


