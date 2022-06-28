package books

import (
	"api-example/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) GetBooks(c *gin.Context) {
	var books []models.Book
	if r := h.Db.Find(&books); r.Error != nil {
		c.AbortWithError(http.StatusNotFound, r.Error)
		return
	}
	c.JSON(http.StatusOK, &books)
}
