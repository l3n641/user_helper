package services

import (
	"fmt"
	"github.com/user_helper/models"
)

type Product struct {
	Database
	TableName string
}

func (p *Product) GetProductsByCategory(categoryName string) []models.Product {
	var products []models.Product
	db := p.GetDb()
	query := db.Table(p.TableName)
	if categoryName != "" {
		query = query.Where("category like ? ", categoryName+"%")

	}
	query.Find(&products)
	return products
}

func (p *Product) UpdateSizeChartBySpu(size string, spuList []string) {
	db := p.GetDb()
	db.Table(p.TableName).Where("SKU in ?", spuList).Update("SIZE", size)
}
func (p *Product) UpdateSizeChartBySProductIds(size string, productIds []int) {
	db := p.GetDb()
	db.Table(p.TableName).Where("id in ?", productIds).Update("size", size)
}
func (p *Product) UpdateCategoryChartBySProductIds(category string, productIds []int) {
	db := p.GetDb()
	db.Table(p.TableName).Where("id in ?", productIds).Update("category", category)
}
func (p *Product) DeleteEmptyData(field string) {
	db := p.GetDb()
	db.Table(p.TableName).Where("? = '' OR ? IS NULL", field, field).Delete(&models.Product{})
}
func (p *Product) DeleteProductById(productIds []int) {
	db := p.GetDb()
	db.Table(p.TableName).Where("id in ?", productIds).Delete(&models.Product{})
}

func (p *Product) RenameProductCategory(oldCategory, newCategory string) {
	db := p.GetDb()
	sql := fmt.Sprintf(" UPDATE %s SET category= REPLACE(category,?,?) WHERE category LIKE ?", p.TableName)
	db.Exec(sql, oldCategory, newCategory, oldCategory+"%")
}

func (p *Product) GetProductCategory() []models.ProductCategory {
	var categoryList []models.ProductCategory
	db := p.GetDb()
	sql := fmt.Sprintf("select COUNT(category) as quantity,category from %s GROUP BY category", p.TableName)
	db.Raw(sql).Scan(&categoryList)
	return categoryList
}
func (p *Product) GetProductCategoryQuantity(categoryName string) int64 {
	var categoryQuantity int64
	db := p.GetDb()
	db.Table(p.TableName).Where("category like ?", categoryName+"%").Count(&categoryQuantity)

	return categoryQuantity
}
