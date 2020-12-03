package main

import (
	"fmt"

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
		fmt.Println("Error retrieving data.")
		db.Close()
	}

	response.Data = IsiData

	defer db.Close()

	fmt.Println("Relations count = %s", response)
}