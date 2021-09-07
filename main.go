package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/joho/godotenv"
	"github.com/mp40/dkk-to-aud/cache"
	"github.com/mp40/dkk-to-aud/calculate"
	"github.com/mp40/dkk-to-aud/read"
)

func loadEnvFile() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getCachedData() (*calculate.Data, error) {
	cache := &cache.Cache
	var data *calculate.Data

	timestamp := time.Now()
	key := timestamp.Format("2006-01-02")

	cachedData, err := cache.Get(key)
	if err != nil {
		return data, err
	}

	var parseErr error
	parseErr = json.Unmarshal([]byte(cachedData), &data)
	if parseErr != nil {
		return data, err
	}

	return data, nil
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	data, cacheErr := getCachedData()

	if cacheErr != nil {
		values := read.Read()

		data = calculate.GetResults(values)
	}

	html, err := template.ParseFiles("view.html")
	check(err)

	err = html.Execute(writer, data)
	check(err)
}

func main() {
	env := os.Getenv("APP_ENV")
	if env != "production" {
		loadEnvFile()
	}

	cache.CreateRedisCache()

	port := os.Getenv("APP_PORT")

	if port == "" {
		log.Fatal("$APP_PORT must be set")
	}

	http.HandleFunc("/", viewHandler)
	err := http.ListenAndServe(":"+port, nil)
	log.Panic(err)
}
