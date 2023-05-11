package psql

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	Database string
	Port     uint32
	User     string
	Password string

	MaxConnections uint64
	AfterTimeout   uint64
}

type Psql struct {
}

func NewPsql() *pgx.ConnPool {

	username := os.Getenv("POSTGRES_USERNAME")
	if username == "" {
		username = "postgres"
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		password = "z20339"
	}

	host := os.Getenv("POSTGRES_HOST")

	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}

	portInt, err := strconv.ParseUint(port, 10, 64)

	logrus.Info("postgres host:", host)

	pgxConfig := pgx.ConnConfig{
		Host:     host, //host.docker.internal
		Database: "postgres",
		Port:     uint16(portInt),
		User:     username,
		Password: password,
	}

	/*pgxConnPoolConfig := pgx.ConnPoolConfig{
		ConnConfig:     pgxConfig,
		MaxConnections: 3,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	}*/

	conn, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     pgxConfig,
		MaxConnections: 3,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	})

	if err != nil {
		log.Fatal(err)
		panic(-1)
	}

	return conn
}

//var conn *pgx.ConnPool
//var err error

func pgxTest() {
	log.Println("start sample")

	pgxConfig := pgx.ConnConfig{
		Host:     "127.0.0.1",
		Database: "postgres",
		Port:     5432,
		User:     "postgres",
		Password: "z20339",
	}
	pgxConnPoolConfig := pgx.ConnPoolConfig{pgxConfig, 3, nil, 0}
	conn, err := pgx.NewConnPool(pgxConnPoolConfig)
	if err != nil {
		log.Fatal(err)
	}

	// declare
	var i int32
	var name string
	var updatedAt time.Time

	// select a row
	err = conn.QueryRow("select id, name from item where id = $1", 2).Scan(&i, &name)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[select a row] id: %d name: %s", i, name)

	// select rows
	rows, err := conn.Query("select id, name, updated_at from item")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&i, &name, &updatedAt); err != nil {
			log.Fatal(err)
		}
		log.Printf(
			"[select rows] id: %d name: %s updated: %s",
			i, name, updatedAt.Format("01/02 15:04:05"))
	}

	// update rows in transaction
	tx, err := conn.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	res, err := tx.Exec("UPDATE item SET updated_at = $1 WHERE id = $2", time.Now(), 2)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[update a row] updated num rows: %d", res.RowsAffected())

	err = conn.QueryRow("select name, updated_at from item where id = $1", 2).Scan(&name, &updatedAt)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[select a row in transaction] name: %s updated: %s", name, updatedAt.Format("01/02 15:04:05"))

	tx.Commit()

	err = conn.QueryRow("select name, updated_at from item where id = $1", 2).Scan(&name, &updatedAt)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[select a row] name: %s updated: %s", name, updatedAt.Format("01/02 15:04:05"))
}
