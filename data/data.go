package data

import(

)

type Product struct {
	ID          string `json : "id"`
	NAME        string `json : "name"`
	AUTHOR      string `json : "author"`
	PublishDate string `json : "publishdate"`
	PRICE string       `json : "price"`
	DESCRIPTION string `json : "description"`
	CREATEDON string   `json : "-"`
	UPDATEDON string   `json : "-"`
	DELETEDON string   `json : "-"`
}


type Products []*Product

func GetProducts() Products{
	return productList
}
var productList = []*Product{
	&Product{
		ID          : "1",
		NAME      :"DESERT PHOENIX",
		AUTHOR: "Suzette Bruggeman",
		PublishDate : "March 7, 2023",
		PRICE  : "$15.00",
		DESCRIPTION : "This is a mesmerizing story that has an irresistible appeal… Desert Phoenix—Inspired by a True Story features clever plotting, focused scenes, and superior storytelling",
	},
	&Product{
		ID          : "2",
		NAME      :"All the Forgivenesses",
		AUTHOR: "Elizabeth Hardinger",
		PublishDate : "September 29, 2020",
		PRICE  : "$11.99",
		DESCRIPTION : "Set in Appalachia and the Midwest at the turn of the twentieth century and inspired by the author's family lore, this exquisite novel paints an intimately rendered portrait of one resilient farm family's challenges and hard-won triumphs--helmed by an unforgettable heroine.",
	},
}


