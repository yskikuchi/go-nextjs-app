package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yskikuchi/go-nextjs-app/handler"
	"github.com/yskikuchi/go-nextjs-app/repository"
	"github.com/yskikuchi/go-nextjs-app/routes"
)

func main() {
	bookingRepo := repository.NewBookingRepository()
	bookingHandler := handler.NewBookingHandler(bookingRepo)

	carRepo := repository.NewCarRepository()
	carHandler := handler.NewCarHandler(carRepo)

	adminRepo := repository.NewAdminRepository()
	adminHandler := handler.NewAdminHandler(adminRepo)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{"http://localhost:3000"},

		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{"*"},

		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{"*"},
	}))

	routes.BookingRoutes(r, bookingHandler)
	routes.CarRoutes(r, carHandler)
	routes.AdminRoutes(r, adminHandler)

	r.Run()
}
