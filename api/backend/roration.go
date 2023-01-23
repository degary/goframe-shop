package backend

import "github.com/gogf/gf/v2/frame/g"

type RotationAddReq struct {
	g.Meta `path:"/backend/rotation/add" method:"post" tags:"图片链接" summary:"创建图片链接接口"`
	PicUrl string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort"     dc:"排序"`
}
type RotationAddRes struct {
	//todo
	RotationId int `json:"rotationId"`
}
