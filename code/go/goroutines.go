package main

import (
	"fmt"
	"time"
)

func main() {
	// START MAIN OMIT
	// in a main()
	for i := 0; i < 10; i++ {
		go worker(i, 500)
	}
	// END MAIN OMIT
	time.Sleep(3 * time.Second) // wait for goroutines to be executed
}

// START FUNC OMIT
func worker(id, n int) {
	for n >= 0 {
		time.Sleep(time.Millisecond)
		n--
	}
	fmt.Printf("%d: done!\n", id)
}

// END FUNC OMIT
