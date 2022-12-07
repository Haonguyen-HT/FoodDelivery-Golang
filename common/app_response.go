package common

type successRes struct {
	Data   interface{} `json:"data" form:"data"`
	Paging interface{} `json:"paging,omitempty" form:"paging"`
	Filter interface{} `json:"filter,omitempty" form:"filter"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return NewSuccessResponse(data, nil, nil)
}