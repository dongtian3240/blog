package common

/***
  分页工具类
**/
type Pager struct {
	PageNum    int
	PageSize   int
	TotalCount int
	TotalPage  int
	DataList   interface{}
}

func NewPager() *Pager {

	return &Pager{PageNum: 1, PageSize: 3}
}
func (this *Pager) SetTotalCount(count int) {
	this.TotalCount = count
	if count%this.PageSize == 0 {
		this.TotalPage = count / this.PageSize
	} else {
		this.TotalPage = count/this.PageSize + 1
	}

}
