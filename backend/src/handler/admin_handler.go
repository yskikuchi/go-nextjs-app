package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yskikuchi/go-nextjs-app/model"
	"github.com/yskikuchi/go-nextjs-app/repository"
	"github.com/yskikuchi/go-nextjs-app/service"
)

type AdminHandler struct {
	Repo *repository.AdminRepository
}

func NewAdminHandler(repo *repository.AdminRepository) *AdminHandler {
	return &AdminHandler{
		Repo: repo,
	}
}

func (h *AdminHandler) Create(c *gin.Context) {
	var admin model.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err := h.Repo.Create(&admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created"})
}

func (h *AdminHandler) Login(c *gin.Context) {
	var admin model.Admin
	var existing_admin model.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing_admin, err := h.Repo.FindByEmail(admin.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "メールアドレスまたはパスワードが正しくありません"})
		return
	}

	if service.VerifyPassword(existing_admin.Password, admin.Password) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "メールアドレスまたはパスワードが正しくありません"})
		return
	}

	token, err := service.GenerateToken(existing_admin.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "ログインしました",
		"access_token": token,
	})
}
