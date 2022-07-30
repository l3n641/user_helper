package api

import (
	"github.com/gin-gonic/gin"
	"github.com/user_helper/services"
	"net/http"
)

type SizeChart struct {
	TableName string   `form:"table_name" json:"table_name"  binding:"required"`
	SpuList   []string `form:"spu_list" json:"spu_list"  binding:"required"`
	SizeChart string   `form:"size_chart" json:"size_chart"  binding:"required"`
}

func UpdateSizeChartBySpu(c *gin.Context) {
	var data SizeChart
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&data); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database := services.Database{}
	database.UpdateSizeChartBySpu(data.TableName, data.SizeChart, data.SpuList)
	c.JSON(200, "success")
}
