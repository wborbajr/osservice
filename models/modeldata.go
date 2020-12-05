package models

import (
	"database/sql"
	"fmt"

	_struct "github.com/wborbajr/osservice/struct"
)

type ModelGetData struct {
	DB *sql.DB
}

type TbOS struct {
	IdOs       int
	IdCliente  int
	IdStatus int
}

func (model ModelGetData) GetOS(doc string, os string) (getStruct []_struct.StructData, err error) {

	// var contentOS TbOS
	var data _struct.StructData

	row := model.DB.QueryRow("SELECT ID_OS, ID_CLIENTE, ID_STATUS FROM TB_OS WHERE ID_OS = ? AND ID_CLIENTE = ?;", os, doc)
	errdb := row.Scan(&data.IdOs, &data.IdCliente, &data.IdStatus)

	if errdb != nil {
		if errdb == sql.ErrNoRows {
			 fmt.Println("Zero rows found")
			//  panic(errdb.Error())
		} else {
			 panic(errdb)
		}
  }

	// if err_db == sql.ErrNoRows  {
	// 	return nil, err_db
	// }

	var _isiStruct []_struct.StructData
	// var data _struct.StructData
	// for row.Next() {
	// 	err2 := row.Scan(
	// 		&data.IdOs,
	// 		&data.IdCliente,
	// 		&data.IdStatus)
	// 	if err2 != nil {
	// 		return nil, err2
	// 	} else {
			_data := _struct.StructData{
				IdOs:      data.IdOs,
				IdCliente: data.IdCliente,
				IdStatus:  data.IdStatus,
			}
			_isiStruct = append(_isiStruct, _data)

	// 	}
	// }

	return _isiStruct, nil

}