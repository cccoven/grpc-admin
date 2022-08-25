package request

type User struct {
	ID       uint32 `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Gender   int32  `json:"gender" form:"gender"`
	Avatar   string `json:"avatar" form:"avatar"`
	Phone    string `json:"phone" form:"phone"`
}

type ModifyUser struct {
	User
	Username string   `json:"username" form:"username" validate:"required"`
	RoleIDs  []uint32 `json:"roles" form:"role_ids"`
}

type ListUser struct {
	Pagination
	User
}

type AdminLogin struct {
	Type     int    `json:"type" form:"type" validate:"required"`       // 登录方式 1 账号密码 2 手机号验证码
	Account  string `json:"account" form:"account" validate:"required"` // 手机号或用户名
	Password string `json:"password" form:"password"`
	Code     string `json:"code" form:"code"` // 验证码
}

type Role struct {
	ID        uint32 `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Desc      string `json:"desc" form:"desc"`
	ParentID  int    `json:"parentID" form:"parent_id"`
	IsDefault int    `json:"isDefault" form:"is_default"`
}

type ModifyRole struct {
	Role
	Name string `json:"name" form:"name" validate:"required,max=50"`
	Desc string `json:"desc" form:"desc" validate:"required,max=200"`
}

type ListRole struct {
	Pagination
	Role
}

type Authority struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type AuthorizeRole struct {
	Authorities []Authority `json:"authorities" validate:"required"`
}

type SetRoleMenus struct {
	Menus []uint32 `json:"menus" validate:"required,gt=0,dive,required"`
}

type AssignRolesToUser struct {
	Roles []uint32 `json:"roles" validate:"required"`
}
