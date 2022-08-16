package main

import (
	"io"
	"log"
	"net/http"
)

func hand(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hellooooooo\n")
}

func main() {
	http.HandleFunc("/", hand)
	err := http.ListenAndServe(":8885", nil)
	if err != nil {
		log.Fatal(err)
	}
}
