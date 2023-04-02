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
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var urlDb = make(map[string]string)
var url string
var y string
var newUrl string
var count int = 1

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

	db, err := sql.Open("mysql", dsn())
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

// Update Database table
func UpdateDb(url string, newUrl string, count int) {
	query := fmt.Sprintf("INSERT INTO url(main_url, short_url,url_count) VALUES('%s','%s', '%d');", url, newUrl, count)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	db, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Print("Error while connecting DB to push the data in the table", err)
		return
	}
	defer db.Close()

	_, err = db.ExecContext(ctx, query)
	if err != nil {
		log.Print("error while pushing the data into the table", err)
		return
	}
	log.Printf("Data is inserted in the table")
}

// Query DB for short
func QueryUrl(url string) string {
	query := fmt.Sprintf("SELECT short_url FROM url WHERE main_url REGEXP '%s';", url)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	db, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Print("unable to open db fro query", err)
		return ""
	}
	defer db.Close()

	var short_url string
	err = db.QueryRowContext(ctx, query).Scan(&short_url)
	if err != nil {
		log.Print("unable to execute query in db on short", err)
		return ""
	}
	log.Print("Query executed successfully on short")
	return short_url
}

// Query DB for Main URL
func QueryUrlMain(url string) string {
	query := fmt.Sprintf("SELECT  main_url FROM url WHERE short_url REGEXP '%s';", url)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	db, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Print("unable to open db for query on main", err)
		return err.Error()
	}
	defer db.Close()

	var main_url string

	err = db.QueryRowContext(ctx, query).Scan(&main_url)
	if err != nil {
		log.Print("unable to execute query in db", err)
		return err.Error()
	}
	log.Print("Query executed successfully on main ", main_url)
	return main_url
}

//  Increase Count

func Count(url string) {
	query := fmt.Sprintf("UPDATE url SET url_count = url_count + 1 WHERE short_url REGEXP '%s';", url)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	db, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Print("unable to open db for query on main", err)
		return
	}
	defer db.Close()

	var count string

	err = db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		log.Print("unable to execute query in db", err)
		return
	}
	log.Print("Query executed successfully on main ", count)
	return
}

// REDIRECT FUNCTION
func HomePage(w http.ResponseWriter, r *http.Request) {

	y := r.URL.Path
	if y == "" {
		fmt.Fprintf(w, "To make a short url, enter your URL after /api?url= for eg:- http://localhost:8000/api?url=www.xyz.com \n")
		return
	} else {
		if strings.Contains(y, "/api/") {
			fmt.Println(y)
			return
		} else {
			s := (QueryUrlMain(y))
			fmt.Println("jai ho")
			http.Redirect(w, r, s, http.StatusSeeOther)
			Count(y)
			return
		}
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
		s := (QueryUrl(url))

		if s != "" {
			fmt.Fprintln(w, "Hey you have already shorten your URL, your short URL is", s)
			return
		}

		x := r.URL.Query()
		y = x["url"][0]

		random := rand.Int()
		str := strconv.Itoa(random)
		hash := sha256.Sum256([]byte(str))
		shortStr := fmt.Sprintf("%x", hash)[:6]
		newUrl = address + shortStr
		fmt.Fprintln(w, "Your short URL is", newUrl)
		urlDb[url] = newUrl
		fmt.Println(url)
		fmt.Println(newUrl)
		fmt.Println(y)
		UpdateDb(url, newUrl, count)
	}
}

// MAIN FUNCTION
func main() {
	CreateDb(dbName)
	CreateTable()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", Api)
	mux.HandleFunc("/", HomePage)

	fmt.Println("Server is listening on localhost 8000")
	http.ListenAndServe(":8000", mux)
}
