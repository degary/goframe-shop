package controller

import (
	"context"

	"github.com/degary/goframe-shop/api/backend"
	"github.com/degary/goframe-shop/internal/model"
	"github.com/degary/goframe-shop/internal/service"
)

// Admin 内容管理
var Admin = cAdmin{}

type cAdmin struct{}

func (a *cAdmin) Add(ctx context.Context, req *backend.AdminAddReq) (res *backend.AdminAddRes, err error) {
	out, err := service.Admin().Add(ctx, model.AdminAddInput{
		AdminAddUpdateBase: model.AdminAddUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			IsAdmin:  req.IsAdmin,
			RoleIds:  req.RoleIds,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.AdminAddRes{AdminId: out.AdminId}, nil
}

func (a *cAdmin) Delete(ctx context.Context, req *backend.AdminDeleteReq) (res *backend.AdminDeleteRes, err error) {
	err = service.Admin().Delete(ctx, req.Id)
	return
}

func (a *cAdmin) Update(ctx context.Context, req *backend.AdminUpdateReq) (res *backend.AdminUpdateRes, err error) {
	err = service.Admin().Update(ctx, model.AdminUpdateInput{
		Id: req.Id,
		AdminAddUpdateBase: model.AdminAddUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			IsAdmin:  req.IsAdmin,
			RoleIds:  req.RoleIds,
		},
	})
	return
}

// Index admin list
func (a *cAdmin) List(ctx context.Context, req *backend.AdminGetListCommonReq) (res *backend.AdminGetListCommonRes, err error) {
	getListRes, err := service.Admin().GetList(ctx, model.AdminGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.AdminGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, err
}
