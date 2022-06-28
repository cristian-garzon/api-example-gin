package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	Db *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := Handler{
		Db: db,
	}

	routes := r.Group("/books")
	routes.GET("/", h.GetBooks)
	routes.GET("/:id", h.GetBook)
	routes.POST("/", h.addBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)

	test := r.Group("/example")

	test.GET("/", h.HelloWorld)

}
