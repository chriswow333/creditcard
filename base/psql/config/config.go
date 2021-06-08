package config

import (
	"fmt"
)

type dbConfig struct {
	Dialect         string
	Database        string
	User            string
	Password        string
	Host            string
	Port            int32
	Charset         string
	URL             string
	MaxIdleConns    int32
	MaxOpenConns    int32
	ConnMaxLifttime int64
	Sslmode         string
}

var DBConfig dbConfig

func initDB() {

	url := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		DBConfig.Host,
		DBConfig.Port,
		DBConfig.User,
		DBConfig.Database,
		DBConfig.Password,
		DBConfig.Sslmode,
	)

	DBConfig.URL = url
}

func init() {

	initDB()
}
