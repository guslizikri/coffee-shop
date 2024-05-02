package handlers

import (
	"coffee-shop/config"
	"coffee-shop/internal/repository"
	"coffee-shop/pkg"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	Email    string `db:"email" json:"email" form:"email"`
	Password string `db:"password" json:"password"`
}

type HandlerAuth struct {
	*repository.RepoUser
}

func NewAuth(r *repository.RepoUser) *HandlerAuth {
	return &HandlerAuth{r}
}

func (h *HandlerAuth) Login(ctx *gin.Context) {
	var data Auth

	if err := ctx.ShouldBind(&data); err != nil {
		pkg.NewRes(500, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	users, err := h.GetAuthData(data.Email)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	if err := pkg.VerifyPass(users.Password, data.Password); err != nil {
		pkg.NewRes(401, &config.Result{
			Data: "Password salah",
		}).Send(ctx)
		return
	}

	jwtt := pkg.NewToken(users.Id, users.Role)
	tokens, err := jwtt.Generate()
	if err != nil {
		pkg.NewRes(500, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, &config.Result{Data: tokens}).Send(ctx)
}
