package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "net/http/pprof"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var redisClient *redis.Client

type Products struct {
	idproducts int
	/*Price         int
	Name          string
	Year          string
	Article       string
	Brand         string
	Model         string
	Length        int
	Photo         string
	Width         int
	InStock       bool
	Amount        int
	IdTypeProduct int
	IdOptions     int*/
}

func initDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	//dsn := "root:123437@tcp(127.0.0.1:3306)/kmkglass"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}
}

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	limit := 15
	query := r.URL.Query().Get("q")
	offset := r.URL.Query().Get("o")
	if query == "" {
		http.Error(w, "Query parameter 'q' is missing", http.StatusBadRequest)
		return
	}
	if offset == "" {
		http.Error(w, "Query parameter 'o' is missing", http.StatusBadRequest)
		return
	}
	//println(query)

	ctx := context.Background()
	cachedResult, err := redisClient.Get(ctx, query+string(limit)+string(offset)).Result()
	if err == redis.Nil {
		sqlQuery := `SELECT p.idproducts
					FROM products p
					JOIN models m ON p.models_id = m.idmodels
					JOIN year_model ym ON p.year_model_id = ym.idyear_model
					WHERE m.name = 'Golf' AND ym.name = '2023' LIMIT 15 OFFSET 1;`
		/*``*/
		//"SELECT * FROM products WHERE brand LIKE ? LIMIT ? OFFSET ?"
		/*`SELECT *
		FROM products
		JOIN typeproduct ON products.idTypeProduct = typeproduct.idTypeProduct
		WHERE typeproduct.name LIKE ? LIMIT ? OFFSET ?`*/
		//
		query = "Golf"
		rows, err := db.Query(sqlQuery /*, "%"+query+"%" , limit, offset*/)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка запроса к MySQL: %v", err), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var results []Products
		for rows.Next() {
			var row Products
			if err := rows.Scan(&row.idproducts); err != nil {
				http.Error(w, fmt.Sprintf("Ошибка сканирования полученных из MySQL данных: %v", err), http.StatusInternalServerError)
				return
			}
			results = append(results, row)
		}
		fmt.Println(results)

		if err := rows.Err(); err != nil {
			http.Error(w, fmt.Sprintf("Ошибка итерации строки: %v", err), http.StatusInternalServerError)
			return
		}

		// Сериализация данных в JSON
		jsonData, err := json.Marshal(results)
		if err != nil {
			log.Fatalf("Ошибка сериализации: %v", err)
		}

		// Установка заголовков ответа и отправка
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
		log.Println("MySQL data")

		//Установка данных в redis
		errRedis := redisClient.Set(ctx, query+string(limit)+string(offset), jsonData, 0).Err()
		if errRedis != nil {
			log.Fatalf("Ошибка записи в Redis: %v", errRedis)
		}

	} else if err != nil {
		http.Error(w, fmt.Sprintf("Redis error: %v", err), http.StatusInternalServerError)
	} else {
		//Отправка данных из Redis
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, cachedResult)
		log.Println("Redis data")
	}

}

func main() {
	initDB()
	initRedis()

	// Запуск сервера pprof для профилирования
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	http.HandleFunc("/search", searchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
