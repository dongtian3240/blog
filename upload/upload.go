package upload

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) UploadImage() {

	file, fileHeader, err := this.GetFile("imgFile")
	var result map[string]interface{} = make(map[string]interface{}, 0)
	if err != nil {

		result["error"] = 1
		result["message"] = "图片不存在!"

	} else if isImage(fileHeader.Header.Get("Content-Type")) == false {
		result["error"] = 1
		result["message"] = "非法的图片!"
	} else {

		//创建目录
		now := time.Now()
		filePath := "/public/" + now.Format("2006") + "/" + now.Format("01") + "/" + now.Format("02")
		wd, _ := os.Getwd()
		fileFullPath := wd + filePath

		if existDir(fileFullPath) == false {
			os.MkdirAll(fileFullPath, os.ModePerm)
		}
		shortPath := now.Format("150405") + "_" + fileHeader.Filename
		path := fileFullPath + "/" + shortPath
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
		fmt.Println("err", err)
		_, err = io.Copy(f, file)
		fmt.Println("err1", err)
		defer f.Close()
		defer file.Close()
		if err != nil {
			fmt.Println(err)
			result["error"] = 1
			result["message"] = "上传失败!"
		} else {

			result["error"] = 0
			result["url"] = filePath + "/" + shortPath
			result["message"] = "上传成功!"
		}
	}

	this.Ctx.Output.JSON(result, true, false)
}

func isImage(contentType string) bool {
	if contentType == "image/x-icon" ||
		contentType == "image/jpeg" ||
		contentType == "application/x-jpg" ||
		contentType == "image/gif" ||
		contentType == "application/x-png" ||
		contentType == "image/png" ||
		contentType == "application/x-bmp" {
		return true
	} else {
		return false
	}
}

func existDir(dir string) bool {

	_, err := os.Stat(dir)
	return err == nil || os.IsExist(err)
}
