package responses

type ResponseBool struct {
	Error *Err `json:"error" xml:"error"`
}

type ResponseSingle struct {
	Data  interface{} `json:"data" xml:"data"`
	Error *Err        `json:"error" xml:"error"`
}

type ResponseArray struct {
	Data       interface{} `json:"data" xml:"data"`
	TotalCount int         `json:"total_count" xml:"total_count"`
	Error      *Err        `json:"error" xml:"error"`
}
