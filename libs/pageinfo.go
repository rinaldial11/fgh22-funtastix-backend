package libs

import "math"

type PageInfo struct {
	CurrentPage int `json:"currentPage"`
	NextPage    int `json:"nextPage"`
	PrevPage    int `json:"prevPage"`
	TotalPage   int `json:"totalPage"`
	TotalData   int `json:"totalData"`
}

func GetPageInfo(page int, limit int, count int) PageInfo {
	var pageInfo PageInfo

	totalPage := int(math.Ceil(float64(count) / float64(limit)))

	nextPage := page + 1
	if nextPage > totalPage {
		nextPage = totalPage
	}

	prevPage := page - 1
	if page == 1 {
		prevPage = 1
	}

	pageInfo.CurrentPage = page
	pageInfo.NextPage = nextPage
	pageInfo.PrevPage = prevPage
	pageInfo.TotalPage = totalPage
	pageInfo.TotalData = count

	return pageInfo
}
