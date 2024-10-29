package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yskikuchi/go-nextjs-app/infra"
	"github.com/yskikuchi/go-nextjs-app/repository"
)

func main() {
	db, err := infra.GetConn()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/bookings", func(c *gin.Context) {
		var tasks []repository.Booking
		db.Find(&tasks)
		c.JSON(http.StatusOK, tasks)
	})

	r.POST("/bookings", func(c *gin.Context) {
		var booking repository.Booking
		if err := c.ShouldBindJSON(&booking); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&booking)
		c.JSON(http.StatusCreated, booking)
	})

	r.GET("/cars", func(c *gin.Context) {
		var cars []repository.Car
		db.Find(&cars)
		c.JSON(http.StatusOK, cars)
	})

	r.POST("/cars", func(c *gin.Context) {
		var car repository.Car
		if err := c.ShouldBindJSON(&car); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&car)
		c.JSON(http.StatusCreated, car)
	})

	r.Run()
}
