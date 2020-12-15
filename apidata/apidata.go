package apidata

import (
	"log"

	"github.com/wborbajr/osservice/konfig"
)

//
// https://goplay.space/#mWw59cjYPh7
// https://play.golang.org/p/mWw59cjYPh7
// https://stackoverflow.com/questions/27795036/create-chan-for-func-with-two-return-args#27795117
//

type ModelOSData struct {
	IdOs      int `json:"idos"`
	IdCliente int `json:"idcliente"`
	IdStatus 	int `json:"idstatus"`
}

var osData = []*ModelOSData{
	{
		IdOs:     1,
		IdCliente:1001,
		IdStatus: 0,
	},
	{
		IdOs:     2,
		IdCliente:2001,
		IdStatus: 1,
	},
}

// GetOS search for OS status
func GetOS(doc string, os string) (ModelOSData, error) {

	log.Println("GetOS")

	_, err := konfig.KonnektAra()

	if err != nil {
		return nil, nil
	}

	log.Println("Connected to Curitiba")

	return []*ModelOSData{
		"OS": osData,
	}, nil

}

