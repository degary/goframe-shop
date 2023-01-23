package roration

import (
	"context"
	"github.com/degary/goframe-shop/internal/dao"
	"github.com/degary/goframe-shop/internal/model"
	"github.com/degary/goframe-shop/internal/service"
	"github.com/gogf/gf/v2/encoding/ghtml"
)

type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}

func New() *sRotation {
	return &sRotation{}
}

// Create 创建内容
func (s *sRotation) Add(ctx context.Context, in model.RotationAddInput) (out model.RotationAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationAddOutput{RotationId: int(lastInsertID)}, err
}
