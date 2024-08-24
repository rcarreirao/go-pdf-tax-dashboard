package cmd

import (
	"go_pdf_tax_dashboard/internal/window"
	"log"

	"github.com/joho/godotenv"
)

func Execute() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	window.Execute()
}
