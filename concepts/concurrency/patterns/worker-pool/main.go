package main

import (
	"fmt"
	"time"
)

// A worker listening for jobs continually.
// takes a job -> wait 1 second -> sends result back.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started\tjob", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished\tjob", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
