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
var y string

func HomePage(w http.ResponseWriter, r *http.Request) {
	if y == "" {
		fmt.Fprintf(w, "To make a short url, enter your URL after /api?url= for eg:- http://localhost:8000/api?url=www.xyz.com \n")
	} else {
		http.Redirect(w, r, y, http.StatusSeeOther)
	}

}

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

func main() {

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/api/", Api)

	fmt.Println("Server is listening on localhost 8000")
	http.ListenAndServe(":8000", nil)
}
