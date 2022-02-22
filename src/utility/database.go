package utility

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func getDatabase() *sql.DB {
	db, err := sql.Open("mysql", "codebrick:111222qq@tcp(codebrick-dev.c9s5okqt95x3.ap-southeast-1.rds.amazonaws.com)/tokotalk")
	if err != nil {
		panic(err)
	}

	return db
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func database_1() {
	db := getDatabase()
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

func database_2() {
	db := getDatabase()
	var name string
	err := db.QueryRow("SELECT name FROM tokotalk.shipment_methods where id = ?", 2).Scan(&name)
	if err != nil {
		panic(err)
	}
	fmt.Println(name)
}

func database_3() {
	db := getDatabase()
	defer db.Close()

	result, err := db.Exec("insert into tokotalk.news (vendor_id, title, text) values (?, ?, ?)", 1, "testTitle", "testText")
	if err != nil {
		panic(err)
	}

	n, err := result.RowsAffected()
	checkError(err)
	if n > 0 {
		fmt.Printf("%v row inserted.\n", n)
	}
}

func database_4() {
	db := getDatabase()
	defer db.Close()

	stmt, err := db.Prepare("update news set title = ? where id = ?")
	checkError(err)
	defer stmt.Close()

	result, err := stmt.Exec("title-title", 37)
	checkError(err)

	n, err := result.RowsAffected()
	if err == nil {
		fmt.Printf("%v row affected.\n", n)
	}
}

func database_5() {
	db := getDatabase()
	defer db.Close()

	tx, err := db.Begin()
	checkError(err)
	defer func() {
		fmt.Println("run defer")
		if r := recover(); r != nil {
			fmt.Println("this is panic happened")
			fmt.Printf("%v , %T \n", r, r)
		}
		tx.Rollback()
	}()

	// a := "123123"
	// _ = a[:99]
	// e1 := errors.New("testError")
	// if e1 != nil {
	// 	panic("hahaha")
	// }

	_, err = tx.Exec("insert into tokotalk.news (vendor_id, title, text) values (?, ?, ?)", 1, "testTitle123", "testText123")
	checkError(err)
	fmt.Println("insert 1")

	_, err = tx.Exec("insert into tokotalk.news (vendor_id, title, text) values (?, ?, ?)", 1, "testTitle2222", "testText2222")
	checkError(err)
	fmt.Println("insert 2")

	err = tx.Commit()
	checkError(err)
}

func StudyDatabase() {
	database_5()
}
