package dto

type Pagination struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"pageSize"`
}

type Authority struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}
