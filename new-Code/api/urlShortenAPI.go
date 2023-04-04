package api

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/practice/new-Code/pkg"
)

// URL SHORTEN FUNCTION
func Api(w http.ResponseWriter, r *http.Request) {

	address := "http://localhost:8000/"

	Url := r.URL.Query().Get("url")
	if Url == "" {
		fmt.Fprintf(w, "Pass the URL")
		return
	} else {
		s := (pkg.QueryUrl(Url))

		if s != "" {
			fmt.Fprintln(w, "Hey you have already shorten your URL, your short URL is", s)
			return
		}

		x := r.URL.Query()
		Y = x["url"][0]

		random := rand.Int()
		str := strconv.Itoa(random)
		hash := sha256.Sum256([]byte(str))
		shortStr := fmt.Sprintf("%x", hash)[:6]
		NewUrl = address + "ly." + shortStr
		fmt.Fprintln(w, "Your short URL is", NewUrl)
		UrlDb[Url] = NewUrl
		pkg.UpdateDb(Url, NewUrl, Count)
	}
}
