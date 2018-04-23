package tests

import (
	"testing"
	"golang_simple_api/config"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"golang_simple_api/db"
)

var newDB = db.NewDB()

func Test_DBConnectAndPing(t *testing.T) {
	var conf = config.Config_PgStore
	//postgres://username:password@ip:database?sslmode=disable
	str := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", conf.DB_type ,conf.DB_username, conf.DB_pw, conf.DB_ip, conf.DB_name)
	db, err := sql.Open(conf.DB_type, str)

	defer db.Close()

	if err != nil {
		t.Error(err)
	}

	err = db.Ping()
	if err != nil {
		t.Error(err)
	}
}