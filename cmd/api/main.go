package main

// "gotron/config"
// "gotron/internal/domain/product"

import (
	"log"
	"net/http"

	"gotron/internal/middleware"

	"gotron/config"

	"github.com/gin-gonic/gin"
	cartdomain "github.com/mfaisal-Ash/gotron/internal/domain/cart"
	orderdomain "github.com/mfaisal-Ash/gotron/internal/domain/order"
	productdomain "github.com/mfaisal-Ash/gotron/internal/domain/product"
)

func main() {
	cfg := config.Load()
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(middleware.RequestLogger())
	router.Use(middleware.Recovery())
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"meta": gin.H{"code": http.StatusNotFound, "message": "Route not found"}})
	})
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"meta": gin.H{"code": http.StatusMethodNotAllowed, "message": "Method not allowed"}})
	})

	productRepo := productdomain.NewInMemoryRepository()
	productService := productdomain.NewService(productRepo)
	if err := productRepo.SeedProducts(); err != nil {
		log.Fatalf("Failed to seed products: %v", err)
	}
	productHandler := productdomain.NewHandler(productService)

	cartRepo := cartdomain.NewInMemoryRepository()
	cartService := cartdomain.NewService(cartRepo, productRepo)
	cartHandler := cartdomain.NewHandler(cartService)

	orderRepo := orderdomain.NewInMemoryRepository()
	orderService := orderdomain.NewService(orderRepo, cartRepo)
	orderHandler := orderdomain.NewHandler(orderService)

	router.GET("/health", func(c *gin.context) {
		c.JSON(http.StatusOK, gin.H{
			"meta": gin.H{
				"code":    http.StatusOK,
				"status":  "success",
				"message": "Service is healthy",
			},
			"data": gin.H{
				"app_name": cfg.AppName,
				"env":      cfg.AppEnv,
			},
		})
	})

	api := router.Group("/api/v1")
	productdomain.RegisterRoutes(api, productHandler)
	cartdomain.RegisterRoutes(api, cartHandler)
	orderdomain.RegisterRoutes(api, orderHandler)

	log.Printf("%s listening on :%s", cfg.AppName, cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
