package handlers

import (
	"coffee-shop/config"
	"coffee-shop/internal/models"
	"coffee-shop/internal/repository"
	"coffee-shop/pkg"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerProduct struct {
	repository.RepoProductIF
}

func NewProduct(r repository.RepoProductIF) *HandlerProduct {
	return &HandlerProduct{r}
}

func (h *HandlerProduct) GetProduct(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")

	limit := ctx.DefaultQuery("limit", "5")
	category := ctx.Query("category")
	search := ctx.Query("search")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	result, err := h.ReadProduct(models.Query{
		Name:     "%" + search + "%",
		Page:     pageInt,
		Limit:    limitInt,
		Category: "%" + category + "%",
	})
	if err != nil {
		pkg.NewRes(http.StatusBadRequest, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, result).Send(ctx)
}

func (h *HandlerProduct) PostProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	product.Image = ctx.MustGet("image").(*string)
	result, err := h.CreateProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	pkg.NewRes(201, result).Send(ctx)

}

func (h *HandlerProduct) PatchProduct(ctx *gin.Context) {
	var product models.Product
	product.Id = ctx.Param("id")
	product.Image = ctx.MustGet("image").(*string)

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	fmt.Println(product.Image)
	fmt.Println(product.Name_product)

	result, err := h.UpdateProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	pkg.NewRes(200, result).Send(ctx)

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

	pkg.NewRes(200, result).Send(ctx)

}
