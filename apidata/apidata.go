package apidata

import (
	"log"

	"github.com/wborbajr/osservice/konfig"
)

//
// https://play.golang.org/p/mWw59cjYPh7
// https://stackoverflow.com/questions/27795036/create-chan-for-func-with-two-return-args#27795117
//

// GetOS search for OS status
func GetOS(doc string, os string) error {

	log.Println("GetOS")

	_, err := konfig.KonnektAra()

	if err != nil {
		log.Fatal("Error connecting to Curitiba database")
	}

	log.Println("Connected to Curitiba")

	return nil
}

