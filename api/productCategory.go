package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/user_helper/services"
	"net/http"
	"strings"
)

type Category struct {
	Name       string     `json:"label"`
	ParentName string     `json:"parent_name"`
	FullName   string     `json:"full_name"`
	Quantity   int64      `json:"quantity"`
	Level      int        `json:"level"`
	Children   []Category `json:"children"`
}

type CategoryTree struct {
	Category
}

func GetProductCategory(c *gin.Context) {
	tableName, ok := c.GetQuery("table_name")
	if !ok {
		c.JSON(400, "table not found")
	}
	tableName = strings.TrimSpace(tableName)
	data := getCategoryList(tableName)
	result := tree(data, "root")
	fmt.Println(result)
	c.JSON(200, result)

}

func UpdateCategoryByProductId(c *gin.Context) {
	type Data struct {
		TableName  string `form:"table_name" json:"table_name"  binding:"required"`
		ProductIds []int  `form:"product_ids" json:"product_ids"  binding:"required"`
		Category   string `form:"category" json:"category"  binding:"required"`
	}
	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productSrv := services.Product{TableName: data.TableName}
	productSrv.UpdateCategoryChartBySProductIds(data.Category, data.ProductIds)
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

	productSrv := services.Product{TableName: data.TableName}
	productSrv.RenameProductCategory(data.OldCategory, data.NewCategory)
	c.JSON(200, "success")

}

func getCategoryList(tableName string) []Category {
	var categoryList = []Category{}
	productSrv := services.Product{TableName: tableName}

	data := productSrv.GetProductCategory()
	for _, category := range data {
		fullName := ""
		parentName := "root"
		cName := category.GetCategories()
		for i, name := range cName {
			if fullName == "" {
				fullName = name
			} else {
				fullName = fullName + "|||" + name
			}
			index := getNodeIndex(fullName, categoryList)
			if index == -1 {
				item := Category{
					Name:       name,
					ParentName: parentName,
					FullName:   fullName,
					Quantity:   productSrv.GetProductCategoryQuantity(fullName),
					Level:      i,
				}
				categoryList = append(categoryList, item)
			}
			parentName = name
		}
	}
	return categoryList
}

func getNodeIndex(name string, categories []Category) int {

	for i := 0; i < len(categories); i++ {
		if categories[i].FullName == name {
			return i
		}
	}
	return -1

}

func tree(list []Category, parentName string) []Category {
	var treeList []Category
	for _, data := range list {
		if data.ParentName == parentName {
			data.Children = tree(list, data.Name)
			treeList = append(treeList, data)
		}
	}
	return treeList
}
