package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/stevenwilkin/treasury/bybit"
)

type fundingMessage struct {
	Current   float64 `json:"current"`
	Predicted float64 `json:"predicted"`
}

var (
	port int
)

func init() {
	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.Parse()
}

func fundingHandler(w http.ResponseWriter, r *http.Request) {
	exchange := &bybit.Bybit{}
	current, predicted := exchange.GetFundingRate()

	fm := fundingMessage{
		Current:   current,
		Predicted: predicted}

	b, err := json.Marshal(fm)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	log.Printf("Starting on http://0.0.0.0:%d\n", port)
	http.Handle("/", http.FileServer(http.Dir("www")))
	http.HandleFunc("/funding", fundingHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
