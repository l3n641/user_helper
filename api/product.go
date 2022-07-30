package api

import (
	"github.com/gin-gonic/gin"
	"github.com/user_helper/models"
	"github.com/user_helper/services"
	"net/http"
	"strings"
)

func GetProductAll(c *gin.Context) {
	var products []models.Product
	tableName, ok := c.GetQuery("table_name")
	if !ok {
		c.JSON(400, "table not found")
	}
	tableName = strings.TrimSpace(tableName)

	productSrv := services.Product{TableName: tableName}
	db := productSrv.GetDb()
	categoryName, ok := c.GetQuery("category_name")
	if !ok {
		db.Table(tableName).Find(&products)
	} else {
		products = productSrv.GetProductsByCategory(categoryName)
	}
	c.JSON(200, products)
}

func DeleteProduct(c *gin.Context) {
	type Data struct {
		TableName  string `form:"table_name" json:"table_name"  binding:"required"`
		EmptyField string `form:"empty_field" json:"empty_field"`
	}

	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database := services.Database{}
	if data.EmptyField != "" {
		database.DeleteEmptyData(data.TableName, data.EmptyField)
	}
	c.JSON(200, "success")
}

func RenameCategory(c *gin.Context) {
	type Data struct {
		TableName   string `form:"table_name" json:"table_name"  binding:"required"`
		OldCategory string `form:"old_category" json:"old_category"  binding:"required"`
		NewCategory string `form:"new_category" json:"new_category"  binding:"required"`
	}

	var data Data

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := services.Database{}
	database.RenameProductCategory(data.TableName, data.OldCategory, data.NewCategory)
	c.JSON(200, "success")

}
