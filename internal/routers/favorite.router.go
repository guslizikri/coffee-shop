package routers

import (
	"coffee-shop/internal/handlers"
	"coffee-shop/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func favorite(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/favorite")

	repo := repository.NewFavorite(d)
	handler := handlers.NewFavorite(repo)

	route.GET("/", handler.GetFavorite)
	route.POST("/", handler.PostFavorite)
	route.DELETE("/:id", handler.DeleteDataFavorite)

}
