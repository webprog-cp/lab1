package handlers

import (
	"net/http"

	"cp_lab1/internal/models"
	"cp_lab1/internal/repositories"

	"github.com/gin-gonic/gin"
)

type ShelterHandler struct {
	shelterRepo repositories.ShelterRepo
	catRepo     repositories.CatRepo
}

func NewShelterHandler(sr repositories.ShelterRepo, cr repositories.CatRepo) *ShelterHandler {
	return &ShelterHandler{
		shelterRepo: sr,
		catRepo:     cr,
	}
}

// GetAll godoc
// @Summary Get all shelters
// @Tags shelters
// @Produce json
// @Success 200 {array} models.Shelter
// @Failure 500 {object} models.ErrorResponse
// @Router /shelters [get]
func (sh *ShelterHandler) GetAll(c *gin.Context) {
	shelters, err := sh.shelterRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shelters)
}

// GetByID godoc
// @Summary Get shelter by ID
// @Tags shelters
// @Produce json
// @Param id path int true "Shelter ID"
// @Success 200 {object} models.Shelter
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /shelters/{id} [get]
func (sh *ShelterHandler) GetByID(c *gin.Context) {
	id, valid := idParam(c)
	if !valid {
		return
	}

	shelter, err := sh.shelterRepo.GetByID(id)
	if err != nil {
		c.JSON(statusByError(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shelter)
}

// Create godoc
// @Summary Create shelter
// @Tags shelters
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param shelter body models.Shelter true "Shelter"
// @Success 201 {object} models.Shelter
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /shelters [post]
func (sh *ShelterHandler) Create(c *gin.Context) {
	var shelter models.Shelter
	if err := c.BindJSON(&shelter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shelter JSON: " + err.Error()})
		return
	}

	shelter, err := sh.shelterRepo.Create(shelter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, shelter)
}

// Update godoc
// @Summary Update shelter
// @Tags shelters
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Shelter ID"
// @Param shelter body models.Shelter true "Shelter"
// @Success 200 {object} models.Shelter
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /shelters/{id} [put]
func (sh *ShelterHandler) Update(c *gin.Context) {
	id, valid := idParam(c)
	if !valid {
		return
	}

	var shelter models.Shelter
	if err := c.BindJSON(&shelter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shelter JSON: " + err.Error()})
		return
	}

	if err := sh.shelterRepo.Update(id, shelter); err != nil {
		c.JSON(statusByError(err), gin.H{"error": err.Error()})
		return
	}

	shelter.ID = id
	c.JSON(http.StatusOK, shelter)
}

// Delete godoc
// @Summary Delete shelter
// @Tags shelters
// @Security ApiKeyAuth
// @Param id path int true "Shelter ID"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /shelters/{id} [delete]
func (sh *ShelterHandler) Delete(c *gin.Context) {
	id, valid := idParam(c)
	if !valid {
		return
	}

	if err := sh.shelterRepo.Delete(id); err != nil {
		c.JSON(statusByError(err), gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteAll godoc
// @Summary Delete all shelters
// @Tags shelters
// @Security ApiKeyAuth
// @Success 204
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /shelters [delete]
func (sh *ShelterHandler) DeleteAll(c *gin.Context) {
	if err := sh.shelterRepo.DeleteAll(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetCats godoc
// @Summary Get shelter cats
// @Tags shelters
// @Produce json
// @Param id path int true "Shelter ID"
// @Success 200 {array} models.Cat
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /shelters/{id}/cats [get]
func (sh *ShelterHandler) GetCats(c *gin.Context) {
	id, valid := idParam(c)
	if !valid {
		return
	}

	if _, err := sh.shelterRepo.GetByID(id); err != nil {
		c.JSON(statusByError(err), gin.H{"error": err.Error()})
		return
	}

	cats, err := sh.catRepo.GetCatsByShelterID(id)
	if err != nil {
		c.JSON(statusByError(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cats)
}
