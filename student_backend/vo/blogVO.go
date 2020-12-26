package vo

type BlogListRequest struct {
	CourseID int `json:"course_id"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
