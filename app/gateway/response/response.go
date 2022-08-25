package response

type Pagination struct {
	Page  int32 `json:"page"`
	List  any   `json:"list"`
	Total int64 `json:"total"`
}
