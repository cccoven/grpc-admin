package request

type RouteGroup struct {
	ID   uint32 `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

type ModifyRouteGroup struct {
	RouteGroup
	Name string `json:"name" form:"createdAt" validate:"required,max=50"`
}

type ListRouteGroup struct {
	Pagination
	RouteGroup
}

type Route struct {
	ID      uint32 `json:"id" form:"id"`
	Path    string `json:"path" form:"path"`
	Method  string `json:"method" form:"method"`
	Desc    string `json:"desc" form:"desc"`
	GroupID uint32 `json:"groupID" form:"group_id"`
}

type ModifyRoute struct {
	Route
	Path    string `json:"path" validate:"required,max=50"`
	Method  string `json:"method" validate:"required,max=10"`
	Desc    string `json:"desc" validate:"required"`
	GroupID uint32 `json:"groupID" validate:"required"`
}

type ListRoute struct {
	Pagination
	Route
}

type Menu struct {
	ID        uint32 `json:"id"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Sort      int32  `json:"sort"`
	ParentID  uint32 `json:"parentID"`
}

type ModifyMenu struct {
	Menu
	Path      string `json:"path" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Component string `json:"component" validate:"required"`
}

type ListMenu struct {
	Pagination
	Menu
}
