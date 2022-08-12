package api

import (
	"github.com/gin-gonic/gin"
	"github.com/user_helper/models"
	"github.com/user_helper/services"
	"net/http"
	"strings"
)

func GetProduct(c *gin.Context) {
	var products []models.Product
	tableName, ok := c.GetQuery("table_name")
	if !ok {
		c.JSON(400, "table not found")
	}
	tableName = strings.TrimSpace(tableName)

	productSrv := services.Product{TableName: tableName}
	categoryName, ok := c.GetQuery("category_name")
	products = productSrv.GetProductsByCategory(categoryName)
	c.JSON(200, products)
}

func DeleteProduct(c *gin.Context) {
	type Data struct {
		TableName  string `form:"table_name" json:"table_name"  binding:"required"`
		ProductIds []int  `form:"product_ids" json:"product_ids"  binding:"required"`
	}

	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productSrv := services.Product{TableName: data.TableName}
	productSrv.DeleteProductById(data.ProductIds)
	c.JSON(200, "success")
}

func DeleteEmptyProduct(c *gin.Context) {
	type Data struct {
		TableName  string `form:"table_name" json:"table_name"  binding:"required"`
		EmptyField string `form:"empty_field" json:"empty_field" binding:"required"`
	}

	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productSrv := services.Product{TableName: data.TableName}
	productSrv.DeleteEmptyData(data.EmptyField)

	c.JSON(200, "success")
}
