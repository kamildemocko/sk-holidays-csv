# SK Holidays to CSV

A simple Go application that fetches Slovak holidays from the Calendarific API and saves them as CSV.

## Features
- Retrieves holidays for given year and country
- Saves holidays to a CSV file with configurable delimiter
- Includes holiday name, description, date, type, and URL in the output

## Usage

```bash
go run ./cmd/skholidayscsv -o /path/to/output.csv [-d ";"] [-c "SK"] [-y "2024"]
```

### Arguments
- `-o` - Required. Full path to output CSV file
- `-d` - Optional. Delimiter character (default: comma)
- `-c` - Optional. Country code (2 characters, default: "SK")
- `-y` - Optional. Year (default: current year)


## Requirements
- Calendarific API key

## .env file

```
api_key=secret
```
