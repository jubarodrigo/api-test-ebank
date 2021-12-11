package app

import (
	"ebanx/handle"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartServer() {
	router := mux.NewRouter()

	router.HandleFunc("/reset", handle.Reset).Methods("GET")
	router.HandleFunc("/balance", handle.GetBalance).Methods("GET")
	router.HandleFunc("/event", handle.Account).Methods("POST")

	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
