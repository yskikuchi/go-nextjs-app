package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yskikuchi/go-nextjs-app/handler"
)

func BookingRoutes(router *gin.Engine, bookingHandler *handler.BookingHandler) {
	router.GET("/bookings", bookingHandler.FindAll)
	router.POST("/bookings", bookingHandler.Create)
}
