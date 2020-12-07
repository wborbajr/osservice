package apis

import (
	"github.com/gofiber/fiber"
	_struct "github.com/wborbajr/osservice/struct"
)

//
// https://play.golang.org/p/mWw59cjYPh7
// https://stackoverflow.com/questions/27795036/create-chan-for-func-with-two-return-args#27795117
//

type tempStruct  struct {
	IsiData []_struct.StructData
	err error
}

func GetOS(c *fiber.Ctx){

}

// func GetOS(c *fiber.Ctx) {

// 	// w.Header().Set("Content-Type", "application/json")

// 	var waitGroup sync.WaitGroup

// 	chanret := make(chan tempStruct)

// 	params := mux.Vars(r)

// 	if params["doc"] != "" && params["os"] != "" {
// 		fmt.Println(params["doc"], params["os"])
// 	}

// 	doc := c.Params["doc"]
// 	os := c.Params["os"]

// 	waitGroup.Add(1)

// 	_modelsara := models.ModelGetData{DB:dbara}
// 	go func() {
// 		defer waitGroup.Done()
// 		tmpdata, temperr := _modelsara.GetOS(doc, os)
// 		chanret <- tempStruct{tmpdata, temperr}
// 		close(chanret)
// 	}()

// 	ara := <-chanret

// 	if ara.err == nil {
// 		var Response _struct.ResponseData
// 		Response.Status = http.StatusOK
// 		Response.Message = "Sukses"
// 		Response.Data = ara.IsiData
// 		restponWithJson(w, http.StatusOK, Response)
// 	}


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

// }

// func restponWithJson(w http.ResponseWriter, code int, payload interface{}) {
// 	response, _ := json.Marshal(payload)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }
