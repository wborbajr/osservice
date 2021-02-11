package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wborbajr/osservice/server"
)

func main() {
	if os.Getenv("FEATURE_TOGGLE") == "TRUE" {
		fmt.Println(os.Getenv("FEATURE_TOGGLE"))
		fmt.Println("Exciting New Feature")
	} else {
		fmt.Println(os.Getenv("FEATURE_TOGGLE"))
		fmt.Println("existing boring feature")
	}

	log.Printf("Starting up OSService")

	log.Printf("Starting the web server...")
	server.SetupApp()
}
