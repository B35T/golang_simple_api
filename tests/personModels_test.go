package tests

import (
	"testing"
	"pop/models"
	"fmt"
)

var person = &models.PersonSchema{
	Name:"test_name",
	Email:"test@email.con",
	Password:"123445_test",
	CoverURL:"url_test",
}

func Test_InsertPerson(t *testing.T) {
	if err := person.Insert(); err != nil {
		t.Error(err)
	}
}

func Test_FindOnePerson(t *testing.T) {
	if _, err := person.FindOne(); err != nil {
		t.Error(err)
	}
}

func Test_UpdatePerson(t *testing.T) {
	person.Email = "newTest@email.com"
	if err := person.Update(); err != nil {
		t.Error(err)
	}

	if result, err := person.FindOne(); err != nil {
		t.Error(err)
	} else if result.Email != "newTest@email.com" {
		t.Error(err)
	}
}

func Test_DeletePerson(t *testing.T) {
	if err := person.Delete(); err != nil {
		t.Error(err)
	}

	if result, err := person.FindOne(); err != nil {
		t.Error(err)
	} else if result != nil {
		fmt.Println(result)
	}
}