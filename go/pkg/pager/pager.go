package pager

import (
	"fmt"

	"github.com/hyahm/golog"
)

// 返回正确的起始值,和末尾数, 总数，页数，显示数
func GetPagingLimitAndPage(count int, page int, limit int) (int, int) {
	// 都小于1了
	if limit == 0 {
		return 0, 0
	}
	if page < 1 {
		return 1, count % limit
	}
	fmt.Println(444444)
	golog.Info(11111)
	// 超出了，返回最大的页码
	if page*limit > count+limit {
		golog.Info(33333)
		if count%limit == 0 {
			return ((count / limit) - 1) * limit, limit
		} else {
			return (count/limit + 1) * limit, count % limit
		}
	} else {
		// if count%limit == 0 {
		golog.Info(22222)
		start := (page - 1) * limit
		golog.Info(start)

		if count-start < limit {
			return (page - 1) * limit, count - start
		} else {
			return (page - 1) * limit, limit
		}

	}
}
