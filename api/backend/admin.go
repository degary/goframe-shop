package backend

import "github.com/gogf/gf/v2/frame/g"

type AdminAddReq struct {
	g.Meta   `path:"/backend/admin/add" method:"post" tags:"管理员" summary:"创建管理员链接接口"`
	Name     string `json:"name"    v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password"    v:"required#密码不能为空" dc:"密码"`
	RoleIds  string `json:"role_ids"    v:"required#角色id不能为空" dc:"角色id"`
	IsAdmin  int    `json:"is_admin"     dc:"是否为超管[1:是超管，0:不是超管]"`
}
type AdminAddRes struct {
	AdminId int `json:"admin_Id"`
}

type AdminDeleteReq struct {
	g.Meta `path:"/backend/admin/delete" method:"delete" tags:"管理员" summary:"删除管理员接口"`
	Id     int `v:"min:1#请选择需要删除的管理员" dc:"管理员id"`
}
type AdminDeleteRes struct{}

type AdminUpdateReq struct {
	g.Meta   `path:"/backend/admin/update/{Id}" method:"post" tags:"管理员" summary:"修改管理员接口"`
	Id       int    `json:"id"      v:"min:1#请选择需要修改的管理员" dc:"管理员Id"`
	Name     string `json:"name"     dc:"用户名"`
	Password string `json:"password"     dc:"密码"`
	RoleIds  string `json:"role_ids"     dc:"角色id"`
	IsAdmin  int    `json:"is_admin"     dc:"是否为超管[1:是超管，0:不是超管]"`
}
type AdminUpdateRes struct{}

type AdminGetListCommonReq struct {
	g.Meta `path:"/backend/admin/list" method:"get" tags:"管理员" summary:"列出管理员接口"`
	CommonPaginationReq
}
type AdminGetListCommonRes struct {
	List  interface{} `json:"list"`  // 列表
	Page  int         `json:"page"`  // 分页码
	Size  int         `json:"size"`  // 分页数量
	Total int         `json:"total"` // 数据总数
}
