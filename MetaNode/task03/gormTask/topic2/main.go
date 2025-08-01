package main

import (
	"fmt"
	"metanode/task03"
	"metanode/task03/gormTask"
)

func main() {
	db := task03.GormInitDb()
	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	var user gormTask.User
	db.Preload("Post").Preload("Post.Comment").First(&user, 1)
	fmt.Println(user)
	//编写Go代码，使用Gorm查询评论数量最多的文章信息。
	var post gormTask.Post
	db.Preload("Comment").Order("comment_count DESC").First(&post)
	fmt.Println(post)
}
