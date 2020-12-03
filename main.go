package main

import (
	"fmt"
	"log"

	"github.com/wborbajr/osservice/config"
	"github.com/wborbajr/osservice/models"
	_struct "github.com/wborbajr/osservice/struct"
)

func main() {

	var response _struct.ResponseData

	db, err := config.Konnekt()
	if err != nil {
		log.Fatal("Error connecting database! Exiting. ", err)
		db.Close()
	}

	_models := models.ModelGetData{DB:db}
	IsiData, err := _models.GetOS()

	if err != nil {
		log.Fatal("Error retrieving data. ", err)
		db.Close()
	}

	response.Data = IsiData

	defer db.Close()

	fmt.Println("Relations count = %s", response)
}