package main

import (
	"fmt"
	"net/http"

	"github.com/practice/new-Code/api"
	"github.com/practice/new-Code/pkg"
)

// MAIN FUNCTION
func main() {

	pkg.CreateDb(pkg.DbName)
	pkg.CreateTable()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", api.Api)
	mux.HandleFunc("/metrics", api.Metrics)
	mux.HandleFunc("/", api.HomePage)

	fmt.Println("Server is listening on localhost 8000")
	http.ListenAndServe(":8000", mux)
}
