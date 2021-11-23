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
	log.SetLevel(log.WarnLevel)

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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "tm-calculator")
}

func Alive() bool {
	timeout := 1 * time.Second
	_, err := net.DialTimeout("tcp", "127.0.0.1:"+port, timeout)
	if err != nil {
		log.Println("Site unreachable, error: ", err)
		return false
	}
	return true
}

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

func Ready() bool {
	return true
}

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
