package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context, pageCurrent *int, pageSize *int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		p := c.DefaultQuery("page", "0")
		l := c.DefaultQuery("limit", "10")

		page, _ := strconv.Atoi(p)
		if page == 0 {
			page = 1
		}

		limit, _ := strconv.Atoi(l)
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		offset := (page - 1) * limit

		*pageCurrent = page
		*pageSize = limit

		return db.Offset(offset).Limit(limit)
	}
}
