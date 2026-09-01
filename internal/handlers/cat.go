package handlers

import (
	"cp_lab1/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CatHandler struct {
	repo repositories.CatRepo
}

func NewCatHandler(repo repositories.CatRepo) *CatHandler {
	return &CatHandler{
		repo: repo,
	}
}

func (ch *CatHandler) GetAll(c *gin.Context) {
	cats := ch.repo.GetAll()

	c.JSON(http.StatusOK, cats)
}
