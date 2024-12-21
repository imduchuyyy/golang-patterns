# Golang Patterns 
For studying purposes, I'm going to implement some patterns in Golang.

## Patterns:
1. Generator: 
The generator pattern is a way to create a function that gives us a series of values. In other programming languages, people might use things called enumerators or iterators. But in Go, the best tool for this job is a channel.

Think of it this way: we have a generator function that gives us something like a list we can pull values from, one at a time.

When we get values from this list, we can use a for range in our code to handle each value as it comes.

2. Worker pool:
In the worker pool pattern, we use multiple workers to handle jobs. Each worker is a separate process, and we have a way of assigning tasks to these workers. In Go, we use goroutines for the workers and channels to give them tasks and to receive the results back.

A worker is just a function that runs in its own goroutine and it's designed to do jobs and usually, a worker will keep running, always ready to pick up a new job from a channel.

When it gets a job, it does the work needed and then gets ready for the next job.
