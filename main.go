package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Music server on port :5556")
	log.Fatal(http.ListenAndServe(":5556", Server()))
}
