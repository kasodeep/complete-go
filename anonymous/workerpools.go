package anonymous

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// it takes data from channel.
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func WorkerPool() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// passing the available jobs to the workers.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// waiting for all the workers to finish.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
