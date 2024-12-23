package pipeline

import "fmt"

type Pipeline struct {
}

func multiplyByTwo(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2 // Take n, double it, and send it to the outgoing channel
		}
		close(out)
	}()
	return out
}

func addOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + 1 // Take n, double it, and send it to the outgoing channel
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n // Take n, double it, and send it to the outgoing channel
		}
		close(out)
	}()
	return out
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Run() {
	// Start with a set of numbers
	in := make(chan int)

	// Set up the pipeline
	c1 := multiplyByTwo(in)
	c2 := addOne(c1)
	c3 := square(c2)

	// Send numbers into the pipeline and close the input channel when done
	go func() {
		for _, n := range []int{1, 4} {
			in <- n
		}
		close(in)
	}()

	// Receive the final output from the pipeline
	for result := range c3 {
		fmt.Println(result)
	}

}
