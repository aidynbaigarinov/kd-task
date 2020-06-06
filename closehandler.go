package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// Handle the termination of the program
func CloseHandler(start time.Time, threads []int) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Printf("\nTermination of the program...\n")
		fmt.Printf("%v\n\n", time.Since(start))
		
		Report(threads)
		os.Exit(0)
	}()
}

// Prints the report after the termination:
// Thread ID | Number of handled requests
func Report(threads []int) {
	for i := range threads {
		fmt.Printf("Thread ID: %d | number of handled requests: %d\n", i, threads[i])
	}
}