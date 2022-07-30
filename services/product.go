package services

import "github.com/user_helper/models"

type Product struct {
	Database
	TableName string
}

func (p *Product) GetProductsByCategory(categoryName string) []models.Product {
	var products []models.Product
	db := p.GetDb()
	db.Table(p.TableName).Where("category like ? ", categoryName+"%").Find(&products)
	return products
}
