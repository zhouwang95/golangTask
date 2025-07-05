package task02

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

/*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     int
}

func CreateTable(db *sqlx.DB) (err error) {
	sqlStr := `
		create table if not exists employees (
		    id bigint primary key auto_increment,
		    name varchar(20),
		    department varchar(100),
		    salary bigint
		);
	`
	_, err = db.Exec(sqlStr)
	return err
}

func InsertData(db *sqlx.DB, name string, department string, salary int64) (err error) {
	_, err = db.Exec("insert into employees(name,department,salary) values (?,?,?)",
		name, department, salary)
	return err
}

func QuertyData(db *sqlx.DB) {
	sqlQuerty := "select * from employees where department = ?"
	var employee []Employee
	err := db.Select(&employee, sqlQuerty, "技术部")
	if err != nil {
		return
	}
	fmt.Println(employee)
}

func QuertyMaxSalary(db *sqlx.DB) {
	sql := "select * from employees order by salary desc limit 1"
	var e Employee
	err := db.Get(&e, sql)
	if err != nil {
		return
	}
	fmt.Println(e)
}
