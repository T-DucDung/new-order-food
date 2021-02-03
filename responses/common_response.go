package responses

type ResponseBool struct {
	Error *Err
}

type ResponseSingle struct {
	Data map[string]string
	Error *Err
}
