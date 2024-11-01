package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yskikuchi/go-nextjs-app/handler"
)

func BookingRoutes(router *gin.Engine, bookingHandler *handler.BookingHandler) {
	router.GET("/bookings", bookingHandler.FindAll)
	router.POST("/bookings", bookingHandler.Create)
	router.POST("/bookings/:id/approve", bookingHandler.Approve)
}

func CarRoutes(router *gin.Engine, carHandler *handler.CarHandler) {
	router.GET("/cars", carHandler.FindAll)
	router.POST("/cars", carHandler.Create)
}
