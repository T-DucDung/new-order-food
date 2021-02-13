package responses

import "errors"

var (
	Success        = errors.New("Success")
	UnSuccess      = errors.New("UnSuccess")
	MapDescription = map[error]string{
		Success:   "Success!",
		UnSuccess: "Unsuccess!",
	}
	MapErrorCode = map[error]int64{
		Success:   200,
		UnSuccess: 201,
	}
)

var UnAuthResponse = ResponseSingle{Data: "UnAuthenticated", Error: &Err{Code: 405, Message: "UnAuthenticated"}}

type Err struct {
	Code    int64  `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
}

func NewErr(err error) *Err {
	return &Err{
		Code:    MapErrorCode[err],
		Message: MapDescription[err],
	}
}
