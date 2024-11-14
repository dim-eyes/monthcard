package data

import (
	"context"
	"encoding/json"
	"fmt"
	pb "monthCard/api/monthCard/v1"
	"monthCard/internal/biz"
	"monthCard/internal/data/constant"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type MonthCard struct {
	Id         int64     `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Userid     int64     `gorm:"column:userid;type:int(11);comment:用户id;NOT NULL" json:"userid"`
	Days       string    `gorm:"column:days;type:varchar(255);comment:天数" json:"days"`
	Issue      int64     `gorm:"column:issue;type:int(11)" json:"issue"`
	ExpireTime int64     `gorm:"column:expire_time;type:int(11);comment:过期时间" json:"expire_time"`
	SaveMoney  int64     `gorm:"column:save_money;type:int(11);comment:节省金钱" json:"save_money"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`
}
type monthCardRepo struct {
	data   *Data
	logger *log.Helper
}

func (r *MonthCard) TableName() string {
	return "monthCard"
}

func NewMonthCardRepo(data *Data, logger log.Logger) biz.MonthCardRepo {
	return &monthCardRepo{data: data, logger: log.NewHelper(logger)}
}

func (r *monthCardRepo) OpenMonthCard(ctx context.Context, req *pb.OpenMonthCardRequest) (*pb.OpenMonthCardReply, error) {
	fmt.Println(time.Now().Unix())
	openStatus := r.IsOpenMonthCard(ctx, req.UserId)
	if openStatus {
		return nil, constant.MONTH_CARD_HAVE_OPENED_ERROR
	}
	currentTime := time.Now().Unix()
	expireTime := currentTime + (constant.DAY_OF_MONTH_CARD * 24 * 60 * 60)

	item := MonthCard{
		Userid:     req.UserId,
		ExpireTime: expireTime,
	}
	res := r.data.db.WithContext(ctx).Create(&item)
	/*res := r.data.db.WithContext(ctx).Model(rm).Where("userid = ?", req.UserId).First(rm)*/
	if res.Error != nil {
		return nil, res.Error
	}

	setData := map[string]interface{}{"userId": req.UserId, "days": constant.DAY_OF_MONTH_CARD, "expireTime": expireTime, "saveMoney": 0}
	r.data.rdb.HMSet(ctx, fmt.Sprintf("monthCard:userId:%d", req.UserId), setData)
	r.data.rdb.ExpireAt(ctx, fmt.Sprintf("monthCard:userId:%d", req.UserId), time.Unix(int64(expireTime), 0))
	r.logger.WithContext(ctx).Info("gormDB: GetMonthCard, userId: ", res)
	return &pb.OpenMonthCardReply{UserId: req.UserId, ExpireTime: expireTime}, nil
}

func (r *monthCardRepo) GetMonthCardRward(ctx context.Context, req *pb.GetMonthCardRewardRequest) (*pb.GetMonthCardRewardReply, error) {
	return &pb.GetMonthCardRewardReply{}, nil
}

func (r *monthCardRepo) GetMonthCardList(ctx context.Context, req *pb.GetMonthCardListRequest) (*pb.GetMonthCardListReply, error) {
	list := r.getShowDays(constant.DAY_OF_MONTH_CARD)
	r.logger.WithContext(ctx).Info("gormDB: GetMonthCard, userId: ", list)
	var cardData []*pb.GetMonthCardListReplyCardData
	for _, v := range list {
		ts := biz.FormatDays{}
		json.Unmarshal([]byte(v), &ts)
		fmt.Println(ts.Day)
		cardData = append(cardData, &pb.GetMonthCardListReplyCardData{Day: ts.Day, Status: ts.Status, RewardId: ts.RewardId, RewardUrl: ts.RewardUrl})
	}
	return &pb.GetMonthCardListReply{List: cardData}, nil
}

func (r *monthCardRepo) IsOpenMonthCard(ctx context.Context, userId int64) bool {
	status := r.data.rdb.HGet(ctx, fmt.Sprintf("monthCard:userId:%d", userId), "userId")
	fmt.Println(status.Val())
	return status.Val() != ""
}

func (r *monthCardRepo) getShowDays(currentDay int64) []string {

	var list = make([]string, 0)
	RewardMap := []string{
		constant.DAY_ONE_REWARD,
		constant.DAY_ONE_REWARD,
		constant.DAY_ONE_REWARD,
		constant.DAY_ONE_REWARD,
		constant.DAY_ONE_REWARD,
		constant.DAY_ONE_REWARD,
		constant.DAY_ONE_REWARD,
		constant.DAY_ONE_REWARD,
		constant.DAY_ONE_REWARD,
		constant.DAY_ONE_REWARD,
		constant.DAY_ONE_REWARD,
		constant.DAY_ONE_REWARD,
		constant.DAY_ONE_REWARD,
		constant.DAY_ONE_REWARD,
		constant.DAY_ONE_REWARD,
	}

	for i := 1; i <= constant.DAY_OF_MONTH_CARD; i++ {
		//messageMap := make(map[string]interface{})
		ts := biz.FormatDays{}
		ts.Day = int64(i)
		if int64(i) > currentDay {
			ts.Status = constant.REWARD_STATUS_WAIT
		}

		if currentDay > int64(i) {
			ts.Status = constant.REWARD_STATUS_OVERTRUE

		}
		ts.Status = constant.REWARD_STATUS_NOT_RECEIVED
		ts.RewardId = 1
		ts.RewardUrl = RewardMap[i-1]
		jsonData, err := json.Marshal(ts)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(jsonData)
		//json.Unmarshal(jsonData, &messageMap)
		list = append(list, string(jsonData))
	}
	return list
}
