package main

import (
	"sync"
	"log"
	"math/rand"
	"time"
)

type task struct {
	id int
	slot int
	loop int
	num float32
}

func problem2(taskCount int) []time.Time {

	log.Printf("problem2: started --------------------------------------------")

	jobs := make(chan int, taskCount)
	tasks := make(chan task, taskCount)
	for j := 1; j <= taskCount; j++ {
		jobs <- j
	}
	close(jobs)

	var wg sync.WaitGroup
	wg.Add(taskCount)

	for inx := 0; inx < 10; inx++ {

		go printRandom2(inx, &wg, jobs, tasks)

	}

	wg.Wait()
	close(tasks)

	limiter := time.Tick(1 * time.Second)

	ts := make([]time.Time, 0)
	for t := range tasks {
		<-limiter
		ts = append(ts, time.Now().UTC())
		log.Printf("ts:%v", time.Now().UTC())
		log.Printf("problem2: job=%d slot=%03d count=%05d rand=%f", t.id, t.slot, t.loop, t.num)
	}

	log.Printf("problem2: finished -------------------------------------------")
	return ts
}

func printRandom2(slot int, wg *sync.WaitGroup, jobs <-chan int, tasks chan<- task) {

	for inx := 0; inx < 10; inx++ {

		select {
		case j, ok := <-jobs:
			if ok {
				t := task{
					id: j,
					slot: slot,
					loop: inx,
					num: rand.Float32(),
				}
				tasks <- t
				wg.Done()
			} else {
				return
			}
		default:
			log.Printf("no value at slot=%03d, moving on", slot)
		}
	}
}
