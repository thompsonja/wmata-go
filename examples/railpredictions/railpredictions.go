package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/thompsonja/wmata-go/pkg/railpredictions"
)

// This is an example of how to use one of the APIs, in this case the railpredictions API.
func main() {
	apiKey := flag.String("api_key", "", "WMATA API key")
	flag.Parse()

	client := railpredictions.New(*apiKey)

	predictions, err := client.GetRailPredictions(context.Background(), "all")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(predictions)
}
