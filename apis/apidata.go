package apis

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wborbajr/osservice/config"
	"github.com/wborbajr/osservice/models"
	_struct "github.com/wborbajr/osservice/struct"
)

func GetOS(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var errcount int

	// IsiData := make(chan []_struct.StructData)
	var IsiData []_struct.StructData
	var err2ara, err2cwb, err2lon, err2nat, err2rec error

	params := mux.Vars(r)

	if params["doc"] != "" && params["os"] != "" {
		fmt.Println(params["doc"], params["os"])
	}

	doc := params["doc"]
	os := params["os"]

	dbara, errara := config.Konnekt_ara()
	dbcwb, errcwb := config.Konnekt_cwb()
	dblon, errlon := config.Konnekt_lon()
	dbnat, errnat := config.Konnekt_nat()
	dbrec, errrec := config.Konnekt_rec()

	if errara == nil {
		_modelsara := models.ModelGetData{DB:dbara}
		// go func() {
			IsiData, err2ara = _modelsara.GetOS(doc, os)
			if err2ara == sql.ErrNoRows {
				errcount++
			}
		// }()
	}
	if errcwb == nil {
		_modelscwb := models.ModelGetData{DB:dbcwb}
		// go func() {
			IsiData, err2cwb = _modelscwb.GetOS(doc, os)
			if err2cwb == sql.ErrNoRows {
				errcount++
			}
		// }()
	}
	if errlon == nil {
		_modelslon := models.ModelGetData{DB:dblon}
		// go func() {
			IsiData, err2lon = _modelslon.GetOS(doc, os)
			if err2lon == sql.ErrNoRows {
				errcount++
			}
		// }()
	}
	if errnat == nil {
		_modelsnat := models.ModelGetData{DB:dbnat}
		// go func() {
			IsiData, err2nat = _modelsnat.GetOS(doc, os)
			if err2nat == sql.ErrNoRows {
				errcount++
			}
		// }()
	}
	if errrec == nil {
		_modelsrec := models.ModelGetData{DB:dbrec}
		// go func() {
			IsiData, err2rec = _modelsrec.GetOS(doc, os)
			if err2rec == sql.ErrNoRows {
				errcount++
			}
		// }()
	}

	var Response _struct.ResponseData

	// record not found
	if errcount == 5 {
		Response.Status = http.StatusInternalServerError
		Response.Message = "OS Not Found"
		Response.Data = nil
		restponWithJson(w, http.StatusInternalServerError, Response)
	} else {
		Response.Status = http.StatusOK
		Response.Message = "Sukses"
		Response.Data = IsiData
		restponWithJson(w, http.StatusOK, Response)
	}







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