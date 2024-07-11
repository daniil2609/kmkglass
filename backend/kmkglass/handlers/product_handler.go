package handlers

import (
	"context"
	"encoding/json"
	"kmkglass/database"
	"kmkglass/models"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
)

func GetProducts(c *gin.Context) {
	// Получаем параметры пагинации
	lastID := c.DefaultQuery("lastId", "0")
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	cacheKey := "products_" + lastID + "_" + strconv.Itoa(pageSize)

	// Проверяем, есть ли кэшированные данные в Redis
	cachedProducts, err := database.RedisClient.Get(database.Ctx, cacheKey).Result()
	if err == redis.Nil {
		// Если данных нет, запрашиваем их из базы данных
		rows, err := database.DB.Query("SELECT * FROM products p WHERE p.idproducts > ? ORDER BY p.idproducts LIMIT ?;", lastID, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var products []models.Product
		for rows.Next() {
			var product models.Product
			if err := rows.Scan(&product.Idproducts,
				&product.Price,
				&product.Name,
				&product.Article,
				&product.Length,
				&product.Photo,
				&product.Width,
				&product.Amount,
				&product.Brands_name,
				&product.Models_name,
				&product.Year_model_name,
				&product.Glass_types_name,
				&product.Glass_options_name); err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			// Получение ссылки на фотографию
			reqParams := make(url.Values)
			presignedURL, err := database.MinioClient.PresignedGetObject(context.Background(), database.BucketName, product.Photo, time.Hour, reqParams)
			if err != nil {
				log.Println("Error generating presigned URL:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			product.Photo = presignedURL.String()
			products = append(products, product)
		}

		// Кэшируем данные в Redis на 10 минут
		productsJSON, _ := json.Marshal(products)
		database.RedisClient.Set(database.Ctx, cacheKey, productsJSON, 10*time.Minute).Err()

		c.JSON(http.StatusOK, products)
		log.Println("MySQL Data")
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		// Если данные есть в кэше, используем их
		var products []models.Product
		json.Unmarshal([]byte(cachedProducts), &products)
		c.JSON(http.StatusOK, products)
		log.Println("Redis Data")
	}
}

func GetYearsModel(c *gin.Context) {
	// Получаем параметры пагинации
	car_model := c.DefaultQuery("model", "0")
	cacheKey := "years_" + car_model

	// Проверяем, есть ли кэшированные данные в Redis
	cachedYears, err := database.RedisClient.Get(database.Ctx, cacheKey).Result()
	if err == redis.Nil {
		// Если данных нет, запрашиваем их из базы данных
		rows, err := database.DB.Query(`SELECT *
										FROM kmkglass.year_model
										WHERE model_name = ?;`, car_model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var years []models.Year_model
		for rows.Next() {
			var year models.Year_model
			if err := rows.Scan(&year.Idyear_model,
				&year.Name,
				&year.Model_name); err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			years = append(years, year)
		}

		// Кэшируем данные в Redis на 10 минут
		productsJSON, _ := json.Marshal(years)
		database.RedisClient.Set(database.Ctx, cacheKey, productsJSON, 10*time.Minute).Err()

		c.JSON(http.StatusOK, years)
		log.Println("MySQL Data")
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		// Если данные есть в кэше, используем их
		var years []models.Year_model
		json.Unmarshal([]byte(cachedYears), &years)
		c.JSON(http.StatusOK, years)
		log.Println("Redis Data")
	}
}

func GetBrands(c *gin.Context) {
	// Получаем параметры пагинации
	cacheKey := "brands"
	// Проверяем, есть ли кэшированные данные в Redis
	cachedBrands, err := database.RedisClient.Get(database.Ctx, cacheKey).Result()
	if err == redis.Nil {
		// Если данных нет, запрашиваем их из базы данных
		rows, err := database.DB.Query(`SELECT *
										FROM kmkglass.brands b;`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var brands []models.Brands
		for rows.Next() {
			var brand models.Brands
			if err := rows.Scan(&brand.Idbrands, &brand.Name); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			brands = append(brands, brand)
		}

		// Кэшируем данные в Redis на 10 минут
		productsJSON, _ := json.Marshal(brands)
		database.RedisClient.Set(database.Ctx, cacheKey, productsJSON, 10*time.Minute).Err()

		c.JSON(http.StatusOK, brands)
		log.Println("MySQL Data")
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		// Если данные есть в кэше, используем их
		var brands []models.Year_model
		json.Unmarshal([]byte(cachedBrands), &brands)
		c.JSON(http.StatusOK, brands)
		log.Println("Redis Data")
	}
}

func GetModelsBrand(c *gin.Context) {
	// Получаем параметры пагинации
	car_brand := c.DefaultQuery("brand", "0")
	cacheKey := "models_" + car_brand

	// Проверяем, есть ли кэшированные данные в Redis
	cachedModels_car, err := database.RedisClient.Get(database.Ctx, cacheKey).Result()
	if err == redis.Nil {
		// Если данных нет, запрашиваем их из базы данных
		rows, err := database.DB.Query(`SELECT *
										FROM kmkglass.models
										WHERE brand_name = ?;`, car_brand)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var models_car []models.Models
		for rows.Next() {
			var model_car models.Models
			if err := rows.Scan(&model_car.Idmodels,
				&model_car.Name,
				&model_car.Brand_name); err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			models_car = append(models_car, model_car)
		}

		// Кэшируем данные в Redis на 10 минут
		productsJSON, _ := json.Marshal(models_car)
		database.RedisClient.Set(database.Ctx, cacheKey, productsJSON, 10*time.Minute).Err()

		c.JSON(http.StatusOK, models_car)
		log.Println("MySQL Data")
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		// Если данные есть в кэше, используем их
		var models_car []models.Year_model
		json.Unmarshal([]byte(cachedModels_car), &models_car)
		c.JSON(http.StatusOK, models_car)
		log.Println("Redis Data")
	}
}

func GetGlassOptionsGlasType(c *gin.Context) {
	// Получаем параметры пагинации
	car_glasstype := c.DefaultQuery("glasstype", "0")
	cacheKey := "glassoptions_" + car_glasstype

	// Проверяем, есть ли кэшированные данные в Redis
	cachedCar_glassoptions, err := database.RedisClient.Get(database.Ctx, cacheKey).Result()
	if err == redis.Nil {
		// Если данных нет, запрашиваем их из базы данных
		rows, err := database.DB.Query(`SELECT *
										FROM kmkglass.glass_options
										WHERE glass_type_name = ?;`, car_glasstype)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var glass_options []models.Glass_options
		for rows.Next() {
			var glass_option models.Glass_options
			if err := rows.Scan(&glass_option.Idglass_options,
				&glass_option.Name,
				&glass_option.Glass_type_name); err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			glass_options = append(glass_options, glass_option)
		}

		// Кэшируем данные в Redis на 10 минут
		productsJSON, _ := json.Marshal(glass_options)
		database.RedisClient.Set(database.Ctx, cacheKey, productsJSON, 10*time.Minute).Err()

		c.JSON(http.StatusOK, glass_options)
		log.Println("MySQL Data")
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		// Если данные есть в кэше, используем их
		var glass_options []models.Year_model
		json.Unmarshal([]byte(cachedCar_glassoptions), &glass_options)
		c.JSON(http.StatusOK, glass_options)
		log.Println("Redis Data")
	}
}

func GetGlassTypes(c *gin.Context) {
	// Получаем параметры пагинации
	cacheKey := "glasstypes"
	// Проверяем, есть ли кэшированные данные в Redis
	cachedGlasstypes, err := database.RedisClient.Get(database.Ctx, cacheKey).Result()
	if err == redis.Nil {
		// Если данных нет, запрашиваем их из базы данных
		rows, err := database.DB.Query(`SELECT *
										FROM kmkglass.glass_types;`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var glasstypes []models.Glass_types
		for rows.Next() {
			var glasstype models.Glass_types
			if err := rows.Scan(&glasstype.Idglass_types, &glasstype.Name); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			glasstypes = append(glasstypes, glasstype)
		}

		// Кэшируем данные в Redis на 10 минут
		productsJSON, _ := json.Marshal(glasstypes)
		database.RedisClient.Set(database.Ctx, cacheKey, productsJSON, 10*time.Minute).Err()

		c.JSON(http.StatusOK, glasstypes)
		log.Println("MySQL Data")
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		// Если данные есть в кэше, используем их
		var glasstypes []models.Year_model
		json.Unmarshal([]byte(cachedGlasstypes), &glasstypes)
		c.JSON(http.StatusOK, glasstypes)
		log.Println("Redis Data")
	}
}

// brands_name - обязательный, хотя бы один
func GetFilterProducts(c *gin.Context) {
	// Получаем параметры пагинации
	lastID := c.DefaultQuery("lastId", "0")
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	brandName := c.DefaultQuery("brandName", "")
	modelName := c.DefaultQuery("modelName", "")
	yearModelName := c.DefaultQuery("yearModelName", "")
	glassTypeName := c.DefaultQuery("glassTypeName", "")
	glassOptionName := c.DefaultQuery("glassOptionName", "")

	cacheKey := "products_" + brandName + "_" + modelName + "_" + yearModelName + "_" + glassTypeName + "_" + glassOptionName + "_" + lastID + "_" + strconv.Itoa(pageSize)
	// Проверяем, есть ли кэшированные данные в Redis
	cachedProducts, err := database.RedisClient.Get(database.Ctx, cacheKey).Result()
	if err == redis.Nil {
		// Если данных нет, запрашиваем их из базы данных
		// Сборка запроса
		query := `
			SELECT p.*
			FROM products p
			WHERE`
		var args []interface{}

		if brandName != "" {
			query += " p.brands_name = ?"
			args = append(args, brandName)
		}
		if modelName != "" {
			query += " AND p.models_name = ?"
			args = append(args, modelName)
		}
		if yearModelName != "" {
			query += " AND p.year_model_name = ?"
			args = append(args, yearModelName)
		}
		if glassTypeName != "" {
			query += " AND p.glass_types_name = ?"
			args = append(args, glassTypeName)
		}
		if glassOptionName != "" {
			query += " AND p.glass_options_name = ?"
			args = append(args, glassOptionName)
		}
		query += " AND p.idproducts > ? ORDER BY p.idproducts LIMIT ?;"
		args = append(args, lastID, pageSize)
		// Выполнение запроса
		rows, err := database.DB.Query(query, args...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var products []models.Product
		for rows.Next() {
			var product models.Product
			if err := rows.Scan(&product.Idproducts,
				&product.Price,
				&product.Name,
				&product.Article,
				&product.Length,
				&product.Photo,
				&product.Width,
				&product.Amount,
				&product.Brands_name,
				&product.Models_name,
				&product.Year_model_name,
				&product.Glass_types_name,
				&product.Glass_options_name); err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			// Получение ссылки на фотографию
			reqParams := make(url.Values)
			presignedURL, err := database.MinioClient.PresignedGetObject(context.Background(), database.BucketName, product.Photo, time.Hour, reqParams)
			if err != nil {
				log.Println("Error generating presigned URL:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			product.Photo = presignedURL.String()
			products = append(products, product)
		}

		// Кэшируем данные в Redis на 10 минут
		productsJSON, _ := json.Marshal(products)
		database.RedisClient.Set(database.Ctx, cacheKey, productsJSON, 10*time.Minute).Err()

		c.JSON(http.StatusOK, products)
		log.Println("MySQL Data")
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		// Если данные есть в кэше, используем их
		var products []models.Product
		json.Unmarshal([]byte(cachedProducts), &products)
		c.JSON(http.StatusOK, products)
		log.Println("Redis Data")
	}
}

func CreateProduct(c *gin.Context) {
	var input models.Product
	input.Price, _ = strconv.Atoi(c.PostForm("price"))
	input.Name = c.PostForm("name")
	input.Article = c.PostForm("article")
	input.Length, _ = strconv.Atoi(c.PostForm("length"))
	input.Width, _ = strconv.Atoi(c.PostForm("width"))
	input.Amount, _ = strconv.Atoi(c.PostForm("amount"))
	input.Brands_name = c.PostForm("brands_name")
	input.Models_name = c.PostForm("models_name")
	input.Year_model_name = c.PostForm("year_model_name")
	input.Glass_types_name = c.PostForm("glass_types_name")
	input.Glass_options_name = c.PostForm("glass_options_name")

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//получение файла
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}
	defer file.Close()

	// Загрузка фотографии в MinIO
	fileName := header.Filename
	_, err = database.MinioClient.PutObject(context.Background(), database.BucketName, fileName, file, header.Size, minio.PutObjectOptions{ContentType: "image/png"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	input.Photo = fileName

	// Сохранение данных в MySQL
	result, err := database.DB.Exec("INSERT INTO products (price, name, article, length, photo, width, amount, brands_name, models_name, year_model_name, glass_types_name, glass_options_name) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", input.Price,
		input.Name,
		input.Article,
		input.Length,
		input.Photo,
		input.Width,
		input.Amount,
		input.Brands_name,
		input.Models_name,
		input.Year_model_name,
		input.Glass_types_name,
		input.Glass_options_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Получение ссылки на фотографию
	reqParams := make(url.Values)
	presignedURL, err := database.MinioClient.PresignedGetObject(context.Background(), database.BucketName, fileName, time.Hour, reqParams)
	if err != nil {
		log.Println("Error generating presigned URL:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Вывод ответа
	input.Idproducts = int(id)
	input.Photo = presignedURL.String()
	c.JSON(http.StatusOK, input)
}
