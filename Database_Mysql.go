package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

type chat struct {
	namaUser string
	pesan string
	id int
}

func connectMysql() *sql.DB {
	dsn := "root:4l3*z@42@/hellosailsDB"
	var db, err = sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	return db
}

func mysqlQuery() {
	var db = connectMysql()

	defer db.Close()

	var rows, err = db.Query("select user,message,id from chat")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	defer rows.Close()

	var result []chat

	for rows.Next() {
		var each = chat{}
		var err = rows.Scan(&each.namaUser, &each.pesan,&each.id)

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
		fmt.Println(each.namaUser)
	}
}

func main() {
	mysqlQuery()
}
