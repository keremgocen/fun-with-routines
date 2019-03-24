package main

import (
	"log"
)

func main() {
	// log.Printf("random numbers printed: %d", problem1())
	ts := problem2(5)
	for i, v := range ts {
		log.Printf("time:%v", v)
		if i > 0 && i < 5 {
			delta := ts[i].Sub(ts[i-1])
			log.Println(delta.Seconds())
			if delta.Seconds() > 0.9 {
				log.Println("OK")
			}
		}
	}
}
