package models

import (
	"database/sql"

	_struct "github.com/wborbajr/osservice/struct"
)


type ModelGetData struct {
	DB *sql.DB
}

func (model ModelGetData) GetOS() (getStruct []_struct.StructData, err error) {

	row, err := model.DB.Query("SELECT ID_OS, ID_CLIENTE, ID_STATUS FROM TB_OS WHERE ID_OS = 117949;")

	if err != nil {
		return nil, err

	} else {
		var _isiStruct []_struct.StructData
		var data _struct.StructData
		for row.Next() {
			err2 := row.Scan(
				&data.IdOs,
				&data.IdCliente,
				&data.IdStatus)
			if err2 != nil {
				return nil, err2
			} else {
				_data := _struct.StructData{
					IdOs:      data.IdOs,
					IdCliente: data.IdCliente,
					IdStatus:  data.IdStatus,
				}
				_isiStruct = append(_isiStruct, _data)

			}
		}

		return _isiStruct, nil
	}
}