package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/practice/new-Code/pkg"
)

// REDIRECT FUNCTION
func HomePage(w http.ResponseWriter, r *http.Request) {

	Y := r.URL.Path
	Y = strings.TrimPrefix(Y, "/")
	if Y == "" {
		fmt.Fprintf(w, "To make a short url, enter your URL after /api?url= for eg:- http://localhost:8000/api?url=www.xyz.com \n")
		return
	} else if strings.Contains(Y, "ly") {
		s := (pkg.QueryUrlMain(Y))
		fmt.Println("On redirect function")
		fmt.Println(Y)
		pkg.Count(Y)
		http.Redirect(w, r, s, http.StatusSeeOther)
		return

	}
}
