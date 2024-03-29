package division

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/sirupsen/logrus"
)

var (
	// Define a prometheus metric for processed divisions
	divisionProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "processed_divisions_total",
		Help: "The total number of processed add events",
	})
)

// Adds Prometheus metrics to each goroutine
func divisionMetrics() {
	go func() {
		for {
			divisionProcessed.Inc()
		}
	}()
}

// IntDivision divides two integers
func IntDivision(a, b int64) int64 {
	result := a / b
	return result
}

// DivisionHandler handles the division request
func DivisionHandler(w http.ResponseWriter, r *http.Request) {

	// Get URL Params
	urlParams := r.URL.Query()
	jsonNum1, ok1 := urlParams["num1"]
	jsonNum2, ok2 := urlParams["num2"]
	if !ok1 || !ok2 {
		log.Errorf("No params found")
	}
	num1, err1 := strconv.ParseInt(jsonNum1[0], 10, 64)
	num2, err2 := strconv.ParseInt(jsonNum2[0], 10, 64)
	if err1 != nil || err2 != nil {
		log.Errorf("Could not parse provided values")
	}

	// Calculate result
	result := IntDivision(num1, num2)

	// Increment processed divisions
	divisionMetrics()

	// Set Headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Write Response
	json.NewEncoder(w).Encode(result)
	log.WithFields(log.Fields{
		"num1": jsonNum1,
		"num2": jsonNum2,
	}).Info("Division Request")
}
