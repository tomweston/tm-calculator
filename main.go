package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/sirupsen/logrus"
	"github.com/tomweston/tm-calculator/pkg/add"
	"github.com/tomweston/tm-calculator/pkg/division"
	"github.com/tomweston/tm-calculator/pkg/random"
	"github.com/tomweston/tm-calculator/pkg/subtract"
)

var (
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_duration_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})

	port string = "5555"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

}

// prometheusMiddleware
func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}

// IndexHandler	handles the root route
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "tm-calculator")
}

// Alive returns true if the server is alive
func Alive() bool {
	_, err := net.DialTimeout("tcp", "localhost:"+port, time.Second)
	if err != nil {
		return false
	}
	return true
}

// LivenessHandler handles the liveness route
func LivenessHandler(w http.ResponseWriter, r *http.Request) {
	if Alive() == true {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"ready": true}`)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"ready": false}`)
	}
}

// Ready returns true if the server is ready
func Ready() bool {
	return true
}

// ReadinessHandler handles the readiness route
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	if Alive() == true && Ready() == true {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"ready": true}`)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"ready": false}`)
	}
}

func handleRequests() {
	r := mux.NewRouter()
	r.Use(prometheusMiddleware)
	r.Path("/metrics").Handler(promhttp.Handler())

	// Default Routes
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/liveness", LivenessHandler)
	r.HandleFunc("/readiness", ReadinessHandler)

	// Use a Mux Subrouter to route calls to corresponding API versions
	v1 := r.PathPrefix("/api/v1/").Subrouter()
	v1.HandleFunc("/add", add.AddHandler)
	v1.HandleFunc("/subtract", subtract.SubtractHandler)
	v1.HandleFunc("/division", division.DivisionHandler)
	v1.HandleFunc("/random", random.RandomHandler)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func main() {
	handleRequests()
}
