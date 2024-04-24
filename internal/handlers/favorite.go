package handlers

import (
	"coffee-shop/internal/models"
	"coffee-shop/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerFavorite struct {
	*repository.RepoFavorite
}

func NewFavorite(r *repository.RepoFavorite) *HandlerFavorite {
	return &HandlerFavorite{r}
}

func (h *HandlerFavorite) GetFavorite(ctx *gin.Context) {
	var Favorite models.Favorite
	page := ctx.DefaultQuery("page", "1")

	limit := ctx.DefaultQuery("limit", "5")
	category := ctx.DefaultQuery("category", "")
	search := ctx.DefaultQuery("search", "")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	if err := ctx.ShouldBind(&Favorite); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := h.ReadFavorite(&Favorite, pageInt, limitInt, category, search)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, result)
}

func (h *HandlerFavorite) PostFavorite(ctx *gin.Context) {
	var Favorite models.Favorite

	if err := ctx.ShouldBind(&Favorite); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := h.CreateFavorite(&Favorite)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, result)
}

func (h *HandlerFavorite) DeleteDataFavorite(ctx *gin.Context) {
	var Favorite models.Favorite
	if err := ctx.ShouldBindUri(&Favorite); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	result, err := h.DeleteFavorite(&Favorite)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, result)
}
