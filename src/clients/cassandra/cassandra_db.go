package cassandra

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	mysql_users_usersname = "mysql_users_username"
	mysql_users_passwords = "mysql_users_password"
	mysql_users_host      = "mysql_users_host"
	mysql_users_db        = "mysql_users_db"
)

var (
	Client   *sql.DB
	username = "root"
	password = ""
	host     = "localhost"
	db       = "cappuccino"
)

func init() {

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		db,
	)
	log.Println("about to connect to %s", datasourceName)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("connected to database")

}
