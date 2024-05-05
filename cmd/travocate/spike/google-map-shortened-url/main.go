package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func main() {
	shortURL := flag.String("url", "", "Shortened Google Maps URL")
	flag.Parse()

	if *shortURL == "" {
		log.Fatal("Usage: go run main.go -url=<shortened Google Maps URL>")
	}

	res, err := http.Get(*shortURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("failed to follow short URL: status code %d", res.StatusCode)
	}

	re := regexp.MustCompile(`/@([-\d.]+),([-\d.]+)`)
	matches := re.FindStringSubmatch(res.Request.URL.Path)

	if len(matches) == 3 {
		latitude := matches[1]
		longitude := matches[2]
		fmt.Printf("Latitude: %s\n", latitude)
		fmt.Printf("Longitude: %s\n", longitude)
	} else {
		fmt.Println("Coordinates not found")
	}
}
