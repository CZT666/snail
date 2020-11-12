package vo

type PageRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type PageResponse struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

func NewPageRequest() (pageRequest *PageRequest) {
	pageRequest = new(PageRequest)
	pageRequest.Page = -1
	pageRequest.Page = -1
	return
}

func NewPageResponse(total int, data interface{}) (pageResponse *PageResponse) {
	pageResponse = new(PageResponse)
	pageResponse.Total = total
	pageResponse.Data = data
	return
}
