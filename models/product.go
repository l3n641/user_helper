package models

import (
	"database/sql"
	"strings"
)

type Product struct {
	ID            uint         `json:"id" gorm:"primarykey"`
	FeaturedImage string       `gorm:"column:featured_image;type:text" json:"featured_image"`
	LANG          string       `gorm:"column:LANG;type:varchar(16)" json:"LANG"`
	CAT0          string       `gorm:"column:CAT-0;type:varchar(255)" json:"CAT-0"`
	Category      string       `gorm:"column:Category;type:varchar(255)" json:"Category"`
	SIZE          string       `gorm:"column:SIZE;type:varchar(512)" json:"SIZE"`
	SKU           string       `gorm:"column:SKU;type:varchar(255)" json:"SKU"`
	StyleName     string       `gorm:"column:Style-Name;type:varchar(255)" json:"Style-Name"`
	TITLE         string       `gorm:"column:TITLE;type:varchar(512)" json:"TITLE"`
	Brand         string       `gorm:"column:Brand;type:varchar(64)" json:"Brand"`
	BrandName     string       `gorm:"column:Brand-name;type:varchar(64)" json:"Brand-name"`
	Model         string       `gorm:"column:model;type:varchar(255)" json:"model"`
	Type          string       `gorm:"column:Type;type:varchar(255)" json:"Type"`
	Gender        string       `gorm:"column:Gender;type:varchar(255)" json:"Gender"`
	GenderName    string       `gorm:"column:Gender-name;type:varchar(255)" json:"Gender-name"`
	Color         string       `gorm:"column:Color;type:varchar(255)" json:"Color"`
	ColorName     string       `gorm:"column:Color-Name;type:varchar(255)" json:"Color-Name"`
	Desc          string       `gorm:"column:desc;type:text" json:"desc"`
	Desc2         string       `gorm:"column:desc2;type:text" json:"desc2"`
	Price         float64      `gorm:"column:price;type:decimal(10,2)" json:"price"`
	Price2        float64      `gorm:"column:price2;type:decimal(10,2)" json:"price2"`
	Description   string       `gorm:"column:Description;type:varchar(255)" json:"Description"`
	Keyword       string       `gorm:"column:Keyword;type:varchar(255)" json:"Keyword"`
	IMGAdd        string       `gorm:"column:IMG-Add;type:varchar(255)" json:"IMG-Add"`
	NPrice        float64      `gorm:"column:NPrice;type:decimal(10,2)" json:"NPrice"`
	OPrice        float64      `gorm:"column:OPrice;type:decimal(10,2)" json:"OPrice"`
	Max           float64      `gorm:"column:max;type:decimal(10)" json:"max"`
	Min           float64      `gorm:"column:min;type:decimal(10)" json:"min"`
	NSize         string       `gorm:"column:NSize;type:varchar(64)" json:"NSize"`
	Date          sql.NullTime `gorm:"column:Date;type:datetime" json:"Date"`
	PageUrl       string       `gorm:"column:PageUrl;type:varchar(512)" json:"PageUrl"`
}

func (p *Product) GetCategories() []string {
	return strings.Split(p.Category, "->")
}

type ProductCategory struct {
	Category string
	Quantity int
}

func (p *ProductCategory) GetCategories() []string {
	return strings.Split(p.Category, "->")
}
