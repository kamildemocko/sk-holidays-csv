package calendarific

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Meta struct {
		Code int `json:"code"`
	} `json:"meta"`
	Response struct {
		Holidays []Holiday `json:"holidays"`
	} `json:"response"`
}

type Holiday struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        struct {
		Iso string `json:"iso"`
	} `json:"date"`
	PrimaryType  string `json:"primary_type"`
	CanonicalURL string `json:"canonical_url"`
}

func GetCurrentHolidays(api_key, year string) ([]Holiday, error) {
	response, err := http.Get(fmt.Sprintf("https://calendarific.com/api/v2/holidays?api_key=%s&country=SK&year=%s", api_key, year))
	if err != nil {
		return []Holiday{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []Holiday{}, fmt.Errorf("expected status code 200, is %d\n", response.StatusCode)
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return []Holiday{}, err
	}

	var data Response
	err = json.Unmarshal(bytes, &data)

	return data.Response.Holidays, nil
}
