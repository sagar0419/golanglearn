package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var urlDb = make(map[string]string)
var url string
var y string

const (
	username = "root"
	password = "password"
	hostname = "127.0.0.1:3306"
	oldDb    = "db"
	dbName   = "sagar"
)

// Database service name
func dsn_old() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, oldDb)
}
func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

// Database Create
func CreateDb(dbname string) {
	db, err := sql.Open("mysql", dsn_old())
	if err != nil {
		log.Printf("Error in db connection", err)
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

// Database Table
func CreateTable() {
	db, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Print("Error While connecting DB", err)
		return
	}
	defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS url(url_id int primary key auto_increment, url_name varchar(20), url_count int);`

	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	log.Printf("Table Created")
}

// REDIRECT FUNCTION
func HomePage(w http.ResponseWriter, r *http.Request) {
	if y == "" {
		fmt.Fprintf(w, "To make a short url, enter your URL after /api?url= for eg:- http://localhost:8000/api?url=www.xyz.com \n")
	} else {
		http.Redirect(w, r, y, http.StatusSeeOther)
	}

}

// URL SHORTEN FUNCTION
func Api(w http.ResponseWriter, r *http.Request) {

	address := "http://localhost:8000/"

	url := r.URL.Query().Get("url")
	if url == "" {
		fmt.Fprintln(w, "Pass the URL")
		return
	} else {
		value, ok := urlDb[url]
		if ok == true {
			fmt.Fprintln(w, "Your short URL is", value)
			return
		}
		x := r.URL.Query()
		y = x["url"][0]

		random := rand.Int()
		str := strconv.Itoa(random)
		hash := sha256.Sum256([]byte(str))
		shortStr := fmt.Sprintf("%x", hash)[:6]
		newUrl := address + shortStr
		fmt.Fprintln(w, "Your short URL is", newUrl)
		urlDb[url] = newUrl
	}
}

// MAIN FUNCTION
func main() {

	CreateTable()
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/api/", Api)

	fmt.Println("Server is listening on localhost 8000")
	http.ListenAndServe(":8000", nil)

}
