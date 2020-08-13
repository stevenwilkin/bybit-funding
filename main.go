package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/stevenwilkin/treasury/bybit"
)

type funding struct {
	Current   float64
	Predicted float64
}

var (
	port     int
	exchange *bybit.Bybit
)

func init() {
	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.Parse()
}

func fundingHandler(w http.ResponseWriter, r *http.Request) {
	current, predicted := exchange.GetFundingRate()

	f := funding{
		Current:   current * 100,
		Predicted: predicted * 100}

	tmpl.Execute(w, f)
}

func main() {
	log.Printf("Starting on http://0.0.0.0:%d\n", port)

	http.HandleFunc("/", fundingHandler)

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("www")))
	http.Handle("/static/", fs)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
