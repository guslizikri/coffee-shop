package routers

import (
	"coffee-shop/internal/handlers"
	"coffee-shop/internal/middleware"
	"coffee-shop/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func favorite(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/favorite")

	repo := repository.NewFavorite(d)
	handler := handlers.NewFavorite(repo)

	route.GET("/", middleware.Auth("user"), handler.GetFavorite)
	route.POST("/", middleware.Auth("user"), handler.PostFavorite)
	route.DELETE("/:id", middleware.Auth("user"), handler.DeleteDataFavorite)

}
