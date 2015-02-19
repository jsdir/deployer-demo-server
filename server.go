package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from `%s`! (%s)\n", os.Getenv("BUILD"), os.Getenv("RELEASE_NAME"))
	fmt.Fprintf(w, "I'm running in the `%s` environment.\n", os.Getenv("ENVIRONMENT"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
