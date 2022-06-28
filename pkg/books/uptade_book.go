package books

import (
	"api-example/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EditRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h Handler) UpdateBook(c *gin.Context) {
	body := EditRequestBody{}
	id := c.Param("id")
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	var book models.Book

	if r := h.Db.First(&book, id); r.Error != nil {
		c.AbortWithError(http.StatusNotFound, r.Error)
		return
	}

	book.Title = body.Title
	book.Author = body.Author
	body.Description = body.Description

	h.Db.Save(&book)

	c.JSON(http.StatusOK, &book)
}
