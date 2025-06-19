package logic

import (
	"context"
	"github.com/Yamon955/ShortVideo/protocol/interaction/pb"
)

func (h *handlerImpl) HandleReport(ctx context.Context, req *pb.ReportReq) (*pb.ReportRsp, error) {
	return &pb.ReportRsp{}, nil
}
