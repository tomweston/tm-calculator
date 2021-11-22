package random

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/sirupsen/logrus"
)

var (
	// Define a prometheus metric for processed random events
	randomProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "processed_random_total",
		Help: "The total number of processed random number events",
	})
)

// Adds Prometheus metrics to each goroutine
func randomMetrics() {
	go func() {
		for {
			randomProcessed.Inc()
		}
	}()
}

// Function to generate a defined set of random numbers
func IntRandom() int {
	return rand.Int()
}

func RandomHandler(w http.ResponseWriter, r *http.Request) {

	// Get URL Params
	urlParams := r.URL.Query()
	json_num, ok1 := urlParams["num"]
	if !ok1 {
		log.Errorf("No params found")
	}
	num, err1 := strconv.ParseInt(json_num[0], 10, 64)
	if err1 != nil {
		log.Errorf("Could not parse provided value")
	}

	// Increment processed random events
	randomMetrics()

	// Set Headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Write Response
	intNum := int(num)
	if num != 0 {
		for i := 0; i < intNum; i++ {
			result := IntRandom()
			json.NewEncoder(w).Encode(result)
		}
	} else {
		for i := 0; i < 10; i++ {
			result := IntRandom()
			json.NewEncoder(w).Encode(result)
		}
	}
}
