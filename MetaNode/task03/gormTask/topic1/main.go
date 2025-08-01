package main

import (
	"metanode/task03"
	"metanode/task03/gormTask"
)

func main() {
	db := task03.GormInitDb(&gormTask.User{}, &gormTask.Post{}, &gormTask.Comment{})
	c := gormTask.Comment{
		Content: "你好",
		Id:      1,
		PostId:  1,
		Status:  0,
	}
	p := gormTask.Post{
		Content: "post1",
		Id:      1,
		UserId:  1,
		Comment: []gormTask.Comment{c},
	}
	u := gormTask.User{
		Id:        1,
		Name:      "user1",
		PostCount: 1,
		Post:      []gormTask.Post{p},
	}
	db.Create(&c)
	db.Create(&p)
	db.Create(&u)
}
