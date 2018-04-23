package controllers

import (
	"net/http"
	"pop/models"
	"encoding/json"
	"errors"
	"fmt"
)


type PersonManager struct {
	person *models.PersonSchema
}


func getWithTypeJson(r *http.Request) (*models.PersonSchema, error) {
	if r.Header.Get("Content-Type") == "application/json" {
		var p *models.PersonSchema
		var decode = json.NewDecoder(r.Body)

		err := decode.Decode(&p)
		if err != nil {
			return nil, err
		}

		return p, nil
	}
	return nil, errors.New("get json error")
}


func FindWithName(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		name := r.FormValue("name")

		manager := PersonManager{
			person: &models.PersonSchema{Name:name},
		}

		if result, err := manager.person.FindOne(); err != nil {
			w.WriteHeader(http.StatusNoContent)
		} else {
			json, err := json.Marshal(result)
			if err != nil {
				w.WriteHeader(http.StatusNoContent)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(json))
		}
		break
	default:
		w.WriteHeader(http.StatusNoContent)
		break
	}
}

func FindAllPerson(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		manager := PersonManager{}
		if result, err := manager.person.FindAll(); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNoContent)
		} else {
			json, err := json.Marshal(result)
			if err != nil {
				w.WriteHeader(http.StatusNoContent)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(json))
		}
		break
	default:
		w.WriteHeader(http.StatusNoContent)
		break
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		person, err := getWithTypeJson(r)
		if err != nil {
			w.WriteHeader(http.StatusFound)
		}

		manager := PersonManager{person}

		if err = manager.person.Insert(); err != nil {
			w.WriteHeader(http.StatusFound)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		break
	default:
		w.WriteHeader(http.StatusFound)
		break
	}
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		person , err := getWithTypeJson(r)
		if err != nil {
			w.WriteHeader(http.StatusFound)
		}

		manager := PersonManager{person}

		if err = manager.person.Update(); err != nil {
			w.WriteHeader(http.StatusFound)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		break
	default:
		w.WriteHeader(http.StatusFound)
		break
	}
}

