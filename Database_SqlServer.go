package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/denisenkom/go-mssqldb"
)

type tBaseStatus struct {
	StatusId   string
	StatusName string
	Descript   string
	Categ      string
	Aktif      bool
}

func connect() *sql.DB {
	dsn := "server=172.16.64.109;user id=845879;password=5879;database=JobSchDB"
	var db, err = sql.Open("mssql", dsn)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	return db
}

func sqlQuery() {
	var db = connect()

	defer db.Close()

	var rows, err = db.Query("select StsId,StsName,Descript,Category,Active from tBase_Status")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	defer rows.Close()

	var result []tBaseStatus

	for rows.Next() {
		var each = tBaseStatus{}
		var err = rows.Scan(&each.StatusId, &each.StatusName,
			&each.Descript, &each.Categ, &each.Aktif)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	for _, each := range result {
		fmt.Println(each.StatusName)
	}
}

func main() {
	//var db = connect()
	//defer db.Close()

	sqlQuery()
}
