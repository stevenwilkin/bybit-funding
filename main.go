package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type fundingMessage struct {
	Current   float64 `json:"current"`
	Predicted float64 `json:"predicted"`
}

type fundingResponse struct {
	Result []struct {
		FundingRate          string `json:"funding_rate"`
		PredictedFundingRate string `json:"predicted_funding_rate"`
	} `json:"result"`
}

var (
	port int
)

func init() {
	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.Parse()
}

func fundingHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://api.bybit.com/v2/public/tickers?symbol=BTCUSD"
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err.Error())
	}

	var response fundingResponse
	json.Unmarshal(body, &response)

	current, _ := strconv.ParseFloat(response.Result[0].FundingRate, 64)
	predicted, _ := strconv.ParseFloat(response.Result[0].PredictedFundingRate, 64)

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
