package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

var urlDb = make(map[string]string)
var url string

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "To make a short url, enter your URL after /api?url= for eg:- http://localhost:8000/api?url=www.xyz.com \n")
}

func LandingUrl(w http.ResponseWriter, r *http.Request) {

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

		random := rand.Int()
		str := strconv.Itoa(random)
		hash := sha256.Sum256([]byte(str))
		shortStr := fmt.Sprintf("%x", hash)[:6]
		newUrl := address + shortStr
		fmt.Fprintln(w, "Your short URL is", newUrl)
		urlDb[url] = newUrl

	}
}

func main() {

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/api/", LandingUrl)

	fmt.Println("Server is listening on localhost 8000")
	http.ListenAndServe(":8000", nil)
}
