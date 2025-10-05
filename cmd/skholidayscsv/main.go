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
	parseArgs()

	apiKey := os.Getenv("api_key")
	thisYear := time.Now().Format("2006")

	fmt.Println("getting calendar")
	h, err := calendarific.GetCurrentHolidays(apiKey, thisYear)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("saving CSV file")
	err = tabularize.SaveHolidaysToCSV(h, *output, []rune(*delimiter)[0])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("> done")
}
