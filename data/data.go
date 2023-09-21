package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"
	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int    `json:"id"`
	NAME        string `json:"name" validate:"required"`
	AUTHOR      string `json:"author" validate:"required"`
	PublishDate string `json:"publishdate" validate:"required"`
	PRICE       string `json:"price" validate:"required"`
	SKU         string `json:"sku" validate:"required,sku"`
	DESCRIPTION string `json:"description" validate:"required"`
	CREATEDON   string `json:"-"`
	UPDATEDON   string `json:"-"`
	DELETEDON   string `json:"-"`
}

type Products []*Product

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

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	return (productList[len(productList)-1].ID + 1)
}

func UpdateProduct(id int, up *Product) error {
	ip, err := findIDOfProduct(id)

	if err != nil {
		return err
	}
	up.ID = id
	productList[ip] = up
	return nil
}

var ErrProductNotFound = fmt.Errorf("product not found")

func findIDOfProduct(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}
	return -1, ErrProductNotFound
}


func DeleteProduct(id int) error {
	idp , err := findIDOfProduct(id) 
	if err != nil {
		return ErrProductNotFound
	}
	productList = append(productList[:idp],productList[id+1:]... )
	return nil
}

func GetOneProduct(id int) (Product , error) {
	var emptyStruct Product
	idProd , err := findIDOfProduct(id) ;
	if err != nil {
		return emptyStruct , ErrProductNotFound
	} 
	return *productList[idProd] , nil
}

var productList = []*Product{
	&Product{
		ID:          1,
		NAME:        "DESERT PHOENIX",
		AUTHOR:      "Suzette Bruggeman",
		PublishDate: "March 7, 2023",
		PRICE:       "$15.99",
		SKU :        "abc-def-ghi",
		DESCRIPTION: "This is a mesmerizing story that has an irresistible appeal… Desert Phoenix—Inspired by a True Story features clever plotting, focused scenes, and superior storytelling",
		CREATEDON:   time.Now().UTC().String(),
		UPDATEDON:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		NAME:        "All the Forgivenesses",
		AUTHOR:      "Elizabeth Hardinger",
		PublishDate: "September 29, 2020",
		PRICE:       "$11.99",
		SKU :        "abc-def-gyu",
		DESCRIPTION: "Set in Appalachia and the Midwest at the turn of the twentieth century and inspired by the author's family lore, this exquisite novel paints an intimately rendered portrait of one resilient farm family's challenges and hard-won triumphs--helmed by an unforgettable heroine.",
		CREATEDON:   time.Now().UTC().String(),
		UPDATEDON:   time.Now().UTC().String(),
	},
}
