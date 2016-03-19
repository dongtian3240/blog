package admin

import (
	"blog/common"
	"blog/models"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type TopicController struct {
	BaseController
}

func (this *TopicController) Index() {

	pageNum, _ := this.GetInt("pageNum", 1)
	pageSize, _ := this.GetInt("pageSize", 3)
	pager := common.NewPager()
	pager.PageNum = pageNum
	pager.PageSize = pageSize
	err := models.FindTopicListByPage(pager)
	if err != nil {
		this.Redirect("500.html", http.StatusServiceUnavailable)
		return
	}
	this.Data["pager"] = pager
	this.TplName = "admin/topic/topic.html"

}

func (this *TopicController) ToNewTopic() {
	labels, err := models.FindAllLabel()
	if err != nil || len(labels) == 0 {
		this.Redirect("500.html", http.StatusOK)
		return
	}
	this.Data["labelList"] = labels
	this.TplName = "admin/topic/new.html"
}

func (this *TopicController) NewTopic() {

	result := common.JsonResult{}

	labelIdStrs := this.GetStrings("LabelId", nil)
	fmt.Println("labels", labelIdStrs)
	title := this.GetString("Title", "")
	content := this.GetString("Content", "")

	if labelIdStrs == nil || len(labelIdStrs) == 0 {
		result.Message = "请至少选择一个博客分类！"
	} else if title == "" {
		result.Message = "博客标题不能为空!"
	} else if content == "" {
		result.Message = "博客内容不能为空!"
	} else {

		keyWords := common.DeleteHtml(content)
		fmt.Println("keyword" + keyWords)
		keyWords = common.SubString(keyWords, 0, 250)
		keyWords = strings.Replace(keyWords, " ", "", -1)
		topic := models.Topic{}
		topic.Title = title
		topic.CreateTime = time.Now()
		topic.KeyWords = keyWords
		topic.Content = content
		err := models.NewTopic(topic, labelIdStrs)
		if err != nil {
			result.Message = "保存失败,请重试!"
		} else {

			result.Success = true
			result.Message = "保存成功!"
		}
	}

	this.Ctx.Output.JSON(result, true, false)

}

func (this *TopicController) ToEditTopic() {

	id, err := this.GetInt("id", 0)

	if id == 0 || err != nil {

		this.Redirect("/500.html", http.StatusFound)
		return
	}

	labels, err := models.FindAllLabel()
	if err != nil || len(labels) == 0 {
		this.Redirect("500.html", http.StatusOK)
		return
	}

	//查找当前博客

	topic, err := models.FindTopicById(id)
	if err != nil {
		this.Redirect("500.html", http.StatusOK)
		return
	}

	this.Data["topic"] = topic
	this.Data["labelList"] = labels
	this.TplName = "admin/topic/edit.html"
}

func (this *TopicController) EditTopic() {

	result := common.JsonResult{}
	topicId, err := this.GetInt("topicId", 0)
	labelIdStrs := this.GetStrings("LabelId", nil)
	fmt.Println("labels", labelIdStrs)
	title := this.GetString("Title", "")
	content := this.GetString("Content", "")

	if labelIdStrs == nil || len(labelIdStrs) == 0 {
		result.Message = "请至少选择一个博客分类！"
	} else if title == "" {
		result.Message = "博客标题不能为空!"
	} else if content == "" {
		result.Message = "博客内容不能为空!"
	} else if topicId == 0 || err != nil {
		result.Message = "当前博客不存在"

	} else {

		keyWords := common.DeleteHtml(content)
		keyWords = common.SubString(keyWords, 0, 250)
		keyWords = strings.Replace(keyWords, " ", "", -1)
		topic := models.Topic{}
		topic.TopicId = topicId
		topic.Title = title
		topic.CreateTime = time.Now()
		topic.KeyWords = keyWords
		topic.Content = content
		err := models.UpdateTopic(topic, labelIdStrs)
		if err != nil {
			fmt.Println("err= ", err)
			result.Message = "更新失败,请重试!"
		} else {

			result.Success = true
			result.Message = "保存成功!"
		}
	}

	this.Ctx.Output.JSON(result, true, false)

}

func (this *TopicController) DeleteTopic() {

	result := common.JsonResult{}
	topicId, err := this.GetInt("topicId", 0)
	if topicId == 0 || err != nil {

		result.Message = "此博客不存在!"
	} else {

		err := models.DeleteTopic(topicId)
		if err != nil {
			result.Message = "删除失败请重试！"
		} else {
			result.Message = "删除成功!"
			result.Success = true
		}
	}

	this.Ctx.Output.JSON(result, true, false)

}
