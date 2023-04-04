package pkg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Query DB for Main URL
func QueryUrlMain(url string) string {
	query := fmt.Sprintf("SELECT  main_url FROM url WHERE short_url REGEXP '%s';", url)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	db, err := sql.Open("mysql", Dsn())
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
