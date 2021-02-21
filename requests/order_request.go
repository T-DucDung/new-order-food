package requests

type RequestOrder struct {
	lastUpdate int64
	Number string
	Address string
	Name string
	Detail  []RequestOrderDetail
}

type RequestOrderDetail struct {
	ProductId int
	Quantity int
}

