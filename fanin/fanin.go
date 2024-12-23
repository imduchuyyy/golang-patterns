package fanin

import (
	"fmt"
	"sync"
)

type FanIn struct{}

func fanIn(inps ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	// Start a goroutine for each input channel.
	wg.Add(len(inps))
	for _, c := range inps {
		go output(c)
	}

	// Start a goroutine to close the output channel// once all input inps have been drained.
	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func NewFanIn() *FanIn {
	return &FanIn{}
}

func (f *FanIn) Run() {
	c1 := generator(1, 3, 5)
	c2 := generator(2, 4, 6)

	c := fanIn(c1, c2)

	for n := range c {
		fmt.Println(n)
	}
}
