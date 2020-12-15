package konfig

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/nakagami/firebirdsql"
)

var dbara string
var dbcwb string
var dblon string
var dbnat string
var dbrec string

var db *sql.DB
var err error

func init(){
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error reading .env file")
	}

	dbara = os.Getenv("ARA")
	dbcwb = os.Getenv("CWB")
	dblon = os.Getenv("LON")
	dbnat = os.Getenv("NAT")
	dbrec = os.Getenv("REC")

}

// KonnektAra connect to Aracaju database
func KonnektAra() (db *sql.DB, err error) {
	db, err = sql.Open("firebirdsql", dbara)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// KonnektCwb - connect to Curitiba database
func KonnektCwb() (db *sql.DB, err error) {
	db, err = sql.Open("firebirdsql", dbcwb)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// KonnektLon Connect to Londrina database
func KonnektLon() (db *sql.DB, err error) {
	db, err = sql.Open("firebirdsql", dblon)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// KonnektNat Connect to Natal database
func KonnektNat() (db *sql.DB, err error) {
	db, err = sql.Open("firebirdsql", dbnat)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// KonnektRec Connect to Recife database
func KonnektRec() (db *sql.DB, err error) {
	db, err = sql.Open("firebirdsql", dbrec)

	if err != nil {
		return nil, err
	}

	return db, nil
}
