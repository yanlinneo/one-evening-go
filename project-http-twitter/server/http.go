package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Server struct {
	TweetsRepository TweetsRepository
}

type Tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type response struct {
	ID int `json:"ID"`
}

type tweetsList struct {
	Tweets []Tweet `json:"tweets"`
}

func (s Server) TweetsEndpoint(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	defer func() {
		duration := time.Since(start)
		fmt.Printf("%s %s %s\n", r.Method, r.URL, duration)
	}()

	/*
		if r.Method == http.MethodPost {
			s.addTweet(w, r)
		} else {
			s.listTweets(w, r)
		}
	*/
}

func (s Server) AddTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tw := Tweet{}

	if err := json.Unmarshal(body, &tw); err != nil {
		log.Printf("Failed to unmarshal payload: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if tw.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := s.TweetsRepository.AddTweet(tw)
	if err != nil {
		log.Printf("Failed to add tweet: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := response{
		ID: id,
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Failed to marshal: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respJSON)
}

func (s Server) ListTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := s.TweetsRepository.Tweets()
	if err != nil {
		log.Printf("Failed to get tweets: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := tweetsList{
		Tweets: tweets,
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Failed to get marshal tweets: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respJSON)
}
