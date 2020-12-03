package main

import (
	"database/sql"
	"fmt"

	_ "github.com/nakagami/firebirdsql"
)

func main() {
    var n int
    conn, _ := sql.Open("firebirdsql", "SYSDBA:masterkey@192.168.0.2/C:/Program Files (x86)/CompuFour/Clipp/Base/CLIPP.FDB")
    defer conn.Close()
    conn.QueryRow("SELECT Count(*) FROM rdb$relations").Scan(&n)
    fmt.Println("Relations count=", n)

}