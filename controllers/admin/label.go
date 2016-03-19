package admin

import (
	"blog/common"
	"blog/models"
	"fmt"
	"net/http"
	"time"
)

type LabelController struct {
	BaseController
}

func (this *LabelController) Index() {

	pageNum, _ := this.GetInt("pageNum", 1)
	pageSize, _ := this.GetInt("pageSize", 3)
	pager := common.NewPager()
	pager.PageNum = pageNum
	pager.PageSize = pageSize
	err := models.FindLabelListByPage(pager)
	if err != nil {
		this.Redirect("500.html", http.StatusServiceUnavailable)
		return
	}
	fmt.Println("err", err)
	fmt.Println("pager is ", pager)
	this.Data["pager"] = pager
	this.TplName = "admin/label/label.html"
}

func (this *LabelController) NewLabel() {

	result := common.JsonResult{}

	labelName := this.GetString("labelName", "")

	if labelName == "" {
		result.Message = "分类名称不能为空"
	} else {
		label := models.Label{}
		label.LabelName = labelName
		label.State = 0
		label.CreateTime = time.Now()
		label.UpdateTime = time.Now()
		label.UserId = 1
		err := models.NewLabel(label)
		if err != nil {
			result.Message = "保存失败，请重试！"
		} else {
			result.Message = "保存成功"
			result.Success = true
		}
	}
	this.Ctx.Output.JSON(result, true, false)

}

func (this *LabelController) FindLabelById() {

	result := common.JsonResult{}
	id, err := this.GetInt("id", 0)
	if err != nil || id == 0 {
		result.Message = "当前记录不存在！"
	} else {

		label, err := models.FindLabelListById(id)
		if err != nil || label.LabelId == 0 {

			result.Message = "当前记录不存在！"
		} else {
			result.Data = label
			result.Message = "成功!"
			result.Success = true
		}
	}

	this.Ctx.Output.JSON(result, true, false)
}

func (this *LabelController) UpdateLabel() {

	result := common.JsonResult{}

	labelId, err := this.GetInt("labelId", 0)
	labelName := this.GetString("labelName", "")
	if labelId == 0 || err != nil {
		result.Message = "当前记录不存在"
	} else if labelName == "" {
		result.Message = "分类名称不能为空!"
	} else {

		err := models.UpdateLabel(labelId, labelName)
		if err != nil {
			result.Message = "修改失败，请重试！"
		} else {
			result.Message = "修改成功"
			result.Success = true
		}

	}

	this.Ctx.Output.JSON(result, true, false)
}

func (this *LabelController) DeleteLabelById() {
	result := common.JsonResult{}
	id, err := this.GetInt("id", 0)
	if err != nil || id == 0 {
		result.Message = "当前记录不存在！"
	} else {

		err := models.DeleteLabelListById(id)
		if err != nil {

			result.Message = "删除失败,请重试！"
		} else {
			result.Message = "删除成功!"
			result.Success = true
		}
	}

	this.Ctx.Output.JSON(result, true, false)
}

func (this *LabelController) FindLabelListByTopicId() {

	result := common.JsonResult{}
	topicId, err := this.GetInt("topicId", 0)

	if topicId == 0 || err != nil {
		result.Message = "此博客不存在!"
	} else {
		labelList, err := models.FindLabelListByTopicId(topicId)
		if err != nil {
			result.Message = "此博客分类不存在!"
		} else {
			result.Data = labelList
			result.Success = true
		}
	}

	this.Ctx.Output.JSON(result, true, false)

}
