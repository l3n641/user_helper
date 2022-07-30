package api

import (
	"github.com/gin-gonic/gin"
	"github.com/user_helper/services"
)

func GetTables(c *gin.Context) {
	database := services.Database{}
	tables := database.GetTables()
	c.JSON(200, tables)
}
