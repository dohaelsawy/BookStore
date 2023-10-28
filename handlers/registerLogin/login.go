package registerLogin

import (
	"net/http"

	data "github.com/dohaelsawy/bookStore/data/person"
)

func (person *People) login(rw http.ResponseWriter, r *http.Request) {
	person.l.Println("hello let me in ...")

	db, err := data.DbCreateObject()
	if err != nil {
		http.Error(rw, "can't open db connection .... panic !!!", http.StatusBadRequest)
		return
	}
	defer db.Close()

	personSite := r.Context().Value(keyPerson{}).(*data.Person)
	personDB, err := data.GetOnePerson(db, personSite.EMAIL)
	if err != nil {
		http.Error(rw, "you have to register first .. don't panic !!!", http.StatusBadRequest)
		return
	}
	err = personDB.ToJson(rw)
	if err != nil {
		http.Error(rw, "can't write it in rw .. panic !!!", http.StatusInternalServerError)
		return
	}
	person.l.Println("person logged in :}")
}
