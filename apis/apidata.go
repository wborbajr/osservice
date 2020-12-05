package apis

import (
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

	params := mux.Vars(r)

	if params["doc"] != "" && params["os"] != "" {
		fmt.Println(params["doc"], params["os"])
	}

	doc := params["doc"]
	os := params["os"]

	fmt.Println("Konnekt_ara()")
	db, err := config.Konnekt_ara()
	var Response _struct.ResponseData

	if err != nil {
		Response.Status = http.StatusInternalServerError
		Response.Message = err.Error()
		Response.Data = nil
		restponWithJson(w, http.StatusInternalServerError, Response)
	}  else {
		_models := models.ModelGetData{DB:db}
		IsiData, err2 := _models.GetOS(doc, os)
		if err2 != nil {
			Response.Status = http.StatusInternalServerError
			Response.Message = err2.Error()
			Response.Data = nil
			restponWithJson(w, http.StatusInternalServerError, Response)

		} else {
			Response.Status = http.StatusOK
			Response.Message = "Sukses"
			Response.Data = IsiData
			restponWithJson(w, http.StatusOK, Response)

		}
	}

}

func restponWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}