package handlers

import (
	"coffee-shop/config"
	"coffee-shop/internal/models"
	"coffee-shop/internal/repository"
	"coffee-shop/pkg"
	"log"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	repository.RepoUserIF
}

func NewUser(r repository.RepoUserIF) *HandlerUser {
	return &HandlerUser{r}
}
func (h *HandlerUser) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := h.GetUserById(id)
	if err != nil {
		log.Println(err)
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, result).Send(ctx)

}

func (h *HandlerUser) GetUser(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "5")
	search := ctx.Query("search")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	result, err := h.ReadUser(models.Query{
		Name:  "%" + search + "%",
		Page:  pageInt,
		Limit: limitInt})
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pkg.NewRes(200, result).Send(ctx)

}

func (h *HandlerUser) PostUser(ctx *gin.Context) {
	var err error
	var user models.User
	user.Role = "user"

	if err = ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}
	user.Password, err = pkg.HashPass(user.Password)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Message: err.Error(),
		}).Send(ctx)
		return
	}

	result, err := h.CreateUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	pkg.NewRes(200, result).Send(ctx)
}

func (h *HandlerUser) PatchUser(ctx *gin.Context) {
	var user models.User
	user.Id = ctx.Param("id")
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	user.Image = ctx.MustGet("image").(*string)
	result, err := h.UpdateUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pkg.NewRes(200, result).Send(ctx)

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

	pkg.NewRes(200, result).Send(ctx)

}
