package main

import (
	"fmt"
	"regexp"
)

func main() {
	names := [][]string{
		{"gopher123"},
		{"alpha99beta"},
		{"1cita2del3"},
	}

	rchan := make(chan string)

	for _, str := range names {
		go restore(str, rchan)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(<-rchan)
	}
}

func restore(name []string, rchan chan string) {
	re := regexp.MustCompile("[0-9]+")
	for _, s := range name {
		s = re.ReplaceAllString(s, "")
		rchan <- s
	}
}
