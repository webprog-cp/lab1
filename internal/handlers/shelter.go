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

func (sh *ShelterHandler) GetAll(c *gin.Context) {
	shelters, err := sh.shelterRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shelters)
}

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

func (sh *ShelterHandler) DeleteAll(c *gin.Context) {
	if err := sh.shelterRepo.DeleteAll(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

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
