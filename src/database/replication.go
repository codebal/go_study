package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func getDatabase(connStr string) *sql.DB {
	db, err := sql.Open("mysql", connStr)
	errorHandle(err)

	err = db.Ping()
	errorHandle(err)

	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1000)
	fmt.Printf("db connected : %+v\n", connStr)
	fmt.Printf("status : %+v\n", db.Stats())

	return db
}

var (
	master, slave                                                                     *sql.DB
	allCount, allCountOk, allCountFail                                                int64
	allElapsedTime, allAvgElapsedTime                                                 int64
	allElapsedTimeOk, allAvgElapsedTimeOk, minElapsedTimeOk, maxElapsedTimeOk         int64
	allElapsedTimeFail, allAvgElapsedTimeFail, minElapsedTimeFail, maxElapsedTimeFail int64
	insertDelayCount                                                                  int32
)

func initDatabases() {
	master = getDatabase(`codebrick:111222qq@tcp(codebrick-test-original.c9s5okqt95x3.ap-southeast-1.rds.amazonaws.com)/tmp?charset=utf8mb4,utf8&sql_mode=TRADITIONAL&multiStatements=true&parseTime=true`)
	slave = getDatabase(`codebrick:111222qq@tcp(codebrick-test-slave-1.c9s5okqt95x3.ap-southeast-1.rds.amazonaws.com)/tmp?charset=utf8mb4,utf8&sql_mode=TRADITIONAL&multiStatements=true&parseTime=true`)
}

func errorHandle(err error) {
	if err != nil {
		panic(err)
	}
}

func insertAndSelect(newName string) {
	startTime := time.Now()
	atomic.AddInt32(&insertDelayCount, 1)
	lastId := insertData(newName)
	atomic.AddInt32(&insertDelayCount, -1)
	//fmt.Printf("%v - new id : %v\n", newName, lastId)

	result := getData(lastId)
	elapsedTime := time.Since(startTime)
	elapsedTimeNano := elapsedTime.Nanoseconds()
	//fmt.Printf("실행시간: %v, %v\n", elapsedTime, elapsedTime.Nanoseconds())

	if result {
		atomic.AddInt64(&allCountOk, 1)
		allElapsedTimeOk += elapsedTimeNano
		allAvgElapsedTimeOk = allElapsedTimeOk / allCountOk
		if minElapsedTimeOk == 0 || elapsedTimeNano < minElapsedTimeOk {
			minElapsedTimeOk = elapsedTimeNano
		}
		if elapsedTimeNano > maxElapsedTimeOk {
			maxElapsedTimeOk = elapsedTimeNano
		}
		fmt.Printf(" -- success id : %v / insert delay : %v --\n", lastId, insertDelayCount)
	} else {
		atomic.AddInt64(&allCountFail, 1)
		allElapsedTimeFail += elapsedTimeNano
		allAvgElapsedTimeFail = allElapsedTimeFail / allCountFail
		if minElapsedTimeFail == 0 || elapsedTimeNano < minElapsedTimeFail {
			minElapsedTimeFail = elapsedTimeNano
		}
		if elapsedTimeNano > maxElapsedTimeFail {
			maxElapsedTimeFail = elapsedTimeNano
		}
		fmt.Printf(" ============= missing id : %v / insert delay : %v ===========\n", lastId, insertDelayCount)
	}

	atomic.AddInt64(&allCount, 1)
	allElapsedTime += elapsedTimeNano
	allAvgElapsedTime = allElapsedTime / allCount
}

func insertData(newName string) int64 {
	result, err := master.Exec("insert test001 (name) values (?)", newName)
	errorHandle(err)
	//fmt.Printf("result : %+v\n", result)
	lastId, err := result.LastInsertId()
	errorHandle(err)

	return lastId
}

func getData(searchId int64) bool {
	rows, err := slave.Query("select * from test001 where id = ?", searchId)
	errorHandle(err)

	//fmt.Printf("rows %+v\n", rows)

	var (
		id        int
		name      string
		createdAt time.Time
	)

	count := 0
	for rows.Next() {
		count++
		err := rows.Scan(&id, &name, &createdAt)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("id:%v, name:%v createdAt:%v\n", id, name, createdAt)
	}
	if count == 0 {
		//fmt.Printf("============================ %v is missing ============\n", searchId)
	}

	if count > 0 {
		return true
	} else {
		return false
	}
}

func oneThread(task string, count int) {
	for i := 1; i <= count; i++ {
		insertAndSelect(fmt.Sprintf("%v - %v", task, i))
		//time.Sleep(10)
	}
}

func massInsert(threadCount int, insertCount int) {
	startTime := time.Now()
	for i := 0; i < threadCount; i++ {
		go oneThread(fmt.Sprintf("t%v", i), insertCount)
	}
	for true {
		if int(allCount) >= (threadCount * insertCount) {
			elapsedTime := time.Since(startTime)
			tps := float64(allCount) / elapsedTime.Seconds()
			fmt.Printf("================= mass insert end == time : %v / tps : %f ==================\n", elapsedTime, tps)
			break
		}
		time.Sleep(10)
	}
}

func getList() {
	rows, err := slave.Query("SELECT * FROM test001 limit 10")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var (
		id        int
		name      string
		createdAt time.Time
	)

	for rows.Next() {
		err := rows.Scan(&id, &name, &createdAt)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("id:%v, name:%v createdAt:%v\n", id, name, createdAt)
	}
}

func replica_1() {
	initDatabases()

	defer func() {
		master.Close()
		slave.Close()
	}()

	massInsert(300, 10)
	//getData(100)

	//fmt.Scanln()

	fmt.Printf("[ total ] - task count : %v / avg time : %f / tatal time : %f\n", allCount, float64(allAvgElapsedTime)/1000000, float64(allElapsedTime)/1000000)
	fmt.Printf("[ ok ] - task count : %v / avg time : %f / max time : %f / min time : %f\n", allCountOk, float64(allAvgElapsedTimeOk)/1000000, float64(maxElapsedTimeOk)/1000000, float64(minElapsedTimeOk)/1000000)
	fmt.Printf("[ fail ] - task count : %v / avg time : %f / max time : %f / min time : %f\n", allCountFail, float64(allAvgElapsedTimeFail)/1000000, float64(maxElapsedTimeFail)/1000000, float64(minElapsedTimeFail)/1000000)
}

func StudyReplication() {
	replica_1()
}
