/*
@Time : 2019-07-10 10:20
@Author : zr
*/
package utils

func CountPage(totalCount int64, onePageCount int) int64 {
	if onePageCount == 0 {
		return 0
	}
	page := totalCount / int64(onePageCount)
	remain := totalCount % int64(onePageCount)
	if remain > 0 {
		page = page + 1
	}

	return page
}

func PageIndex2RowOffset(pageIndex int, onePageCount int) int64 {
	if onePageCount == 0 {
		return 0
	}
	pageIndex = pageIndex - 1
	offset := int64(pageIndex) * int64(onePageCount)

	return offset
}
