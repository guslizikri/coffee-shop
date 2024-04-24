package routers

import (
	"coffee-shop/internal/handlers"
	"coffee-shop/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	repo := repository.NewProduct(d)
	handler := handlers.NewProduct(repo)

	route.GET("/", handler.GetProduct)
	route.POST("/", handler.PostProduct)
	route.PATCH("/:id", handler.PatchProduct)
	route.DELETE("/:id", handler.DeleteDataProduct)

}
