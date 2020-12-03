package config

import (
	"database/sql"

	_ "github.com/nakagami/firebirdsql"
)

func Konnekt() (db *sql.DB, err error) {
	db, err = sql.Open("firebirdsql", "SYSDBA:masterkey@192.168.0.2/C:/Program Files (x86)/CompuFour/Clipp/Base/CLIPP.FDB?charset=WIN1252")

	if err != nil {
		return nil, err
	}

	return db, nil
}