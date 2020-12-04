package config

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/nakagami/firebirdsql"
)

func Konnekt_ara() (db *sql.DB, err error) {

	var dbara string

	err_env := godotenv.Load(".env")

	if err_env != nil {
		panic("Error loading .env file: ")
	}
	dbara = os.Getenv("ARA")

	db, err = sql.Open("firebirdsql", dbara)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Konnekt_cwb() (db *sql.DB, err error) {

	var dbcwb string

	err_env := godotenv.Load(".env")

	if err_env != nil {
		panic("Error loading .env file: ")
	}
	dbcwb = os.Getenv("CWB")

	db, err = sql.Open("firebirdsql", dbcwb)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Konnekt_lon() (db *sql.DB, err error) {

	var dblon string

	err_env := godotenv.Load(".env")

	if err_env != nil {
		panic("Error loading .env file: ")
	}
	dblon = os.Getenv("LON")

	db, err = sql.Open("firebirdsql", dblon)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Konnekt_nat() (db *sql.DB, err error) {

	var dbnat string

	err_env := godotenv.Load(".env")

	if err_env != nil {
		panic("Error loading .env file: ")
	}
	dbnat = os.Getenv("NAT")

	db, err = sql.Open("firebirdsql", dbnat)

	if err != nil {
		return nil, err
	}

	return db, nil
}


func Konnekt_rec() (db *sql.DB, err error) {

	var dbrec string

	err_env := godotenv.Load(".env")

	if err_env != nil {
		panic("Error loading .env file: ")
	}
	dbrec = os.Getenv("REC")

	db, err = sql.Open("firebirdsql", dbrec)

	if err != nil {
		return nil, err
	}

	return db, nil
}
