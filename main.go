package main

import (
	"log"
)

func main() {
	log.Printf("random numbers printed: %d", problem1())
	ts := problem2(5)
	log.Println("timestamps for problem2:")
	for _, v := range ts {
		log.Printf("%v", v)
	}
}
