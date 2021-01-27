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

type Err struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func NewErr(err error) *Err {
	return &Err{
		Code:    MapErrorCode[err],
		Message: MapDescription[err],
	}
}
