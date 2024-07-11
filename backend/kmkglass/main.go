package main

import (
	"kmkglass/database"
	"kmkglass/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	log.Println("MySQL init")
	database.InitRedis()
	log.Println("Redis init")
	database.InitMinio()
	log.Println("Minio init")

	r := gin.Default()

	r.GET("/products", handlers.GetProducts)
	r.GET("/years", handlers.GetYearsModel)
	r.GET("/brands", handlers.GetBrands)
	r.GET("/models", handlers.GetModelsBrand)
	r.GET("/glassoptions", handlers.GetGlassOptionsGlasType)
	r.GET("/glasstypes", handlers.GetGlassTypes)
	r.GET("/filterproducts", handlers.GetFilterProducts)

	r.POST("/products", handlers.CreateProduct)

	r.Run(":8080")
}
