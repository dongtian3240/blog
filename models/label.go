package models

import (
	"time"

	"blog/common"

	"github.com/astaxie/beego/orm"
)

/**标签**/
type Label struct {
	LabelId    int `orm:"pk;auto"`
	LabelName  string
	State      int
	UserId     int
	CreateTime time.Time
	UpdateTime time.Time
}

func (this *Label) TableName() string {
	return "t_label"
}

func FindLabelListByPage(pager *common.Pager) error {

	var labels []Label = make([]Label, 0)
	_, err := orm.NewOrm().QueryTable("t_label").Filter("state",
		0).Limit(pager.PageSize,
		(pager.PageNum-1)*pager.PageSize).OrderBy("-update_time").All(&labels)

	if err != nil {
		return err
	}

	count, err := FindLabelListCount()
	if err != nil {
		return err
	}
	pager.DataList = labels

	pager.SetTotalCount(count)
	//获取总数

	return nil
}

func FindLabelListCount() (int, error) {
	count, err := orm.NewOrm().QueryTable("t_label").
		Filter("state", 0).Count()
	if err == nil {
		return int(count), err
	} else {
		return 0, err
	}
}

func NewLabel(label Label) error {

	_, err := orm.NewOrm().Insert(&label)
	return err
}

func FindLabelListById(id int) (Label, error) {

	label := Label{}
	err := orm.NewOrm().QueryTable("t_label").
		Filter("label_id", id).Filter("state", 0).
		One(&label)

	return label, err
}

func DeleteLabelListById(id int) error {

	label := Label{LabelId: id}
	_, err := orm.NewOrm().Delete(&label)

	return err
}

func UpdateLabel(labelId int, labelName string) error {

	label := Label{}
	label.LabelId = labelId
	label.UpdateTime = time.Now()
	label.LabelName = labelName
	_, err := orm.NewOrm().Update(&label, "label_id", "label_name", "update_time")
	return err
}
func FindAllLabel() ([]Label, error) {

	var labels []Label
	_, err := orm.NewOrm().QueryTable("t_label").Filter("state", 0).All(&labels)
	return labels, err
}

func FindLabelListByTopicId(topicId int) ([]Label, error) {
	var labelList []Label
	var tls []TopicLabel
	count, err := orm.NewOrm().QueryTable("t_topic_label").Filter("topic_id", topicId).All(&tls)
	if count == 0 || err != nil {
		return labelList, err
	}
	cnt := len(tls)

	var ids []int = make([]int, cnt)
	for ind, tl := range tls {
		ids[ind] = tl.LabelId
	}

	_, err = orm.NewOrm().QueryTable("t_label").Filter("labelId__in", ids).Filter("state", 0).All(&labelList)

	return labelList, err

}
