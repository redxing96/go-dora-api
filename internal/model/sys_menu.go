package model

// 获取所有菜单输入
type GetAllSysMenuInput struct {
	BaseInput
	Type int `json:"type" dc:"类型 1:菜单 2:接口 3:按钮"`
}

// 获取所有菜单输出
type GetAllSysMenuOutput struct {
}

type MenuAllResponse struct {
	List []*MenuItem `json:"list" dc:"列表"`
}

type MenuAddResponse struct {
	Id int64 `json:"id" dc:"菜单ID"`
}

type MenuItem struct {
	Id        int64        `json:"id" dc:"菜单ID"`
	Pid       int64        `json:"pid" dc:"父级ID"`
	Path      string       `json:"path" dc:"路径"`
	Name      string       `json:"name" dc:"名称"`
	Sort      int          `json:"sort" dc:"排序"`
	Component string       `json:"component" dc:"组件"`
	Meta      MenuItemMeta `json:"meta" dc:"元数据"`
	Status    int          `json:"status" dc:"状态 1:正常 2:禁用"`
	Children  []*MenuItem  `json:"children" dc:"子菜单"`
}

type MenuItemMeta struct {
	Title         string `json:"title" dc:"标题"`
	Icon          string `json:"icon" dc:"图标"`
	Hidden        bool   `json:"hidden" dc:"是否隐藏"`
	KeepAlive     bool   `json:"keep_alive" dc:"是否缓存"`
	ActiveMenu    string `json:"active_menu" dc:"激活菜单的path"`
	AlwaysShow    bool   `json:"always_show" dc:"是否总是显示为父菜单"`
	IsLargeScreen bool   `json:"is_large_screen" dc:"是否仅在大屏显示"`
	IsFirstLevel  bool   `json:"is_first_level" dc:"是否是一级导航"`
	IsSecondLevel bool   `json:"is_second_level" dc:"是否是二级导航"`
	NoRedirect    bool   `json:"no_redirect" dc:"是否禁止重定向"`
	IsLink        string `json:"is_link" dc:"是否是外部链接"`
}

type MenuDeleteInput struct {
	Ids []int `json:"ids" dc:"路由权限ID"`
}

type MenuDeleteOutput struct {
	Ids []int `json:"ids" dc:"路由权限ID"`
}

type MenuUpdateInput struct {
	Id            int    `json:"id" dc:"路由权限ID"`
	Pid           int    `json:"pid" dc:"父级ID"`
	Type          int    `json:"type" dc:"类型 1:菜单 2:接口"`
	Path          string `json:"path" dc:"路径"`
	Sort          int    `json:"sort" dc:"排序"`
	Component     string `json:"component" dc:"组件"`
	Title         string `json:"title" dc:"标题"`
	Icon          string `json:"icon" dc:"图标"`
	Hidden        int    `json:"hidden" dc:"是否隐藏 2:否 1:是"`
	KeepAlive     int    `json:"keep_alive" dc:"是否缓存 2:否 1:是"`
	ActiveMenu    string `json:"active_menu" dc:"激活菜单"`
	AlwaysShow    int    `json:"always_show" dc:"总是显示 2:否 1:是"`
	IsLargeScreen int    `json:"is_large_screen" dc:"是否大屏 2:否 1:是"`
	IsFirstLevel  int    `json:"is_first_level" dc:"是否一级菜单 2:否 1:是"`
	IsSecondLevel int    `json:"is_second_level" dc:"是否二级菜单 2:否 1:是"`
	NoRedirect    int    `json:"no_redirect" dc:"是否重定向 2:否 1:是"`
	IsLink        int    `json:"is_link" dc:"是否链接 2:否 1:是"`
	Remark        string `json:"remark" dc:"备注"`
	Status        int    `json:"status" dc:"状态 1-正常 2-禁用"`
}

type MenuUpdateOutput struct {
	Id int `json:"id" dc:"路由权限ID"`
}
