package main

import (
	"kmkglass/database"
	"kmkglass/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	database.InitRedis()

	r := gin.Default()

	r.GET("/products", handlers.GetProducts)
	r.GET("/years", handlers.GetYearsModel)
	r.GET("/brands", handlers.GetBrands)
	r.GET("/models", handlers.GetModelsBrand)
	r.GET("/glassoptions", handlers.GetGlassOptionsGlasType)
	r.GET("/glasstypes", handlers.GetGlassTypes)
	r.GET("/filterproducts", handlers.GetFilterProducts)

	r.Run(":8080")
}
