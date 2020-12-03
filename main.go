package main

import (
	"database/sql"
	"fmt"

	_ "github.com/nakagami/firebirdsql"
)

// var (
// 	DB_HOST="192.168.0.2"
// 	DB_PORT="3050"
// 	DB_DATABASE="C:\\Program Files (x86)\\CompuFour\\Clipp\\Base\\CLIPP.FDB"
// 	DB_USERNAME="SYSDBA"
// 	DB_PASSWORD="masterkey"
// 	DB_CHARSET="WIN1252"
// )

func main() {
	var n int
	conn, _ := sql.Open("firebirdsql", "SYSDBA:masterkey@192.168.0.2C:\\Program Files (x86)\\CompuFour\\Clipp\\Base\\CLIPP.FDB")
	defer conn.Close()

	conn.QueryRow("SELECT Count(*) FROM rdb$relations").Scan(&n)

	fmt.Println("Relations count = ", n)
}