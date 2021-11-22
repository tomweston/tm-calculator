package random

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

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
	json_num, ok := urlParams["num"]
	if !ok {
		log.Infoln("No params found, sending 10 random numbers")

		// TODO: send 10 random numbers
		json_num := []string{"10"}
		json_num_string := strings.Join(json_num, " ")
		if defaultNum, err := strconv.Atoi(json_num_string); err == nil {
			for i := 0; i < defaultNum; i++ {
				result := IntRandom()
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				json.NewEncoder(w).Encode(result)
			}
		}
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

	for i := 0; i < intNum; i++ {
		result := IntRandom()
		json.NewEncoder(w).Encode(result)
	}
}
