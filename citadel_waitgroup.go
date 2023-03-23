package main

import (
	"fmt"
	"regexp"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	names := []string{
		"gopher123",
		"alpha99beta",
		"1cita2del3",
		"sag456ar",
	}

	for s := 0; s < len(names); s++ {
		// for s, names := range names {
		wg.Add(1)
		go restore(names[s])
	}

	wg.Wait()
}

func restore(names string) {
	defer wg.Done()
	// for i := 0; i < len(names); i++ {
	re := regexp.MustCompile("[0-9]+")
	// mut.Lock()
	s := re.ReplaceAllString(names, "")
	// mut.Unlock()
	fmt.Printf("%s \n", s)

}
