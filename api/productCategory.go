package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/user_helper/services"
	"strings"
)

type Category struct {
	Name       string     `json:"label"`
	ParentName string     `json:"parent_name"`
	FullName   string     `json:"full_name"`
	Quantity   int        `json:"quantity"`
	Level      int        `json:"level"`
	Children   []Category `json:"children"`
}

type CategoryTree struct {
	Category
}

func ProductCategory(c *gin.Context) {
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
func getCategoryList(tableName string) []Category {
	var categoryList = []Category{}
	database := services.Database{}
	data := database.GetProductCategory(tableName)
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
					Quantity:   database.GetProductCategoryQuantity(tableName, fullName),
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
