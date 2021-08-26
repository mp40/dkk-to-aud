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

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
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
	http.HandleFunc("/", viewHandler)
	err := http.ListenAndServe(":"+getEnvVariable("PORT"), nil)
	log.Fatal(err)
}
