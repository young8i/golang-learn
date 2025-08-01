package main

import (
	"gorm.io/gorm"
	"metanode/task03"
	"metanode/task03/gormTask"
)

func main() {
	db := task03.GormInitDb()
	//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	db.Callback().Create().After("gorm:create").Register("update_post_count", func(db *gorm.DB) {
		var user gormTask.User
		db.Model(&user).Where("id = ?", db.Statement.Dest.(gormTask.Post).UserId).Update("post_count", gormTask.User{PostCount: db.Statement.Dest.(gormTask.Post).UserId})
	})
	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	db.Callback().Delete().After("gorm:delete").Register("update_comment_status", func(db *gorm.DB) {
		var comment gormTask.Comment
		db.Model(&comment).Where("id = ?", db.Statement.Dest.(gormTask.Comment).PostId).Update("status", "无评论")
	})
}
