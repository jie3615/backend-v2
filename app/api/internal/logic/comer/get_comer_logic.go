package comer

import (
	"context"
	"errors"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comer"
	"metaLand/data/model/comerprofile"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewGetComerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerLogic {
	return &GetComerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerLogic) GetComer() (resp *types.ComerResponse, err error) {
	// todo: add your logic here and delete this line// 从token中获取用户comerid
	comerInfo, ok := l.ctx.Value("comerInfo").(*comer.Comer)
	if !ok {
		return nil, errors.New("user not found")
	}

	resp = &types.ComerResponse{
		Activation:     false,
		Address:        comerInfo.Address,
		Avatar:         "",
		Banner:         "",
		CustomDomain:   "",
		Id:             int(comerInfo.ID),
		InvitationCode: "",
		IsConnected:    false,
		IsSeted:        false,
		Location:       "",
		Name:           "",
		TimeZone:       "America/Los_Angeles",
	}

	comerProfile, err := comerprofile.FindComerProfile(l.svcCtx.DB, uint64(comerInfo.ID))
	if err != nil {
		return nil, err
	}
	if comerProfile != nil {
		resp.Avatar = comerProfile.Avatar
		resp.Name = comerProfile.Name
		resp.Location = comerProfile.Location
		resp.IsSeted = true
		resp.Location = comerProfile.Website
	}

	return resp, nil
}
