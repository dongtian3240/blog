package models

import (
	"time"
)

/****
  评论相关
***/
type Comment struct {
	CommentId  int `orm:"pk;auto"`
	Content    string
	UserId     int
	UserName   string
	State      int
	Type       int
	SourceId   int
	CreateTime time.Time
	UpdateTime time.Time
}

func (this *Comment) TableName() string {
	return "t_comment"
}
