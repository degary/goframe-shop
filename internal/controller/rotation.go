package controller

import (
	"context"

	"github.com/degary/goframe-shop/api/backend"
	"github.com/degary/goframe-shop/internal/model"
	"github.com/degary/goframe-shop/internal/service"
)

// Rotation 内容管理
var Rotation = cRotation{}

type cRotation struct{}

func (a *cRotation) Add(ctx context.Context, req *backend.RotationAddReq) (res *backend.RotationAddRes, err error) {
	out, err := service.Rotation().Add(ctx, model.RotationAddInput{
		RotationAddUpdateBase: model.RotationAddUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RotationAddRes{RotationId: out.RotationId}, nil
}
