package main

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func SayHello(name string, err bool) (string, error) {
	if err {
		return "", errors.New("something wrong")
	}
	return "Hello " + name, nil
}

func example() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/query", queryHandler)
	router.GET("/param/:id", paramHandler)
	router.POST("/body", bodyHandler)
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.String(200, "Hello World")
	// })

	router.Run(":3001")
}

func rootHandler(ctx *gin.Context) {
	ctx.String(200, "Hello World")
}

type qString struct {
	Page  string `form:"page"`
	Limit string `form:"limit"`
}

func queryHandler(ctx *gin.Context) {
	// page := ctx.Query("page")
	// limit := ctx.Query("limit")
	var query qString

	err := ctx.ShouldBind(&query)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(200, gin.H{
		"page":  query.Page,
		"limit": query.Limit,
	})
}

type pString struct {
	Id int `uri:"id"`
}

func paramHandler(ctx *gin.Context) {
	// id := ctx.Param("id")
	var param pString
	if err := ctx.ShouldBindUri(&param); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(200, gin.H{
		"id": param.Id,
	})
}

type User struct {
	Username string `form:"username" json:"username" binding:"required" `
	Password string `form:"password" json:"password" binding:"required"`
}

func bodyHandler(ctx *gin.Context) {
	var body User

	err := ctx.ShouldBind(&body)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(200, body)
}
