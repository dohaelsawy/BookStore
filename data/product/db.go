package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbCreateObject() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/bookstore")
	if err != nil {
		return nil, err
	}
	return db, err
}

func GetProducts(db *sql.DB) (Products, error) {
	result, err := db.Query("SELECT * FROM book")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	for result.Next() {
		var p Product
		err := result.Scan(&p.ID, &p.NAME, &p.PublishDate, &p.PRICE, &p.SKU, &p.DESCRIPTION, &p.CREATEDON, &p.UPDATEDON, &p.AUTHOR)
		if err != nil {
			return nil, err
		}
		productList = append(productList, &p)
	}

	return productList, nil
}

func GetOneProduct(db *sql.DB, id int) (Product, error) {
	var p Product
	result, err := db.Query("SELECT * FROM book WHERE book_id = ?", id)
	if err != nil {
		return p, err
	}
	defer db.Close()
	for result.Next() {
		err = result.Scan(&p.ID, &p.NAME, &p.PublishDate, &p.PRICE, &p.SKU, &p.DESCRIPTION, &p.CREATEDON, &p.UPDATEDON, &p.AUTHOR)
		if err != nil {
			return p, err
		}
	}
	return p, nil
}

func InsertProduct(db *sql.DB , product *Product)  error{
	prod , err := db.Prepare("INSERT INTO book(name,publish_date,price,sku,description,created_on,updated_on,author) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer prod.Close()
	_, err = prod.Exec(&product.NAME,&product.PublishDate,&product.PRICE,&product.SKU,&product.DESCRIPTION,&product.CREATEDON,&product.UPDATEDON,&product.AUTHOR)
	if err != nil {
		return err
	}
	return nil
}


func DeleteProduct(db *sql.DB , id int) error {
	stat , err := db.Prepare("DELETE FROM book WHERE book_id = ?")
	if err != nil {
		return err
	}
	_ , err = stat.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct (db *sql.DB , id int , prodcut *Product) error {
	stat , err := db.Prepare("UPDATE book SET name = ?,publish_date = ? ,price = ? , sku = ? , description = ? , created_on = ? , updated_on = ? , author = ?  WHERE book_id = ?")
	if err != nil {
		return err
	}
	_ , err = stat.Exec(prodcut.NAME , prodcut.PublishDate , prodcut.PRICE , prodcut.SKU,prodcut.DESCRIPTION,prodcut.CREATEDON , prodcut.UPDATEDON , prodcut.AUTHOR , id)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}