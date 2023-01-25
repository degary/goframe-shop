package backend

import "github.com/gogf/gf/v2/frame/g"

type PositionAddReq struct {
	g.Meta    `path:"/backend/position/add" method:"post" tags:"手工位" summary:"创建手工位链接接口"`
	PicUrl    string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsId   uint   `json:"goods_id" v:"required#商品id不能为空" dc:"商品id"`
	Sort      int    `json:"sort"     dc:"排序"`
}
type PositionAddRes struct {
	PositionId int `json:"positionId"`
}

type PositionDeleteReq struct {
	g.Meta `path:"/backend/position/delete" method:"delete" tags:"手工位" summary:"删除手工位接口"`
	Id     int `v:"min:1#请选择需要删除的手工位" dc:"手工位id"`
}
type PositionDeleteRes struct{}

type PositionUpdateReq struct {
	g.Meta    `path:"/backend/position/update/{Id}" method:"post" tags:"手工位" summary:"修改手工位接口"`
	Id        int    `json:"id"      v:"min:1#请选择需要修改的手工位" dc:"手工位Id"`
	PicUrl    string `json:"pic_url" v:"required#手工位图片链接不能为空" dc:"手工位模型"`
	Link      string `json:"link"    v:"required#手工位跳转链接不能为空" dc:"手工位跳转链接"`
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsId   uint   `json:"goods_id" v:"required#商品id不能为空" dc:"商品id"`
	Sort      int    `json:"sort"    dc:"排序"`
}
type PositionUpdateRes struct{}

type PositionGetListCommonReq struct {
	g.Meta `path:"/backend/position/list" method:"get" tags:"手工位" summary:"列出手工位接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type PositionGetListCommonRes struct {
	List  interface{} `json:"list"`  // 列表
	Page  int         `json:"page"`  // 分页码
	Size  int         `json:"size"`  // 分页数量
	Total int         `json:"total"` // 数据总数
}
