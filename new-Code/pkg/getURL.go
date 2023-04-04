package pkg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Query DB for short
func QueryUrl(url string) string {
	query := fmt.Sprintf("SELECT short_url FROM url WHERE main_url REGEXP '%s';", url)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	db, err := sql.Open("mysql", Dsn())
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
