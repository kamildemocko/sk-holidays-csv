package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kamildemocko/sk-holidays-csv/internal/calendarific"
	"github.com/kamildemocko/sk-holidays-csv/internal/tabularize"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	apiKey := os.Getenv("api_key")
	thisYear := time.Now().Format("2006")

	h, err := calendarific.GetCurrentHolidays(apiKey, thisYear)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	outPath := "output.csv"
	err = tabularize.SaveHolidaysToCSV(h, outPath)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
