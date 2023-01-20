package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccesResponse(data, paging, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, Filter: filter}

}
