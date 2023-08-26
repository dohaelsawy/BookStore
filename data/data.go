package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int    `json:"id"`
	NAME        string `json:"name"`
	AUTHOR      string `json:"author"`
	PublishDate string `json:"publishdate"`
	PRICE       string `json:"price"`
	DESCRIPTION string `json:"description"`
	CREATEDON   string `json:"-"`
	UPDATEDON   string `json:"-"`
	DELETEDON   string `json:"-"`
}

type Products []*Product

func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Products) ToJson(w io.Writer) error {
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

var productList = []*Product{
	&Product{
		ID:          1,
		NAME:        "DESERT PHOENIX",
		AUTHOR:      "Suzette Bruggeman",
		PublishDate: "March 7, 2023",
		PRICE:       "$15.99",
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
		DESCRIPTION: "Set in Appalachia and the Midwest at the turn of the twentieth century and inspired by the author's family lore, this exquisite novel paints an intimately rendered portrait of one resilient farm family's challenges and hard-won triumphs--helmed by an unforgettable heroine.",
	},
}
