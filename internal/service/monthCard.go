package service

import (
	"context"
	pb "monthCard/api/monthCard/v1"
	"monthCard/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type MonthCardService struct {
	pb.UnimplementedMonthCardServer
	uc  *biz.MonthCardUsecase
	log *log.Helper
}

func NewMonthCardService(uc *biz.MonthCardUsecase, logger log.Logger) *MonthCardService {
	return &MonthCardService{uc: uc, log: log.NewHelper(logger)}
}

func (s *MonthCardService) OpenMonthCard(ctx context.Context, req *pb.OpenMonthCardRequest) (*pb.OpenMonthCardReply, error) {
	return s.uc.OpenMonthCard(ctx, req)

}

func (s *MonthCardService) GetMonthCardRward(ctx context.Context, req *pb.GetMonthCardRewardRequest) (*pb.GetMonthCardRewardReply, error) {
	return s.uc.GetMonthCardRward(ctx, req)
}

func (s *MonthCardService) GetMonthCardList(ctx context.Context, req *pb.GetMonthCardListRequest) (*pb.GetMonthCardListReply, error) {
	return s.uc.GetMonthCardList(ctx, req)
}
