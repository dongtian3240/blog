package models

/****
  微博和标签关联
**/
type TopicLabel struct {
	Id      int `orm:"pk;auto"`
	TopicId int
	LabelId int
}

func (this *TopicLabel) TableName() string {
	return "t_topic_label"
}
