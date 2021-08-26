package read

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"github.com/joho/godotenv"
)

func getEnvVariable(key string) string {
	env := os.Getenv("APP_ENV")
	if env != "production" {
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	return os.Getenv(key)
}

func Read() [][]interface{} {
	ctx := context.Background()

	secrets := getEnvVariable("JWT_CONFIG")

	config, err := google.JWTConfigFromJSON([]byte(secrets), "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := config.Client(ctx)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := getEnvVariable("SHEET_ID")
	readRange := getEnvVariable("READ_RANGE")

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	return resp.Values
}
