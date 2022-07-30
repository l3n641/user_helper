package models

import (
	"database/sql"
	"strings"
	"time"
)

type Product struct {
	ID           uint `gorm:"primarykey"`
	CreateTime   time.Time
	ModifiedTime time.Time
	DeleteTime   sql.NullTime
	Name         string
	Category     string
	Spu          string
	CollectionId string
	Brand        string
	Size         string
	Price        float64
	DPrice       float64
	IsValidated  int8
	Pictures     string
	Description  string
	ProductAttrs string
	SiteName     string
	SiteUrl      string
}

func (p *Product) GetCategories() []string {
	return strings.Split(p.Category, "|||")
}

type ProductCategory struct {
	Category string
	Quantity int
}

func (p *ProductCategory) GetCategories() []string {
	return strings.Split(p.Category, "|||")
}
