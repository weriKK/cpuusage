package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/promauto"
// )

// var (
// 	lciLoad = promauto.NewGauge(prometheus.GaugeOpts{
// 		Name: "myapp_cpu_usage_percent_sec",
// 		Help: "Measured CPU usage percentage per second",
// 	})
// )

// func lciHandler(w http.ResponseWriter, r *http.Request) {
// 	lciHdr := r.Header.Get("Lci")
// 	if lciHdr == "" {
// 		http.Error(w, "empty Lci header", http.StatusBadRequest)
// 		return
// 	}

// 	parameters := strings.Split(lciHdr, ";")
// 	if len(parameters) != 3 {
// 		http.Error(w, "invalid Lci header", http.StatusBadRequest)
// 		return
// 	}

// 	lci := map[string]string{}
// 	for _, p := range parameters {
// 		kv := strings.SplitN(p, ":", 2)
// 		if len(kv) != 2 || kv[0] == "" || kv[1] == "" {
// 			http.Error(w, "invalid Lci header parameter", http.StatusBadRequest)
// 			return
// 		}

// 		lci[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
// 	}

// 	lciLoad.Set()

// 	fmt.Printf("%+v\n", lci)
// }
