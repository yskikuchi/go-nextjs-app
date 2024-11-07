package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yskikuchi/go-nextjs-app/handler"
	"github.com/yskikuchi/go-nextjs-app/middleware"
)

func BookingRoutes(router *gin.Engine, bookingHandler *handler.BookingHandler) {
	router.GET("/bookings", middleware.Authenticate, bookingHandler.FindAll)
	router.GET("/bookings/summaries", bookingHandler.FindAllSummaries)
	router.POST("/bookings", bookingHandler.Create)
	router.POST("/bookings/:id/approve", middleware.Authenticate, bookingHandler.Approve)
	router.POST("/bookings/:id/cancel", middleware.Authenticate, bookingHandler.Cancel)
	router.POST("/bookings/search", bookingHandler.Search)
}

func CarRoutes(router *gin.Engine, carHandler *handler.CarHandler) {
	router.GET("/cars", carHandler.FindAll)
	router.POST("/cars", middleware.Authenticate, carHandler.Create)
	router.PATCH("/cars/:id", middleware.Authenticate, carHandler.Update)
}

func AdminRoutes(router *gin.Engine, adminHandler *handler.AdminHandler) {
	router.POST("/admins", adminHandler.Create)
	router.POST("/admins/login", adminHandler.Login)
}
