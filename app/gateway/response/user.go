package response

type Role struct {
	ID        uint32 `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	ParentID  uint32 `json:"parentId"`
	IsDefault int32  `json:"isDefault"`
	Children  []Role `json:"children"`
}

type Login struct {
	UserID    uint32 `json:"userID"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

type User struct {
	ID        uint32 `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Username  string `json:"username"`
	Gender    int32  `json:"gender"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	Roles     []Role `json:"roles,omitempty"`
}
