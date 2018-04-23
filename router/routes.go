package router

import (
	"net/http"
	c "golang_simple_api/controllers"
)

func Init() {

	//person
	http.HandleFunc("/person/all", c.FindAllPerson)
	http.HandleFunc("/person/", c.FindWithName) // person/?name=MyName
	http.HandleFunc("/person/register", c.Register)
	http.HandleFunc("/person/update", c.UpdateProfile)

	// listen
	http.ListenAndServe(":3000",nil)
}
