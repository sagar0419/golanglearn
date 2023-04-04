package main

import (
	"fmt"
	"time"
)

func prac(a, b int) {
	for i := a; i <= b; i++ {
		fmt.Println(i)
	}
}

func main() {
	a := 2
	b := 6
	go prac(a, b)
	time.Sleep(6 * time.Second)
	for i := a; i <= b; i++ {
		fmt.Printf("%c \n", 'A'+i-1)
	}
}
