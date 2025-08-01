package main

import "metanode/task03"

// employees 表，包含字段 id 、 name 、 department 、 salary
type Employee struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func main() {
	db := task03.SqlXInitDb()
	//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	var employees []Employee
	db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	var employee Employee
	db.Get(&employee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
}
