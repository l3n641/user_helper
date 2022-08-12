package services

import (
	"fmt"
	"github.com/user_helper/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
)

type Database struct {
	Dsn       string
	DbName    string
	dbConnect *gorm.DB
}

func (d *Database) GetTables() []models.Table {
	var tables []models.Table
	db := d.GetDb()
	sql := fmt.Sprintf("SELECT * FROM INFORMATION_SCHEMA.TABLES  WHERE table_schema  = '%s'", d.DbName)
	db.Raw(sql).Scan(&tables)
	return tables
}

func (d *Database) GetDb() *gorm.DB {
	if d.dbConnect != nil {
		return d.dbConnect
	}
	dsn := d.getDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	dbDebug := os.Getenv("DB_DEBUG")
	d.dbConnect = db

	if dbDebug != "" {
		return db.Debug()
	}
	return db
}

func (d *Database) getDsn() string {
	username := os.Getenv("DB_USER")                // 账号
	password := os.Getenv("DB_PASSWORD")            // 密码
	host := os.Getenv("DB_HOST")                    // 数据库地址，可以是Ip或者域名
	port, err := strconv.Atoi(os.Getenv("DB_PORT")) // 数据库端口
	if err != nil {
		panic("数据库端口错误")
	}
	dbname := os.Getenv("DB_NAME") // 数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	d.Dsn = dsn
	d.DbName = dbname
	return dsn
}
