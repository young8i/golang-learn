package gormTask

type User struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	PostCount int
	Post      []Post
}
type Post struct {
	Id           int `gorm:"primaryKey"`
	Content      string
	UserId       int
	CommentCount int
	Comment      []Comment
}
type Comment struct {
	Id      int `gorm:"primaryKey"`
	Content string
	PostId  int
	Status  int
}
