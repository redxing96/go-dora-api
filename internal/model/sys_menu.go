/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-25 10:26:14
 * @LastEditTime: 2025-07-02 18:47:19
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/model/sys_menu.go
 */
package model

// 获取所有菜单输入
type GetAllSysMenuInput struct {
	BaseInput
	Type   []int `json:"type" dc:"类型 1:菜单 2:接口 3:按钮 4:目录"`
	Status int   `json:"status" dc:"状态 1-正常 2-禁用"`
}

// 获取所有菜单输出
type GetAllSysMenuOutput struct {
}

type MenuAllResponse struct {
	List []*MenuItem `json:"list" dc:"列表"`
}

type MenuAddResponse struct {
	Id   int64  `json:"id" dc:"菜单ID"`
	Name string `json:"name" dc:"名称"`
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
	IconSvg       string `json:"icon_svg" dc:"图标svg"`
	Type          int    `json:"type" dc:"类型 1菜单,2接口,3按钮"`
	IsHidden      int    `json:"is_hidden" dc:"是否隐藏"`
	IsKeepAlive   int    `json:"is_keep_alive" dc:"是否缓存"`
	ActiveMenu    string `json:"active_menu" dc:"激活菜单的path"`
	IsLargeScreen int    `json:"is_large_screen" dc:"是否仅在大屏显示"`
	Link          string `json:"link" dc:"是否是外部链接"`
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
	Name          string `json:"name" dc:"名称"`
	Icon          string `json:"icon" dc:"图标"`
	IconSvg       string `json:"icon_svg" dc:"图标svg"`
	IsHidden      int    `json:"hidden" dc:"是否隐藏 2:否 1:是"`
	IsKeepAlive   int    `json:"keep_alive" dc:"是否缓存 2:否 1:是"`
	ActiveMenu    string `json:"active_menu" dc:"激活菜单"`
	IsLargeScreen int    `json:"is_large_screen" dc:"是否大屏 2:否 1:是"`
	Link          string `json:"is_link" dc:"是否链接 2:否 1:是"`
	Remark        string `json:"remark" dc:"备注"`
	Status        int    `json:"status" dc:"状态 1-正常 2-禁用"`
}

type MenuUpdateOutput struct {
	Id int `json:"id" dc:"路由权限ID"`
}

type MenuAddInput struct {
	Pid           int
	Type          int
	Name          string
	Path          string
	Sort          int
	Component     string
	Title         string
	Icon          string
	IconSvg       string
	IsHidden      int
	IsKeepAlive   int
	ActiveMenu    string
	IsLargeScreen int
	Link          string
	Remark        string
	Status        int
}

type MenuAddOutput struct {
	Id int64 `json:"id" dc:"菜单ID"`
}

type GetManageMenuInput struct {
	ManagerID int64
	MenuType  []int
	Status    int
}

type BtnList struct {
	Id        int64  `json:"id" dc:"菜单ID"`
	Pid       int64  `json:"pid" dc:"父级ID"`
	Name      string `json:"name" dc:"名称"`
	Path      string `json:"path" dc:"路径"`
	Component string `json:"component" dc:"组件"`
}
