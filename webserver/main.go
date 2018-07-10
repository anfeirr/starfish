package main

import (
	"io"
	"net/http"
)

func starFish(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>我们的征程是星辰大海</h1>")
}

func main() {
	http.HandleFunc("/", starFish)
	http.ListenAndServe(":8000", nil)
}
