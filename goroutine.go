package main

import (
	"fmt"
	"time"
)

func prac() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func main() {
	go prac()
	time.Sleep(1 * time.Second)
	alpha := []string{"a, b, c, d, e"}
	for _, char := range alpha {
		fmt.Printf(char)
	}
}
