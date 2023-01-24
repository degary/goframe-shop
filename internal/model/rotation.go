package model

// RotationAddUpdateBase 创建/修改内容基类
type RotationAddUpdateBase struct {
	PicUrl string
	Link   string
	Sort   int
}

// RotationAddInput 创建内容
type RotationAddInput struct {
	RotationAddUpdateBase
}

// RotationAddOutput 创建内容返回结果
type RotationAddOutput struct {
	RotationId int `json:"rotation_id"`
}

type RotationUpdateInput struct {
	RotationAddUpdateBase
	Id int
}
