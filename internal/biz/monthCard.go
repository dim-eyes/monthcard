package biz

import (
	"context"
	pb "monthCard/api/monthCard/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type MonthCard struct {
	UserId     int64
	ExpireTime int64
}

type MonthCardReward struct {
}

type FormatDays struct {
	Day       int64  `json:"day"`
	Status    int64  `json:"status"`
	RewardId  int64  `json:"rewardId"`
	RewardUrl string `json:"rewardUrl"`
}

type MonthCardList struct {
	Code int64
	Msg  string
	Data string
}

type MonthCardRepo interface {
	OpenMonthCard(ctx context.Context, req *pb.OpenMonthCardRequest) (*pb.OpenMonthCardReply, error)
	GetMonthCardRward(ctx context.Context, req *pb.GetMonthCardRewardRequest) (*pb.GetMonthCardRewardReply, error)
	GetMonthCardList(ctx context.Context, req *pb.GetMonthCardListRequest) (*pb.GetMonthCardListReply, error)
}

type MonthCardUsecase struct {
	repo MonthCardRepo
	log  *log.Helper
}

func NewMonthCardUsecase(repo MonthCardRepo, logger log.Logger) *MonthCardUsecase {
	return &MonthCardUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *MonthCardUsecase) OpenMonthCard(ctx context.Context, req *pb.OpenMonthCardRequest) (*pb.OpenMonthCardReply, error) {
	uc.log.WithContext(ctx).Infof("OpenMonthCard: %v", req.UserId)
	return uc.repo.OpenMonthCard(ctx, req)
}

func (uc *MonthCardUsecase) GetMonthCardRward(ctx context.Context, req *pb.GetMonthCardRewardRequest) (*pb.GetMonthCardRewardReply, error) {
	uc.log.WithContext(ctx).Infof("GetMonthCardRward: %v")
	return uc.repo.GetMonthCardRward(ctx, req)
}

func (uc *MonthCardUsecase) GetMonthCardList(ctx context.Context, req *pb.GetMonthCardListRequest) (*pb.GetMonthCardListReply, error) {
	uc.log.WithContext(ctx).Infof("GetMonthCardList: %v")
	return uc.repo.GetMonthCardList(ctx, req)
}
