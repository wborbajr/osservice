package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/wborbajr/osservice/config"
	"github.com/wborbajr/osservice/models"
	_struct "github.com/wborbajr/osservice/struct"
)

//
// https://play.golang.org/p/mWw59cjYPh7
// https://stackoverflow.com/questions/27795036/create-chan-for-func-with-two-return-args#27795117
//

type Ret struct {
	IsiData []_struct.StructData
	err error
}

func GetOS(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var waitGroup sync.WaitGroup

	chanret := make(chan Ret)

	params := mux.Vars(r)

	if params["doc"] != "" && params["os"] != "" {
		fmt.Println(params["doc"], params["os"])
	}

	doc := params["doc"]
	os := params["os"]

	dbara, errara := config.Konnekt_ara()
	// dbcwb, errcwb := config.Konnekt_cwb()
	// dblon, errlon := config.Konnekt_lon()
	// dbnat, errnat := config.Konnekt_nat()
	// dbrec, errrec := config.Konnekt_rec()

	waitGroup.Add(1)

	if errara == nil {
		_modelsara := models.ModelGetData{DB:dbara}
		go func() {
			defer waitGroup.Done()
			tmpdata, temperr := _modelsara.GetOS(doc, os)
			chanret <- Ret{tmpdata, temperr}
			close(chanret)
		}()
	}

	ara := <-chanret

	if ara.err == nil {
		var Response _struct.ResponseData
		Response.Status = http.StatusOK
		Response.Message = "Sukses"
		Response.Data = ara.IsiData
		restponWithJson(w, http.StatusOK, Response)
	}


	// var Response _struct.ResponseData

	// // record not found
	// if errcount == 5 {
	// 	Response.Status = http.StatusInternalServerError
	// 	Response.Message = "OS Not Found"
	// 	Response.Data = nil
	// 	restponWithJson(w, http.StatusInternalServerError, Response)
	// } else {
	// 	Response.Status = http.StatusOK
	// 	Response.Message = "Sukses"
	// 	Response.Data = IsiData
	// 	restponWithJson(w, http.StatusOK, Response)
	// }







	// var Response _struct.ResponseData

	// if err != nil {
	// 	Response.Status = http.StatusInternalServerError
	// 	Response.Message = err.Error()
	// 	Response.Data = nil
	// 	restponWithJson(w, http.StatusInternalServerError, Response)
	// }  else {
	// 	_models := models.ModelGetData{DB:db}
	// 	IsiData, err2 := _models.GetOS(doc, os)
	// 	if err2 != nil {
	// 		Response.Status = http.StatusInternalServerError
	// 		Response.Message = err2.Error()
	// 		Response.Data = nil
	// 		restponWithJson(w, http.StatusInternalServerError, Response)

	// 	} else {
	// 		Response.Status = http.StatusOK
	// 		Response.Message = "Sukses"
	// 		Response.Data = IsiData
	// 		restponWithJson(w, http.StatusOK, Response)

	// 	}
	// }

}

func restponWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
