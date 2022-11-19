package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go measureCpuStats()

	http.Handle("/metrics", promhttp.Handler())
	// http.HandleFunc("/lci", lciHandler)
	http.ListenAndServe(":8080", nil)
}
