package models

import (
	"pop/db"
	"fmt"
	"pop/lib"
)

var newDB = db.NewDB()

type PersonSchema struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
	CoverURL string `json:"cover_url"`
	//DateCreate time.Time `json:"date_create"`
}


func (p *PersonSchema) FindAll() ([]*PersonSchema,error) {
	sql := fmt.Sprintf("SELECT Name, Email, CoverURL FROM Persons")
	row, err := newDB.DB.Query(sql)
	if err != nil {
		return nil, err
	}

	var ps []*PersonSchema

	for row.Next() {
		var pa PersonSchema
		if err = row.Scan(&pa.Name, &pa.Email, &pa.CoverURL); err != nil {
			return nil, err
		}

		ps = append(ps, &pa)
	}
	return ps, nil
}

func (p *PersonSchema) FindOne() (*PersonSchema, error) {
	sql := fmt.Sprintf("SELECT Name, Email, CoverURL FROM Persons WHERE Name='%s';", p.Name)
	row, err := newDB.DB.Query(sql)

	if err != nil {
		panic(err)
	}

	var ps PersonSchema

	for row.Next() {
		if err = row.Scan(&ps.Name, &ps.Email, &ps.CoverURL); err != nil {
			return nil,err
		}
	}

	//fmt.Println(&ps)

	return &ps ,nil
}

func (p *PersonSchema) Insert() error {
	id,_ := lib.RandomString(16)

	sql := fmt.Sprintf("INSERT INTO Persons (ID, Email, Password, Name, CoverURL)" +
		" VALUES ('%s', '%s', '%s', '%s', '%s');", id, p.Email, id, p.Name, p.CoverURL)
	_, err := newDB.DB.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}

func (p *PersonSchema) Update() error {
	sql := fmt.Sprintf("UPDATE Persons SET Email='%s' , Password='%s'" +
		"WHERE Name='%s'", p.Email, p.Password, p.Name)

	_ ,err := newDB.DB.Exec(sql)

	if err != nil {
		return err
	}
	return nil
}

func (p *PersonSchema) Delete() error {
	sql := fmt.Sprintf("DELETE FROM Persons WHERE Name='%s'", p.Name)

	_, err := newDB.DB.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}