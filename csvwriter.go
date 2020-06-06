package main

import (
	"encoding/csv"
	"os"
	"sync"
)

type CustomCSVWriter struct {
	mutex *sync.Mutex
	csvWriter *csv.Writer
}

// Returns Custom CSV Writer
func NewCSVWriter() *CustomCSVWriter {
	w := csv.NewWriter(os.Stdout)
	return &CustomCSVWriter{csvWriter: w, mutex: &sync.Mutex{}}
}

// Writes to STDOUT
func (w *CustomCSVWriter) Write(row []string) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.csvWriter.Write(row)
}

// Flushes the data
func (w *CustomCSVWriter) Flush() error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.csvWriter.Flush()
	return w.csvWriter.Error()
}