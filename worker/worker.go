package worker

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Worker struct {
}

func NewWorker() *Worker {
	return &Worker{}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)

		time.Sleep(time.Second) // simulate time-consuming task
		results <- j * j        // square the number and send the result back

		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func (w *Worker) Run() {
	log.Println(`
	In the worker pool pattern, we use multiple workers to handle jobs. Each worker is a separate process, and we have a way of assigning tasks to these workers. In Go, we use goroutines for the workers and channels to give them tasks and to receive the results back.
	A worker is just a function that runs in its own goroutine and it's designed to do jobs and usually, a worker will keep running, always ready to pick up a new job from a channel.
	When it gets a job, it does the work needed and then gets ready for the next job:
	`)

	numJobs := 10
	numWorkers := 3
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	wg := sync.WaitGroup{}

	// Start the workers.
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs to the workers and close the jobs channel when done.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all the workers to finish.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and print the results.
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}

}
