package main

import "strconv"

// Get the data from `results` channel and write it to STDOUT with Custom CSV Writer
func PrintToStdout(results <-chan Result, w *CustomCSVWriter) {
	defer w.Flush()
	result := <- results
	data := []string{
		result.Url, 
		result.StatusCode, 
		strconv.Itoa(result.LenBody), 
		result.ReqTime.String(),
	}

	w.Write(data)
	wg.Done()
}
