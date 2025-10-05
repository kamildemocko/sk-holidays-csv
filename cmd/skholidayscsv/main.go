package main

import (
	"fmt"
	"os"
	"strings"

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

	fmt.Println("getting calendar")
	h, err := calendarific.GetCurrentHolidays(
		apiKey,
		strings.ToUpper(*country),
		*year,
	)
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
