package utility

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func database_1() {
	db, err := sql.Open("mysql", "codebrick:111222qq@tcp(codebrick-dev.c9s5okqt95x3.ap-southeast-1.rds.amazonaws.com)/tokotalk")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM shipment_methods limit 10")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var (
		id   int
		name string
	)
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id:%v, name:%v \n", id, name)
	}
}

func StudyDatabase() {
	database_1()
}
