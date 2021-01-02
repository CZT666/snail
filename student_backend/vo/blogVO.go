package vo

type BlogListRequest struct {
	BlogID int `json:"id"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
