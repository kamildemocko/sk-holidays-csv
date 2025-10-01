package tabularize

import (
	"encoding/csv"
	"os"

	"github.com/kamildemocko/sk-holidays-csv/internal/calendarific"
)

func SaveHolidaysToCSV(holidays []calendarific.Holiday, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"name", "description", "date", "primary_type", "url"}
	err = writer.Write(header)
	if err != nil {
		return err
	}

	for _, holiday := range holidays {
		record := []string{
			holiday.Name,
			holiday.Description,
			holiday.Date.Iso,
			holiday.PrimaryType,
			holiday.CanonicalURL,
		}

		err = writer.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}
