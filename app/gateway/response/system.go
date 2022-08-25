package response

type RouteGroup struct {
	ID        uint32 `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Name      string `json:"name"`
}

type Route struct {
	ID        uint32     `json:"id"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
	Path      string     `json:"path"`
	Method    string     `json:"method"`
	Desc      string     `json:"desc"`
	GroupID   uint32     `json:"groupID"`
	Group     RouteGroup `json:"group"`
}

type Menu struct {
	ID        uint32 `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Sort      int32  `json:"sort"`
	ParentID  uint32 `json:"parentID"`
	Children  []Menu `json:"children"`
}
