//  package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	links := []string{
// 		"https://www.google.com",
// 		"https://www.facebook.com",
// 		"https://www.youtube.com",
// 		"http://www.netflix.com",
// 		"https://www.udemy.com",
// 	}

// 	c := make(chan string)

// 	for _, link := range links {
// 		go checklink(link, c)
// 	}
// 	for l := range c {
// 		go func(link string) {
// 			time.Sleep(5 * time.Second)
// 			checklink(link, c)
// 		}(l)
// 	}
// }

// func checklink(link string, c chan string) {
// 	_, err := http.Get(link)
// 	if err != nil {
// 		fmt.Println(link, "Site is  down")
// 		c <- link
// 	}
// 	fmt.Println(link, "Link is up")
// 	c <- link
// }
