package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kamildemocko/sk-holidays-csv/internal/calendarific"
	"github.com/kamildemocko/sk-holidays-csv/internal/tabularize"
)

var output = flag.String("o", "", "Output dir with filename for CSV, \nExample: -o C:\\programs\\cal\\output.csv")

func init() {
	_ = godotenv.Load(".env")
}

func parseArgs() {
	flag.Parse()

	if *output == "" {
		flag.Usage()
		os.Exit(1)
	}

	if !filepath.IsAbs(*output) {
		fmt.Println("Please provide full path")
		flag.Usage()
		os.Exit(1)
	}

	extension := filepath.Ext(*output)
	if extension == "" {
		fmt.Println("Path appears to be directory")
		fmt.Println("Please provide full path to a CSV")
		flag.Usage()
		os.Exit(1)
	}

	parentDir := filepath.Dir(*output)
	if _, err := os.Stat(parentDir); err != nil {
		fmt.Println("Input dir does not exist")
		flag.Usage()
		os.Exit(1)
	}
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
	err = tabularize.SaveHolidaysToCSV(h, *output)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("> done")
}
