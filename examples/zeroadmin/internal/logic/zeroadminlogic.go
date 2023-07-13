package logic

import (
	"context"

	"github.com/quarkcms/quark-go/v2/examples/zeroadmin/internal/svc"
	"github.com/quarkcms/quark-go/v2/examples/zeroadmin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ZeroadminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZeroadminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZeroadminLogic {
	return &ZeroadminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ZeroadminLogic) Zeroadmin(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	return
}
