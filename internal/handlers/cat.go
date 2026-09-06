package handlers

import (
	"cp_lab1/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
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
	cats, err := ch.repo.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, cats)
}

func (ch *CatHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	cat, err := ch.repo.GetByID(id)
	if err != nil {
		// error here
	}

	c.JSON(http.StatusOK, cat)
}
