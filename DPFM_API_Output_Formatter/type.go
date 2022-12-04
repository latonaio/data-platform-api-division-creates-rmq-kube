package dpfm_api_output_formatter

type Division struct {
	Division     string       `json:"Division"`
	DivisionText DivisionText `json:"DivisionText"`
}

type DivisionText struct {
	Division     string `json:"Division"`
	Language     string `json:"Language"`
	DivisionName string `json:"DivisionName"`
}
