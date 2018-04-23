package config

const (
	db_name = "users_store"
	db_ip = "192.168.1.6:5432"
	db_username = "admin"
	db_pw = "root"
	db_type = "postgres"
)

type configDB struct {
	DB_name string
	DB_ip string
	DB_username string
	DB_pw string
	DB_type string
}


var Config_PgStore = func() *configDB {
	return &configDB{
		DB_name:db_name,
		DB_ip:db_ip,
		DB_username:db_username,
		DB_pw:db_pw,
		DB_type:db_type,
	}
}()