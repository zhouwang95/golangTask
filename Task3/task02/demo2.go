package task02

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

/*
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，
并将结果映射到 Book 结构体切片中，确保类型安全。
*/
type Book struct {
	ID     int
	Title  string
	Author string
	Price  uint
}

func CreateTableBook(db *sqlx.DB) (err error) {
	sqlStr := `
		create table if not exists books (
		    id bigint primary key auto_increment,
		    title varchar(100),
		    author varchar(20),
		    price int
		);
	`
	_, err = db.Exec(sqlStr)
	return err
}

func InsertBookData(db *sqlx.DB) {
	db.Exec("insert into books(title,author,price) values (?,?,?)",
		"go语言", "张三", 100)
	db.Exec("insert into books(title,author,price) values (?,?,?)",
		"java语言", "李四", 30)
	db.Exec("insert into books(title,author,price) values (?,?,?)",
		"c语言", "王五", 100)

}

func QuertyBook(db *sqlx.DB) {
	var book []Book
	err := db.Select(&book,"select * from books where price > ?",50)
	if err != nil {
		return
	}
	fmt.Println(book)
}
