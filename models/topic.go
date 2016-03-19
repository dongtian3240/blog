package models

import (
	"blog/common"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Topic struct {
	TopicId     int `orm:"pk;auto"`
	Title       string
	KeyWords    string
	Content     string
	Author      string
	UserId      int
	State       int
	ViewCount   int
	ReplyCount  int
	CreateTime  time.Time
	UpdateTime  time.Time
	PublishTime time.Time
}

func (this *Topic) TableName() string {
	return "t_topic"
}

func FindTopicListByPage(pager *common.Pager) error {

	var topics []Topic
	_, err := orm.NewOrm().QueryTable("t_topic").Filter("state", 0).
		Limit(pager.PageSize, (pager.PageNum-1)*pager.PageSize).All(&topics)

	if err == nil {

		pager.DataList = topics
		pager.SetTotalCount(FindTopicCount())
	}
	return err
}

func FindTopicCount() int {

	id, err := orm.NewOrm().QueryTable("t_topic").Filter("state", 0).Count()
	if err != nil {
		return 0
	} else {
		return int(id)
	}
}

func NewTopic(topic Topic, labelListStr []string) error {

	ormer := orm.NewOrm()
	//开始一个事务
	ormer.Begin()

	id, err := ormer.Insert(&topic)
	if err != nil {
		return err
	} else {

		for _, va := range labelListStr {

			var topicLabel TopicLabel = TopicLabel{}
			lId, err := strconv.Atoi(va)
			topicLabel.LabelId = lId
			if err != nil {
				return err
			}
			topicLabel.TopicId = int(id)
			_, err = ormer.Insert(&topicLabel)
			if err != nil {
				return err
			}

		}
	}

	defer func() {

		if err != nil {
			ormer.Rollback()
		} else {
			ormer.Commit()
		}
	}()

	return err
}

func FindTopicById(id int) (Topic, error) {

	var topic = Topic{}
	err := orm.NewOrm().QueryTable("t_topic").Filter("topic_id", id).One(&topic)
	return topic, err
}

func UpdateTopic(topic Topic, labelListStr []string) error {

	ormer := orm.NewOrm()
	//开始一个事务
	ormer.Begin()

	//删除分类与博客关联表
	_, err := ormer.QueryTable("t_topic_label").Filter("topic_id", topic.TopicId).Delete()
	if err != nil {
		return err
	}
	topic.UpdateTime = time.Now()
	topic.State = 0
	fmt.Println("===TopicId====", topic.TopicId)
	_, err = ormer.Update(&topic, "state", "update_time", "title", "keywords", "content")
	if err != nil {
		return err
	} else {

		for _, va := range labelListStr {

			var topicLabel TopicLabel = TopicLabel{}
			lId, err := strconv.Atoi(va)
			topicLabel.LabelId = lId
			if err != nil {
				return err
			}
			topicLabel.TopicId = topic.TopicId
			_, err = ormer.Insert(&topicLabel)
			if err != nil {
				return err
			}

		}
	}

	defer func() {

		if err != nil {
			ormer.Rollback()
		} else {
			ormer.Commit()
		}
	}()

	return err
}

func DeleteTopic(topicId int) error {
	ormer := orm.NewOrm()
	ormer.Begin()

	_, err := ormer.Delete(&Topic{TopicId: topicId})
	if err != nil {
		return err

	} else {
		_, err = ormer.QueryTable("t_topic_label").Filter("topic_id", topicId).Delete()

	}

	defer func() {

		if err != nil {
			ormer.Rollback()
		} else {
			ormer.Commit()
		}
	}()

	return err
}
