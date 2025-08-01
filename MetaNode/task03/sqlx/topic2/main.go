package main

import "metanode/task03"

type Book struct {
	Id     int    `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Price  int    `db:"price"`
}

func main() {
	db := task03.SqlXInitDb()
	//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，
	//并将结果映射到 Book 结构体切片中，确保类型安全。
	var books []Book
	db.Select(&books, "SELECT * FROM books WHERE price > ?", 50)
}
