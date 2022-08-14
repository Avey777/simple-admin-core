package logic

import (
	"context"
	"github.com/suyuan32/simple-admin-core/api/common/errorx"

	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetRoleByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByIdLogic {
	return &GetRoleByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleByIdLogic) GetRoleById(in *core.IDReq) (*core.RoleInfo, error) {
	var role model.Role
	result := l.svcCtx.DB.Where("id = ?", in.ID).First(&role)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, errorx.GetInfoFailed)
	}
	return &core.RoleInfo{
		Id:            uint64(role.ID),
		Name:          role.Name,
		Value:         role.Value,
		DefaultRouter: role.DefaultRouter,
		Status:        role.Status,
		Remark:        role.Remark,
		OrderNo:       role.OrderNo,
		CreateAt:      role.CreatedAt.UnixMilli(),
	}, nil
}
