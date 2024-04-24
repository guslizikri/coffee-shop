package handlers

import (
	"coffee-shop/internal/models"
	"coffee-shop/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerProduct struct {
	*repository.RepoProduct
}

func NewProduct(r *repository.RepoProduct) *HandlerProduct {
	return &HandlerProduct{r}
}

func (h *HandlerProduct) GetProduct(ctx *gin.Context) {
	var Product models.Product
	page := ctx.DefaultQuery("page", "1")

	limit := ctx.DefaultQuery("limit", "5")
	category := ctx.DefaultQuery("category", "")
	search := ctx.DefaultQuery("search", "")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	if err := ctx.ShouldBind(&Product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := h.ReadProduct(&Product, pageInt, limitInt, category, search)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, result)
}

func (h *HandlerProduct) PostProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := h.CreateProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, result)
}

func (h *HandlerProduct) PatchProduct(ctx *gin.Context) {
	var product models.Product
	product.Id = ctx.Param("id")
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	result, err := h.UpdateProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, result)
}

func (h *HandlerProduct) DeleteDataProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindUri(&product); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	result, err := h.DeleteProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, result)
}
