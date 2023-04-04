package pkg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//  Increase Count

func Count(Url string) {
	query := fmt.Sprintf("UPDATE url SET url_count = url_count + 1 WHERE short_url REGEXP '%s';", Url)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	db, err := sql.Open("mysql", Dsn())
	if err != nil {
		log.Print("unable to open db for query on main", err)
		return
	}
	defer db.Close()

	_, err = db.ExecContext(ctx, query)
	if err != nil {
		log.Print("unable to execute count query in databse ", err)
		return
	}
	log.Print(" Count Query executed successfully  ")
	return
}

// Metrics
type domain struct {
	counts int
	main   string
}

func MetricsDb() ([]domain, error) {
	query := `SELECT url_count, main_url FROM url;`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	db, err := sql.Open("mysql", Dsn())
	if err != nil {
		log.Print("unable to open db for query on main", err)
		panic(err)
	}
	defer db.Close()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Print("unable to prepare statement for query on metrics", err)
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Print("unable to execute query in db", err)
		panic(err)
	}
	defer rows.Close()

	var counts []domain
	for rows.Next() {
		var d domain
		if err := rows.Scan(&d.counts, &d.main); err != nil {
			log.Print("unable to scan row in db", err)
			panic(err)
		}
		counts = append(counts, d)
	}

	if err := rows.Err(); err != nil {
		log.Print("error in db rows", err)
		panic(err)
	}

	return counts, nil
}
