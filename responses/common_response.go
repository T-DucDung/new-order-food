package responses

type ResponseBool struct {
	Error *Err
}

type ResponseSingle struct {
	Data  interface{}
	Error *Err
}

type ResponseArray struct {
	Data       interface{}
	TotalCount int
	Error      *Err
}
