package model

type GetManageDetailInput struct {
	Id      int    `json:"id" dc:"管理员ID"`
	Account string `json:"account" dc:"管理员账号"`
}

type GetManageDetailOutput struct {
	Id       int    `json:"id" dc:"管理员ID"`
	Account  string `json:"account" dc:"管理员账号"`
	Password string `json:"password" dc:"管理员密码"`
	Status   int    `json:"status" dc:"管理员状态 0-初始化 1-正常 2-冻结"`
	IsSuper  int    `json:"is_super" dc:"是否超管 1-是"`
}

type MenuAddInput struct {
	Pid           int
	Type          int
	Path          string
	Sort          int
	Component     string
	Title         string
	Icon          string
	Hidden        int
	KeepAlive     int
	ActiveMenu    string
	AlwaysShow    int
	IsLargeScreen int
	IsFirstLevel  int
	IsSecondLevel int
	NoRedirect    int
	IsLink        int
	Remark        string
	Status        int
}

type MenuAddOutput struct {
	Id int64 `json:"id" dc:"菜单ID"`
}
