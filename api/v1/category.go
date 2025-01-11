package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_mysql/service"
)

func ListCategory(c *gin.Context) {
	var listCategory service.CategoryService
	if err := c.ShouldBind(&listCategory); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		res := listCategory.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}
