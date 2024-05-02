package routers

import (
	"coffee-shop/internal/handlers"
	"coffee-shop/internal/middleware"
	"coffee-shop/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	repo := repository.NewProduct(d)
	handler := handlers.NewProduct(repo)

	route.GET("/", handler.GetProduct)
	route.POST("/", middleware.Auth("admin"), middleware.UploadFile, handler.PostProduct)
	route.PATCH("/:id", middleware.Auth("admin"), middleware.UploadFile, handler.PatchProduct)
	route.DELETE("/:id", middleware.Auth("admin"), handler.DeleteDataProduct)

}
