package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dexterorion/smartmei/handler"
)

func main() {
	log.Print("service started")
	log.Fatal("search service closed",
		http.ListenAndServe(
			fmt.Sprintf(":%d", 5656),
			handler.New(),
		))
}
