package requests

type RequestDis struct {
	Rank       int     `json:"rank" xml:"rank"`
	Rate       float32 `json:"rate" xml:"rate"`
	Accumulate int     `json:"accumulate" xml:"accumulate"`
}
