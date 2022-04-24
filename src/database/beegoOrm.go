package database

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Test001 struct {
	Id        int
	Name      string    `orm:"size(100)"`
	CreatedAt time.Time `orm:"column(created_at)"`
}

func initDatabase() {

	dbConn := "codebrick:111222qq@tcp(codebrick-test-original.c9s5okqt95x3.ap-southeast-1.rds.amazonaws.com)/tmp?charset=utf8mb4,utf8&sql_mode=TRADITIONAL&multiStatements=true"
	//dbConn += "codebrick:111222qq@tcp(codebrick-test-slave-1.c9s5okqt95x3.ap-southeast-1.rds.amazonaws.com:3306)/tmp?charset=utf8mb4,utf8&sql_mode=TRADITIONAL&multiStatements=true"

	//orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbConn)

	orm.RegisterModel(new(Test001))

}

func StudyBeegoOrm() {
	initDatabase()

	o := orm.NewOrm()

	qs := o.QueryTable("test001")

	var tests []*Test001

	qs.All(&tests)

	for i, test := range tests {
		fmt.Printf("%v, %v\n", i, test)
	}

}
