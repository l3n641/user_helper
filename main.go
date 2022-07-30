package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/user_helper/api"
	"github.com/user_helper/blueprint"
	"log"
	"net/http"
	"os"
)

var DbName string
var DbUserName string
var DbPassword string
var DbHost = "127.0.0.1"
var DbPort = 3306
var IsDebug = false

func init() {

}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	image_path := os.Getenv("IMAGE_PATH")
	if image_path == "" {
		panic("图片目录还未设置")
	}

	r := gin.Default()
	r.Use(Cors())
	r.LoadHTMLGlob("templates/*")
	r.GET("/", blueprint.Sizechat)
	r.GET("/table", api.GetTables)
	r.GET("/product", api.GetProductAll)
	r.GET("/product_category", api.ProductCategory)
	r.PUT("/size_chart", api.UpdateSizeChartBySpu)
	r.POST("/delete_product", api.DeleteProduct)
	r.PUT("/rename_category", api.RenameCategory)
	r.Static("/image", image_path)
	r.Static("/js", "./static/js")
	r.Static("/css", "./static/css")
	r.Run(":8000")
}
