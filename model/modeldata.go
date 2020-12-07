package model

import (
	"database/sql"
)

type ModelGetData struct {
	DB *sql.DB
}

// type TbOS struct {
// 	IdOs       int
// 	IdCliente  int
// 	IdStatus int
// }

// func (model ModelGetData) GetOS(doc string, os string) (getStruct []_struct.StructData, err error) {

// 	var data _struct.StructData

// 	dbara, errara := config.Konnekt_ara()
// 	dbcwb, errcwb := config.Konnekt_cwb()
// 	dblon, errlon := config.Konnekt_lon()
// 	dbnat, errnat := config.Konnekt_nat()
// 	dbrec, errrec := config.Konnekt_rec()


// 	if errara != nil {
// 		log.Println("Aracaju - OFFLine")
// 	}
// 	if errcwb != nil {
// 		log.Println("Curitiba - OFFLine")
// 	}
// 	if errlon != nil {
// 		log.Println("Londrina - OFFLine")
// 	}
// 	if errnat != nil {
// 		log.Println("Natal - OFFLine")
// 	}
// 	if errrec != nil {
// 		log.Println("Recife - OFFLine")
// 	}


// 	row := model.DB.QueryRow("SELECT ID_OS, ID_CLIENTE, ID_STATUS FROM TB_OS WHERE ID_OS = ? AND ID_CLIENTE = ?;", os, doc)
// 	errdb := row.Scan(&data.IdOs, &data.IdCliente, &data.IdStatus)

// 	if errdb != nil {
// 		if errdb == sql.ErrNoRows {
// 			 fmt.Println("Zero rows found")
// 			//  panic(errdb.Error())
// 		} else {
// 			 panic(errdb)
// 		}
//   	}
// 	var _isiStruct []_struct.StructData

// 	_data := _struct.StructData{
// 		IdOs:      data.IdOs,
// 		IdCliente: data.IdCliente,
// 		IdStatus:  data.IdStatus,
// 	}
// 	_isiStruct = append(_isiStruct, _data)

// 	return _isiStruct, nil

// }