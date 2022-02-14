package main

import (
	"fmt"
)

const amountOfWorkers = 20

type work struct {
	id string
}

type worker struct {
	id       int
	jobChan  <-chan work
	doneChan chan<- bool
}

func main() {
	jobChan := make(chan work, amountOfWorkers)
	doneChan := make(chan bool)

	for i := 0; i < amountOfWorkers; i++ {
		go worker{i, jobChan, doneChan}.start()
	}

	workToBedone := []work{{"a"}, {"b"}, {"c"}}
	for _, workToDo := range workToBedone {
		jobChan <- workToDo
	}

	for {
		if len(jobChan) == 0 {
			close(jobChan)
			break
		}
	}

	for range make([]bool, amountOfWorkers) {
		<-doneChan
	}
}

func (w worker) start() {
	defer func() {
		w.doneChan <- true
	}()

	for workReceived := range w.jobChan {
		fmt.Println("work done ", w.id, workReceived)
	}
}
