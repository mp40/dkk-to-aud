package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
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

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	values := read.Read()

	data := calculate.GetResults(values)

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

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", viewHandler)
	err := http.ListenAndServe(":"+port, nil)
	log.Panic(err)
}
