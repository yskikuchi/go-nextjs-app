package handler

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yskikuchi/go-nextjs-app/model"
	"github.com/yskikuchi/go-nextjs-app/repository"
	cv "github.com/yskikuchi/go-nextjs-app/validator"
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

	err := validateBooking(h, &booking)
	if err != nil {
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

func (h *BookingHandler) Approve(c *gin.Context) {
	id := c.Param("id")
	booking, err := h.Repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	booking.Status = "confirmed"
	booking, err = h.Repo.Update(booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Approved"})
}

func (h *BookingHandler) Cancel(c *gin.Context) {
	id := c.Param("id")
	booking, err := h.Repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	booking.Status = "canceled"
	booking, err = h.Repo.Update(booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Canceled"})
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

// TODO: 別ファイルへ切り出したい...
func validateBooking(h *BookingHandler, booking *model.Booking) error {
	validate := validator.New()
	validate.RegisterValidation("mustBeAfterNow", cv.MustBeAfterNow)

	if err := validate.Struct(booking); err != nil {
		return err
	}

	// 既存の予約と重複しないかを検証する
	_, err := h.Repo.FindByCarIDAndTimeRange(booking.CarID.String(), booking.StartTime.Format(time.RFC3339), booking.EndTime.Format(time.RFC3339))

	// 既存の予約が見つかった場合はエラーを返す
	if err == nil {
		return fmt.Errorf("選択した時間帯で既に予約が入っています")
	}

	return nil
}
