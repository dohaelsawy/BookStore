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

func GetPeopleList(db *sql.DB) (People, error) {
	result, err := db.Query("SELECT * FROM person")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	for result.Next() {
		var p Person
		err := result.Scan(&p.ID, &p.FirstName, &p.LastName, &p.EMAIL, &p.PASSWORD, &p.PhoneNumber, &p.CITY, &p.ROLE, &p.TOKEN)
		if err != nil {
			return nil, err
		}
		PepoleList = append(PepoleList, &p)
	}

	return PepoleList, nil
}

func GetOnePerson(db *sql.DB, id int) (Person, error) {
	var p Person
	result, err := db.Query("SELECT * FROM book WHERE book_id = ?", id)
	if err != nil {
		return p, err
	}
	defer db.Close()
	for result.Next() {
		err = result.Scan(&p.ID, &p.FirstName, &p.LastName, &p.EMAIL, &p.PASSWORD, &p.PhoneNumber, &p.CITY, &p.ROLE, &p.TOKEN)
		if err != nil {
			return p, err
		}
	}
	return p, nil
}

func InsertNewPerson(db *sql.DB , person *Person)  error{
	prod , err := db.Prepare("INSERT INTO person(first_name,last_name,email,password,phone_number,token,role,city) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer prod.Close()
	_, err = prod.Exec(&person.FirstName,&person.LastName,&person.EMAIL,&person.PASSWORD,&person.PhoneNumber,&person.TOKEN,&person.ROLE,&person.CITY)
	if err != nil {
		return err
	}
	return nil
}


func DeleteOnePerson(db *sql.DB , id int) error {
	stat , err := db.Prepare("DELETE FROM person WHERE person_id = ?")
	if err != nil {
		return err
	}
	_ , err = stat.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePerson(db *sql.DB , id int , person *Person) error {
	stat , err := db.Prepare("UPDATE person SET first_name = ?,last_name = ? ,email = ? , password = ? , phone_number = ? , token = ? , role = ? , city = ?  WHERE person_id = ?")
	if err != nil {
		return err
	}
	_ , err = stat.Exec(person.FirstName , person.LastName , person.EMAIL , person.PASSWORD,person.PhoneNumber,person.TOKEN , person.ROLE , person.CITY , id)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}