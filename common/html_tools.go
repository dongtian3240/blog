package common

import (
	"regexp"
	"strings"
)

/**定义去掉的正则**/
var (
	htmlPatten   string = "<[^>]+>"
	scriptPatten string = "<script[^>]*?>[\\s\\S]*?<\\/script>"
	cssPatten    string = "<style[^>]*?>[\\s\\S]*?<\\/style>"
)

/**
 过滤掉html代码 css代码 和script代码
**/
func DeleteHtml(html string) string {

	//去掉html
	reg := regexp.MustCompile(htmlPatten)
	strs := reg.FindAllString(html, -1)
	for _, va := range strs {

		html = strings.Replace(html, va, "", -1)

	}

	//去掉script
	reg = regexp.MustCompile(scriptPatten)
	strs = reg.FindAllString(html, -1)
	for _, va := range strs {

		html = strings.Replace(html, va, "", -1)

	}

	reg = regexp.MustCompile(cssPatten)
	strs = reg.FindAllString(html, -1)
	for _, va := range strs {

		html = strings.Replace(html, va, "", -1)

	}

	return html
}

/**字符串截取**/
func SubString(str string, start, length int) string {

	str = strings.Trim(str, " ")
	rs := []rune(str)
	le := len(rs)
	end := start + length

	if start > le || start < 0 {
		start = 0
	}
	if end > le {
		end = le
	}
	return string(rs[start:end])
}
