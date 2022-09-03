package main

import (
	"fmt"
	"time"
)

type ControlMsg int

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

const (
	DoExit = iota
	ExitOk
)

func doubler(jobs <-chan Job, results chan<- Result, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exit goroutine")
				control <- ExitOk
				return
			default:
				panic("Unhandled control message")
			}
		case job := <-jobs:
			results <- Result{result: job.data * 2, job: job}
		}
	}
}

func main() {
	jobs := make(chan Job, 50)
	results := make(chan Result, 50)
	control := make(chan ControlMsg)

	go doubler(jobs, results, control)

	for i := 0; i < 30; i++ {
		jobs <- Job{i}
	}

	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timed out")
			control <- DoExit
			// Wait for response from goroutine before exit.
			<-control
			fmt.Println("program exit")
			return
		}
	}
}
