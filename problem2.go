package main

import (
	"sync"
	"log"
	"math/rand"
	"time"
)

// task represents the task of printing a number by a go routine
type task struct {
	id int
	slot int
	loop int
	num float32
}

// problem2 prints one random float32 every second
// for a total of taskCount passed as argument
// returns a list of timestamps for each task
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

	// limiter channel will receive a value every 1 second
	// which allows us to limit the rate of printing tasks
	limiter := time.Tick(1 * time.Second)

	ts := make([]time.Time, 0)
	for t := range tasks {
		<-limiter
		// timestamps for each task is recorded for testing the duration between tasks
		ts = append(ts, time.Now().UTC())
		log.Printf("problem2: job=%d slot=%03d count=%05d rand=%f", t.id, t.slot, t.loop, t.num)
	}

	log.Printf("problem2: finished -------------------------------------------")
	return ts
}

// printRandom2 reads jobs channel to decide if it should generate a random value
// generated value then written into the tasks channel for communication with the main process,
// along with detailed information about the task.
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
