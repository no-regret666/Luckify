package logic

import (
	"Luckify/app/checkin/model"
	"Luckify/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"time"

	"Luckify/app/checkin/cmd/rpc/internal/svc"
	"Luckify/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCheckinRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCheckinRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCheckinRecordLogic {
	return &AddCheckinRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCheckinRecordLogic) AddCheckinRecord(in *pb.AddCheckinRecordReq) (*pb.AddCheckinRecordResp, error) {
	dbRecord := &model.CheckinRecord{}
	_ = copier.Copy(dbRecord, in)
	dbRecord.LastCheckinDate.Time = time.Unix(in.LastCheckinDate, 0)
	dbRecord.LastCheckinDate.Valid = true

	err := l.svcCtx.CheckinRecordModel.Insert(l.ctx, nil, dbRecord)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_CHECKIN_RECORD_INSERT_ERROR), "CheckinRecordModel Insert: %+v,err: %v", dbRecord, err)
	}

	return &pb.AddCheckinRecordResp{}, nil
}
