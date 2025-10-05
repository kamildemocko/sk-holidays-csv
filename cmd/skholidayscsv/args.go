package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var output = flag.String("o", "", "Output dir with filename for CSV, \nExample: -o C:\\programs\\cal\\output.csv")
var delimiter = flag.String("d", ",", "Used delimiter, default: ','")
var country = flag.String("c", "SK", "Country, default: 'SK'")
var year = flag.String("y", "", "Year, default: current year")

func parseArgs() {
	flag.Parse()

	validateArgs()
}

func validateArgs() {
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

	if len(*delimiter) != 1 {
		fmt.Println("Delimiter must be a single character")
		flag.Usage()
		os.Exit(1)
	}

	if len(*country) != 2 {
		fmt.Println("Country must be a code of 2 characters")
		flag.Usage()
		os.Exit(1)
	}

	if *year == "" {
		*year = time.Now().Format("2006")
	}
}
