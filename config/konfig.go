package config

import (
	"database/sql"

	_ "github.com/nakagami/firebirdsql"
)

func KonnektCWB() (db *sql.DB, err error) {

	// dbcwb := os.Getenv("CWB")

	sql.Open("firebirdsql", "SYSDBA:masterkey@192.168.0.2/C:/Program Files (x86)/CompuFour/Clipp/Base/CLIPP.FDB?charset=WIN1252")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func KonnektLON() (db *sql.DB, err error) {
	db, err = sql.Open("firebirdsql", "SYSDBA:masterkey@192.168.0.2/C:/Program Files (x86)/CompuFour/Clipp/Base/CLIPP.FDB?charset=WIN1252")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func KonnektNAT() (db *sql.DB, err error) {
	db, err = sql.Open("firebirdsql", "SYSDBA:masterkey@192.168.0.2/C:/Program Files (x86)/CompuFour/Clipp/Base/CLIPP.FDB?charset=WIN1252")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func KonnektREC() (db *sql.DB, err error) {
	db, err = sql.Open("firebirdsql", "SYSDBA:masterkey@192.168.0.2/C:/Program Files (x86)/CompuFour/Clipp/Base/CLIPP.FDB?charset=WIN1252")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func KonnektARA() (db *sql.DB, err error) {
	db, err = sql.Open("firebirdsql", "SYSDBA:masterkey@192.168.0.2/C:/Program Files (x86)/CompuFour/Clipp/Base/CLIPP.FDB?charset=WIN1252")

	if err != nil {
		return nil, err
	}

	return db, nil
}


