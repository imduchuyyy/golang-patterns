package generator

import "log"

type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{}
}

func evenGenerator() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; ; i += 2 {
			c <- i
		}
	}()

	return c
}

func (g *Generator) Run() {
	log.Println(`
	The generator pattern is a way to create a function that gives us a series of values. 
	In other programming languages, people might use things called enumerators or iterators. 
	But in Go, the best tool for this job is a channel.

	Think of it this way: we have a generator function that gives us something like a list we can pull values from, one at a time.

	When we get values from this list, we can use a for range in our code to handle each value as it comes.

	Remember, though, don't make things more complicated than they need to be. If we're working with a small group of items that our computer can handle easily, then just putting those items in a list and giving back that list is perfectly fine.
	`)

	for v := range evenGenerator() {
		log.Println(v)
	}
}
