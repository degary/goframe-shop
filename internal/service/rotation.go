// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/degary/goframe-shop/internal/model"
)

type (
	IRotation interface {
		Add(ctx context.Context, in model.RotationAddInput) (out model.RotationAddOutput, err error)
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context, in model.RotationUpdateInput) error
		GetList(ctx context.Context, in model.RotationGetListInput) (out *model.RotationGetListOutput, err error)
	}
)

var (
	localRotation IRotation
)

func Rotation() IRotation {
	if localRotation == nil {
		panic("implement not found for interface IRotation, forgot register?")
	}
	return localRotation
}

func RegisterRotation(i IRotation) {
	localRotation = i
}
