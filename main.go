package main

import (
	"log"

	"github.com/wborbajr/osservice/server"
)



func main() {
	log.Printf("Starting the web server...")
	server.SetupApp()

}