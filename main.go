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
		fmt.Println("Error connecting database! Exiting.")
		db.Close()
	}

	_models := models.ModelGetData{DB:db}
	IsiData, err2 := _models.GetOS()

	if err2 != nil {
		log.Fatal("Error retrieving data. ", err2)
		db.Close()
	}

	response.Data = IsiData

	defer db.Close()

	fmt.Println("Relations count = %s", response)
}