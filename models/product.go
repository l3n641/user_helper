package models

import (
	"database/sql"
	"strings"
	"time"
)

type Product struct {
	ID           uint         `json:"id" gorm:"primarykey"`
	CreateTime   time.Time    `json:"create_time"`
	ModifiedTime time.Time    `json:"modified_time"`
	DeleteTime   sql.NullTime `json:"delete_time"`
	Name         string       `json:"name"`
	Category     string       `json:"category"`
	Spu          string       `json:"spu"`
	CollectionId string       `json:"collection_id"`
	Brand        string       `json:"brand"`
	Size         string       `json:"size"`
	Price        float64      `json:"price"`
	DPrice       float64      `json:"d_price"`
	IsValidated  int8         `json:"is_validated"`
	Pictures     string       `json:"pictures"`
	Description  string       `json:"description"`
	ProductAttrs string       `json:"product_attrs"`
	SiteName     string       `json:"site_name"`
	SiteUrl      string       `json:"site_url"`
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
