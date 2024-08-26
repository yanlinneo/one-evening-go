package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"twitter/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func main() {
	s := server.Server{
		TweetsRepository: &server.TweetsMemoryRepository{},
	}

	// http.HandleFunc("/tweets", s.TweetsEndpoint)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	go spamTweets()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	rateLimit := httprate.LimitByIP(10, 1*time.Minute)
	r.With(rateLimit).Post("/tweets", s.AddTweet)
	r.Get("/tweets", s.ListTweets)

	//r.Post("/tweets", s.AddTweet)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func spamTweets() {
	for {
		// Prepare payload
		addTweetPayload := server.Tweet{
			Message:  "ass",
			Location: "ass",
		}

		// Marshal payload
		marshaledPayload, err := json.Marshal(addTweetPayload)
		if err != nil {
			fmt.Println(err)
		}

		// Send HTTP POST request
		url := "http://localhost:8080/tweets"

		resp, respErr := http.Post(url, "application/json", bytes.NewBuffer(marshaledPayload))
		if respErr != nil {
			fmt.Println(respErr)
		}

		// Close response body
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		// (Optionally read and print the response)
		fmt.Println(body)
	}
}
