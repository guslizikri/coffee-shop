package handlers

import (
	"coffee-shop/internal/models"
	"coffee-shop/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repository.RepoUser
}

func NewUser(r *repository.RepoUser) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) GetUser(ctx *gin.Context) {
	var user models.User
	page := ctx.DefaultQuery("page", "1")

	limit := ctx.DefaultQuery("limit", "5")
	category := ctx.DefaultQuery("category", "")
	search := ctx.DefaultQuery("search", "")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := h.ReadUser(&user, pageInt, limitInt, category, search)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, result)
}

func (h *HandlerUser) PostUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := h.CreateUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, result)
}

func (h *HandlerUser) PatchUser(ctx *gin.Context) {
	var user models.User
	user.Id = ctx.Param("id")
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	result, err := h.UpdateUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, result)
}

func (h *HandlerUser) DeleteDataUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	result, err := h.DeleteUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, result)
}
