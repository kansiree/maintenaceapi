package manageMysql

import (
	"database/sql"
	"testing"
)

var Driver = "mysql"
var Port = "3306"
var DataBase = "t016ffukzsi0y5ie"
var Host = "u3r5w4ayhxzdrw87.cbetxkdyhwsb.us-east-1.rds.amazonaws.com"
var User = "sz0debklevf8wjhf"
var Password = "gu2af8swu50tjc3k"

func ConnectDB() (db *sql.DB, err error) {
	// db, err = sql.Open(Driver, "sz0debklevf8wjhf:gu2af8swu50tjc3k@tcp(u3r5w4ayhxzdrw87.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/t016ffukzsi0y5ie")

	db, err = sql.Open(Driver, getDatabaseURL(Driver, User, Password, Port, Host, DataBase))
	if err != nil {
		panic(err.Error())
	}
	return db, err
}

func getDatabaseURL(driver string, user string, password string, port string, host string, databaeName string) string {
	if driver == "mysql" {
		url := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaeName
		return url
	} else {
		return "not support database type"
	}

}

func TestConnectDB(t *testing.T) {
	ConnectDB()
}
