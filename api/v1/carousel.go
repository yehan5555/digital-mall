package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_mysql/service"
)

func ListCarousels(c *gin.Context) {
	var listCarousels service.CarouselService
	if err := c.ShouldBind(&listCarousels); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		res := listCarousels.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}
