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

// GetAll godoc
// @Summary Get all cats
// @Tags cats
// @Produce json
// @Success 200 {array} models.Cat
// @Failure 500 {object} map[string]string
// @Router /cats [get]
func (ch *CatHandler) GetAll(c *gin.Context) {
	cats, err := ch.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cats)
}

// GetByID godoc
// @Summary Get cat by ID
// @Tags cats
// @Produce json
// @Param id path int true "Cat ID"
// @Success 200 {object} models.Cat
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cats/{id} [get]
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

// Create godoc
// @Summary Create cat
// @Tags cats
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param cat body models.Cat true "Cat"
// @Success 201 {object} models.Cat
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cats [post]
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

// Update godoc
// @Summary Update cat
// @Tags cats
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Cat ID"
// @Param cat body models.Cat true "Cat"
// @Success 200 {object} models.Cat
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cats/{id} [put]
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

// Delete godoc
// @Summary Delete cat
// @Tags cats
// @Security ApiKeyAuth
// @Param id path int true "Cat ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cats/{id} [delete]
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

// DeleteAll godoc
// @Summary Delete all cats
// @Tags cats
// @Security ApiKeyAuth
// @Success 204
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cats [delete]
func (ch *CatHandler) DeleteAll(c *gin.Context) {
	if err := ch.repo.DeleteAll(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
