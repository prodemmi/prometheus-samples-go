package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	userMetrics := RegisterUserMetrics()
	go SetupUserSeeder(userMetrics)

	promHandler := promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{})

	pMux := http.NewServeMux()
	pMux.Handle("/metrics", promHandler)
	pMux.HandleFunc("/ping", func(r http.ResponseWriter, req *http.Request) {
		r.Write([]byte("PONG"))
	})

	// Alert Testing
	go func() {
		// time.Sleep(20 * time.Second)
		// os.Exit(0)
	}()

	log.Fatal(http.ListenAndServe(":8085", pMux))
}
