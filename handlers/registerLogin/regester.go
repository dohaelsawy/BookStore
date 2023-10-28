package registerLogin

import (
	"net/http"
	data "github.com/dohaelsawy/bookStore/data/person"
)
type keyPerson struct{}

func (person *People) Register(rw http.ResponseWriter, r *http.Request) {
	person.l.Println("hello i'm new here i want to join .. ")
	db, err := data.DbCreateObject()
	if err != nil {
		http.Error(rw, "can't open db connection .... panic !!!",http.StatusBadRequest)
		return
	}
	defer db.Close()
	newPerson := r.Context().Value(keyPerson{}).(*data.Person)
	peopleList , err := data.GetPeopleList(db)
	for _ , p := range peopleList {
		if p.EMAIL == newPerson.EMAIL {
			http.Error(rw, "hey we have you already just jump in ...",http.StatusInternalServerError)
			return
		}
	}
	newPerson.TOKEN = "token"
	newPerson.ROLE = 0
	err = data.InsertNewPerson(db, newPerson)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	person.l.Println("new person joined :}")
}
