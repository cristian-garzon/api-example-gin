package books

import (
	"api-example/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) GetBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if r := h.Db.First(&book, id); r.Error != nil {
		c.AbortWithError(http.StatusNotFound, r.Error)
		return
	}
	c.JSON(http.StatusOK, &book)
}
