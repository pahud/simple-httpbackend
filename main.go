package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path[1:]
	if len(p) == 0 {
		for _, pair := range os.Environ() {
			fmt.Fprintf(w, "%s\n", pair)
		}
	} else {
		// fmt.Fprintf(w, "[DEBUG] path: %s\n", p)
		fmt.Fprintf(w, "%s\n", os.Getenv(p))
	}

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
