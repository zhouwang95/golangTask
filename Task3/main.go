package main

import (
	"Task3/task03"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *sqlx.DB

// InitMySQL 初始化数据库
func InitMySQL() (err error) {
	dsn := "root:zw123456@tcp(127.0.0.1:3306)/sqlx"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect server failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return
}

func main() {
	/*err := InitMySQL()
	if err != nil {
		return
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)*/
	/*err1 := task02.CreateTableBook(db)
	if err1 != nil {
		return
	}*/
	//task02.InsertBookData(db)

	//task02.QuertyBook(db)

	/*err1 := task02.CreateTable(db)
	if err1 != nil {
		return
	}
	task02.InsertData(db, "张三", "技术部", 15000)
	task02.InsertData(db, "李四", "技术部", 16000)
	task02.InsertData(db, "王五", "研发部", 17000)
	*/
	//task02.QuertyData(db)
	//task02.QuertyMaxSalary(db)

	db, err := gorm.Open(mysql.Open("root:zw123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=true&loc=Local"))
	if err != nil {
		panic(err)
	}
	//task01.Run(db)
	task03.Run(db)
	/*task03.DeleteComm(db)
	task03.Querty1(db)
	task03.Querty2(db)*/

}
