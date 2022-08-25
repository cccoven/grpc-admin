package request

type Pagination struct {
	Page     int32 `json:"page" form:"page"`
	PageSize int32 `json:"pageSize" form:"page_size"`
}

type MultipleIDs struct {
	IDs []uint32 `json:"ids" validate:"required"`
}
