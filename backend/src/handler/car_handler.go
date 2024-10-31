package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yskikuchi/go-nextjs-app/model"
	"github.com/yskikuchi/go-nextjs-app/repository"
)

type CarHandler struct {
	Repo *repository.CarRepository
}

func NewCarHandler(repo *repository.CarRepository) *CarHandler {
	return &CarHandler{
		Repo: repo,
	}
}

func (h *CarHandler) FindAll(c *gin.Context) {
	cars, err := h.Repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cars)
}

func (h *CarHandler) Create(c *gin.Context) {
	var car model.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car, err := h.Repo.Create(&car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created"})
}
