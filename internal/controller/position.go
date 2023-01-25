package controller

import (
	"context"

	"github.com/degary/goframe-shop/api/backend"
	"github.com/degary/goframe-shop/internal/model"
	"github.com/degary/goframe-shop/internal/service"
)

// Position 内容管理
var Position = cPosition{}

type cPosition struct{}

func (a *cPosition) Add(ctx context.Context, req *backend.PositionAddReq) (res *backend.PositionAddRes, err error) {
	out, err := service.Position().Add(ctx, model.PositionAddInput{
		PositionAddUpdateBase: model.PositionAddUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			Sort:      req.Sort,
			GoodsId:   req.GoodsId,
			GoodsName: req.GoodsName,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionAddRes{PositionId: out.PositionId}, nil
}

func (a *cPosition) Delete(ctx context.Context, req *backend.PositionDeleteReq) (res *backend.PositionDeleteRes, err error) {
	err = service.Position().Delete(ctx, req.Id)
	return
}

func (a *cPosition) Update(ctx context.Context, req *backend.PositionUpdateReq) (res *backend.PositionUpdateRes, err error) {
	err = service.Position().Update(ctx, model.PositionUpdateInput{
		Id: req.Id,
		PositionAddUpdateBase: model.PositionAddUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			Sort:      req.Sort,
			GoodsId:   req.GoodsId,
			GoodsName: req.GoodsName,
		},
	})
	return
}

// Index rotation list
func (a *cPosition) List(ctx context.Context, req *backend.PositionGetListCommonReq) (res *backend.PositionGetListCommonRes, err error) {
	getListRes, err := service.Position().GetList(ctx, model.PositionGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, err
}
