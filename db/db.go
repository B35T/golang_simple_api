package db

import (
	"database/sql"
	"pop/config"
	"fmt"
	_ "github.com/lib/pq"
)

//var Store = NewDB()

type PgStore struct {
	DB *sql.DB
}

func NewDB() *PgStore {
	var conf = config.Config_PgStore
	//postgres://username:password@ip:database?sslmode=disable
	str := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", conf.DB_type ,conf.DB_username, conf.DB_pw, conf.DB_ip, conf.DB_name)
	db, err := sql.Open(conf.DB_type, str)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &PgStore{db}
}

