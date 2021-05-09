package main

import (
	"XCloud/cmd/objects"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	s := os.Getenv("STORAGE_ROOT") //NUMBER_OF_PROCESSORS
	//s := os.Getenv("NUMBER_OF_PROCESSORS")
	fmt.Println(s)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
