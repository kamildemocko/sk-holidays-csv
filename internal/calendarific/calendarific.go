package calendarific

type Response struct {
	Meta struct {
		Code int `json:"code"`
	} `json:"meta"`
	Response struct {
		Holidays Holiday `json:"response"`
	}
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
