package routers

import (
	"coffee-shop/internal/handlers"
	"coffee-shop/internal/middleware"
	"coffee-shop/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repository.NewUser(d)
	handler := handlers.NewUser(repo)

	route.GET("/", handler.GetUser)
	route.POST("/", handler.PostUser)
	route.PATCH("/:id", middleware.Auth("user", "admin"), middleware.UploadFile, handler.PatchUser)
	route.DELETE("/:id", middleware.Auth("user", "admin"), handler.DeleteDataUser)

}
