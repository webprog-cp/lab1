package handlers

import (
	"strconv"
	"net/http"

	"cp_lab1/internal/repositories"
	"cp_lab1/internal/models"

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

func (ch *CatHandler) GetByShelterID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	cats, err := ch.repo.GetCatsByShelterID(id)
	if err != nil {
		//err
	}

	c.JSON(http.StatusOK, cats)
}

func (ch *CatHandler) Create(c *gin.Context) {
	var cat models.Cat

	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	if err := ch.repo.Create(cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (ch *CatHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	var cat models.Cat
	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	if err := ch.repo.Update(id, cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	c.JSON(http.StatusOK, cat)
}

func (ch *CatHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	if err := ch.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (ch *CatHandler) DeleteAll(c *gin.Context) {
	if err := ch.repo.DeleteAll(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{})
}
