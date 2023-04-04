package pkg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Update Database table
func UpdateDb(url string, newUrl string, count int) {
	query := fmt.Sprintf("INSERT INTO url(main_url, short_url,url_count) VALUES('%s','%s', '%d');", url, newUrl, count)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	db, err := sql.Open("mysql", Dsn())
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
