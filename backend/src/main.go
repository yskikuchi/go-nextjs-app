package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yskikuchi/go-nextjs-app/handler"
	"github.com/yskikuchi/go-nextjs-app/infra"
	"github.com/yskikuchi/go-nextjs-app/model"
	"github.com/yskikuchi/go-nextjs-app/repository"
	"github.com/yskikuchi/go-nextjs-app/routes"
)

func main() {
	db, err := infra.GetConn()
	if err != nil {
		log.Fatal(err)
	}

	bookingRepo := repository.NewBookingRepository()
	bookingHandler := handler.NewBookingHandler(bookingRepo)

	r := gin.Default()

	routes.BookingRoutes(r, bookingHandler)

	r.GET("/cars", func(c *gin.Context) {
		var cars []model.Car
		db.Find(&cars)
		c.JSON(http.StatusOK, cars)
	})

	r.POST("/cars", func(c *gin.Context) {
		var car model.Car
		if err := c.ShouldBindJSON(&car); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&car)
		c.JSON(http.StatusCreated, car)
	})

	r.Run()
}
