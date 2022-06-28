package books

import (
	"api-example/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  models.Book
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (h Handler) addBook(c *gin.Context) {
	body := AddBookRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var book models.Book
	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description
	if r := h.Db.Create(&book); r.Error != nil {
		c.AbortWithError(http.StatusNotFound, r.Error)
		return
	}
	c.JSON(http.StatusCreated, &book)
}
