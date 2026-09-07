package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"cp_lab1/internal/const_errors"
	"cp_lab1/internal/models"
	"cp_lab1/internal/repositories"

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

func idParam(c *gin.Context) (uint64, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be an unsigned integer"})
		return 0, false
	}
	return id, true
}

func statusByError(err error) int {
	if errors.Is(err, const_errors.NoEntityByID) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}

func (ch *CatHandler) GetAll(c *gin.Context) {
	cats, err := ch.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cats)
}

func (ch *CatHandler) GetByID(c *gin.Context) {
	id, valid := idParam(c)
	if !valid {
		return
	}

	cat, err := ch.repo.GetByID(id)
	if err != nil {
		c.JSON(statusByError(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cat)
}

func (ch *CatHandler) GetByShelterID(c *gin.Context) {
	id, valid := idParam(c)
	if !valid {
		return
	}

	cats, err := ch.repo.GetCatsByShelterID(id)
	if err != nil {
		c.JSON(statusByError(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cats)
}

func (ch *CatHandler) Create(c *gin.Context) {
	var cat models.Cat
	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cat JSON: " + err.Error()})
		return
	}

	cat, err := ch.repo.Create(cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cat)
}

func (ch *CatHandler) Update(c *gin.Context) {
	id, valid := idParam(c)
	if !valid {
		return
	}

	var cat models.Cat
	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cat JSON: " + err.Error()})
		return
	}

	if err := ch.repo.Update(id, cat); err != nil {
		c.JSON(statusByError(err), gin.H{"error": err.Error()})
		return
	}

	cat.ID = id
	c.JSON(http.StatusOK, cat)
}

func (ch *CatHandler) Delete(c *gin.Context) {
	id, valid := idParam(c)
	if !valid {
		return
	}

	if err := ch.repo.Delete(id); err != nil {
		c.JSON(statusByError(err), gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (ch *CatHandler) DeleteAll(c *gin.Context) {
	if err := ch.repo.DeleteAll(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
