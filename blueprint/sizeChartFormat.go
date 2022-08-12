package blueprint

import (
	"github.com/gin-gonic/gin"
	"github.com/user_helper/services"
	"net/http"
)

func Sizechat(c *gin.Context) {
	database := services.Database{}
	tables := database.GetTables()
	c.HTML(http.StatusOK, "sizechart_format.html", gin.H{"tables": tables})
}
