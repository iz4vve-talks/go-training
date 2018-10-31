package main

import (
	"os"
	"strconv"
)

func main() {
	count, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < count; i++ {
		// do nothing
	}
}
