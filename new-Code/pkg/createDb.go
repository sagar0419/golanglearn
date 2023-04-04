package pkg

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Database Create
func CreateDb(dbname string) {
	db, err := sql.Open("mysql", Dsn_old())
	if err != nil {
		log.Print("Error in db connection", err)
		return
	}

	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return
	}
	log.Printf("rows affected %d\n", no)

}

// Create Database Table
func CreateTable() {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	db, err := sql.Open("mysql", Dsn())
	if err != nil {
		log.Print("Error While connecting DB", err)
		return
	}
	defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS url(id int primary key auto_increment, main_url varchar(2048) , short_url varchar(200) , url_count int );`

	// _, err = db.Exec(query)
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		log.Print(err)
	}
	log.Printf("Table Created")
}
