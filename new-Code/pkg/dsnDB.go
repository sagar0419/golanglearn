package pkg

import "fmt"

const (
	Username = "root"
	Password = "password"
	Hostname = "127.0.0.1:3306"
	OldDb    = "db"
	DbName   = "sagar"
)

// Database service name
func Dsn_old() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", Username, Password, Hostname, OldDb)
}
func Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", Username, Password, Hostname, DbName)
}
