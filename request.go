package main

import (
	"io/ioutil"
	"log"
	"time"
)

// Makes request to given URL
func Request(urls <-chan string, results chan<- Result, i int, threads []int) {

	netClient := NewClient()

	// Get url from the Urls channel
	for input := range urls {
		start := time.Now()

		// Make GET request
		response, err := netClient.Get(input)
		if err != nil {
			log.Printf("Couldn't make a request to %v | %v\n", input, err)
			return
		}
		defer response.Body.Close()
		
		// Time spent on a GET request
		elapsed := time.Since(start)
		
		// Increment number of handled requests
		threads[i]++

		// Read the body of response
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("Couldn't read body of response | %v\n", err)
			return
		}

		// Make new `result` variable and fill it with the response data
		result := Result{
			Url: input,
			StatusCode: response.Status,
			LenBody: len(body),
			ReqTime: elapsed,
		}

		// Send result to channel
		results <- result
	}
}
