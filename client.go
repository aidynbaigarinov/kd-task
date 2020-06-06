package main

import (
	"net"
	"net/http"
	"time"
)

// Returns a new Client with custom dial timeouts
func NewClient() *http.Client {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
	    	Timeout: 10 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	return &http.Client{
		Timeout: time.Second * 30,
		Transport: netTransport,
	}
}