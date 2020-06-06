package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
	"sync"
)

// Struct of http request result
type Result struct {
	Url string
	StatusCode string
	LenBody int
	ReqTime time.Duration
}

var wg sync.WaitGroup

// Scans the STDIN
func scanInput(urls chan string, results chan Result, w *CustomCSVWriter) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		wg.Add(1)
		go PrintToStdout(results, w)
		input := scanner.Text()
		urls <- input
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner Error: %v\n", err)
	}

}

// Main program
func main() {
	start := time.Now()

	// Get the number of CPUs
	cpuNum := runtime.NumCPU()

	// Set maximum number of usable CPU cores
	runtime.GOMAXPROCS(cpuNum)

	threads := make([]int, cpuNum)

	// Handling the termination of the program
	CloseHandler(start, threads)

	// Create channels with incoming urls and request results
	urls := make(chan string)
	results := make(chan Result)

	// Create 'workers' based on number of CPU cores to handle requests
	for i := 0; i < cpuNum; i++ {
		go Request(urls, results, i, threads)
	}

	w := NewCSVWriter()
	
	// Scan the stdin and create goroutines for printing the results
	scanInput(urls, results, w)
	wg.Wait()
	close(urls)
	close(results)
	Report(threads)
}
