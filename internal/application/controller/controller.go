package controller

import (
	"blog-go-api/internal/application/dto"
	"blog-go-api/internal/application/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	repository repository.Repository
}

func New(repository repository.Repository) *Controller {
	return &Controller{repository: repository}
}

func (c *Controller) Create(ctx *gin.Context) {

	var input dto.PostDTO

	err := ctx.BindJSON(input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	res, err = c.repository.Create(&input)
	if err != nil {
		return
	}

}
