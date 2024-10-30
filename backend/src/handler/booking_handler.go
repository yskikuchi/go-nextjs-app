package handler

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yskikuchi/go-nextjs-app/model"
	"github.com/yskikuchi/go-nextjs-app/repository"
)

type BookingHandler struct {
	Repo *repository.BookingRepository
}

func NewBookingHandler(repo *repository.BookingRepository) *BookingHandler {
	return &BookingHandler{
		Repo: repo,
	}
}

func (h *BookingHandler) FindAll(c *gin.Context) {
	bookings, err := h.Repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

func (h *BookingHandler) Create(c *gin.Context) {
	var booking model.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	referenceNumber, err := generateReferenceNumber(h)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	booking.ReferenceNumber = referenceNumber
	booking, err = h.Repo.Create(&booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"referenceNumber": booking.ReferenceNumber})
}

// 予約参照番号を生成する
func generateReferenceNumber(h *BookingHandler) (string, error) {
	const maxAttempts = 5
	const length = 8
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for attempts := 0; attempts < maxAttempts; attempts++ {
		refNum := make([]byte, length)
		for i := range refNum {
			num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
			if err != nil {
				return "", err
			}
			refNum[i] = charSet[num.Int64()]
		}

		_, err := h.Repo.FindByReferenceNumber(string(refNum))
		if err != nil {
			return string(refNum), nil
		}
	}

	return "", fmt.Errorf("failed to generate reference number")
}
