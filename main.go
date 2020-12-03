package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/wborbajr/osservice/apis"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/osservice/getos",apis.GetOS).Methods("GET")

	err := http.ListenAndServe(":3001",router)

	if err != nil {
		log.Println(err)
	}

}