package main

import (
	"flag"
)

func main() {
	count := flag.Int("i", 100000000, "count")
	flag.Parse()

	for i := 0; i < *count; i++ {

	}
}
