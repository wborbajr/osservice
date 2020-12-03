package models

import (
	"database/sql"

	_struct "github.com/wborbajr/osservice/struct"
)


type ModelGetData struct {
	DB *sql.DB
}

func (model ModelGetData) GetOS() (getStruct []_struct.StructData, err error) {

	row, err := model.DB.Query("SELECT * FORM V_OS")

	if err != nil {
		return nil, err

	} else {
		var _isiStruct []_struct.StructData
		var data _struct.StructData
		for row.Next() {
			err2 := row.Scan(
				&data.Id,
				&data.FirstName,
				&data.LastName)
			if err2 != nil {
				return nil, err2
			} else {
				_data := _struct.StructData{
					Id:        data.Id,
					FirstName: data.FirstName,
					LastName:  data.LastName,
				}
				_isiStruct = append(_isiStruct, _data)

			}
		}

		return _isiStruct, nil
	}
}