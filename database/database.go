package database

import (
	"database/sql"
	"os"

	_ "github.com/nakagami/firebirdsql"
)

var dbara string
var dbcwb string
var dblon string
var dbnat string
var dbrec string

var DB *sql.DB
var err error

func init(){

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error reading .env file: ", err)
	// }

	dbara = os.Getenv("ARA")
	dbcwb = os.Getenv("CWB")
	dblon = os.Getenv("LON")
	dbnat = os.Getenv("NAT")
	dbrec = os.Getenv("REC")

}

// KonnektAra connect to Aracaju database
func KonnektAra() (error) {
	DB, err = sql.Open("firebirdsql", dbara)

	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	return nil
}

// KonnektCwb - connect to Curitiba database
func KonnektCwb() (error) {
	DB, err = sql.Open("firebirdsql", dbcwb)

	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	return nil
}

// KonnektLon Connect to Londrina database
func KonnektLon() (error) {
	DB, err = sql.Open("firebirdsql", dblon)

	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	return nil

}

// KonnektNat Connect to Natal database
func KonnektNat() (error) {
	DB, err = sql.Open("firebirdsql", dbnat)

	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	return nil

}

// KonnektRec Connect to Recife database
func KonnektRec() (error) {
	DB, err = sql.Open("firebirdsql", dbrec)

	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	return nil

}
