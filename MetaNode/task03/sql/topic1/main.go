package main

import (
	"log"
	"metanode/task03"
)

// Student /**
type Student struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`  // 学生姓名
	Age   int    `db:"age"`   // 学生年龄
	Grade string `db:"grade"` // 学生年级
}

func main() {
	//初始化数据库
	db := task03.InitDb()
	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	_, err := db.Exec("INSERT INTO students (name, age, grade) VALUES (?, ?, ?)", student.Name, student.Age, student.Grade)
	if err != nil {
		log.Fatal("插入失败！", err)
	}
	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息
	query, err := db.Query("SELECT * FROM students WHERE age > 18")
	if err != nil {
		return
	}
	for query.Next() {
		var student Student
		err := query.Scan(&student.Id, &student.Name, &student.Age, &student.Grade)
		if err != nil {
			return
		}
		// 打印查询结果
		println(student.Name, student.Age, student.Grade)
	}
	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	_, err1 := db.Exec("UPDATE students SET grade = '四年级' WHERE name = '张三'")
	if err1 != nil {
		log.Fatal("更新失败！")
	}
	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	_, err2 := db.Exec("DELETE FROM students WHERE age < 15")
	if err2 != nil {
		log.Fatal("删除失败！")
	}
}
