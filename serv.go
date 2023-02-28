package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	PATH = "/time"
	PORT = ":8795"
)

func main() {
	http.HandleFunc(PATH, TimeHandler)
	fmt.Printf("Starting server at port %s", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
This is for 1 com
This is for 2 cum
this is for commit 1 before vlad
this is commit 2