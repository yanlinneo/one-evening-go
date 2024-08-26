package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	calls []string
	stats = map[string]int{}
)

func main() {
	// Start an HTTP server
	http.HandleFunc("/hello", yourFunction)
	log.Fatal(http.ListenAndServe(":8080", nil))
	// The server is started on the port you pass to http.ListenAndServe. It will run forever, until you close the application (by pressing Ctrl+C). In the example above, it will be available at http://localhost:8080.

	// Your solution goes here. Good luck!
}

func yourFunction(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // request -> url -> query (?name=) -> name

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	calls = append(calls, name)
	stats[name]++

	fmt.Printf("calls: %#v\n", calls)
	fmt.Printf("stats: %#v\n\n", stats)

	fmt.Fprint(w, "Hello, ", name)
}
