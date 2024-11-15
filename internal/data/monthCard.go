package data

import (
	"context"
	"encoding/json"
	"fmt"
	pb "monthCard/api/monthCard/v1"
	"monthCard/internal/biz"
	"monthCard/internal/data/constant"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type MonthCard struct {
	Id         int64     `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Userid     int64     `gorm:"column:userid;type:int(11);comment:用户id;NOT NULL" json:"userid"`
	Days       string    `gorm:"column:days;type:varchar(255);comment:天数" json:"days"`
	ExpireTime int64     `gorm:"column:expire_time;type:int(11);comment:过期时间" json:"expire_time"`
	SaveMoney  int64     `gorm:"column:save_money;type:int(11);comment:节省金钱" json:"save_money"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`
}
type monthCardRepo struct {
	data   *Data
	logger *log.Helper
}

type monthCardInfo struct {
	currentDay  int64
	accrueMoney int64
	surplusDay  int64
	days        string
}

func (r *MonthCard) TableName() string {
	return "monthCard"
}

func NewMonthCardRepo(data *Data, logger log.Logger) biz.MonthCardRepo {
	return &monthCardRepo{data: data, logger: log.NewHelper(logger)}
}

func (r *monthCardRepo) OpenMonthCard(ctx context.Context, req *pb.OpenMonthCardRequest) (*pb.OpenMonthCardReply, error) {
	fmt.Println(time.Now().Unix())
	openStatus := r.isOpenMonthCard(ctx, req.UserId)
	if openStatus {
		return nil, constant.MONTH_CARD_HAVE_OPENED_ERROR
	}
	currentTime := r.getCurrentTime()
	expireTime := currentTime + (constant.DAY_OF_MONTH_CARD * 24 * 60 * 60)

	item := MonthCard{
		Userid:     req.UserId,
		ExpireTime: expireTime,
	}
	res := r.data.db.WithContext(ctx).Create(&item)
	if res.Error != nil {
		return nil, res.Error
	}

	setData := map[string]interface{}{"userId": req.UserId, "days": "", "expireTime": expireTime, "saveMoney": 0}
	r.data.rdb.HMSet(ctx, fmt.Sprintf("monthCard:userId:%d", req.UserId), setData)
	r.data.rdb.ExpireAt(ctx, fmt.Sprintf("monthCard:userId:%d", req.UserId), time.Unix(expireTime, 0))
	return &pb.OpenMonthCardReply{UserId: req.UserId, ExpireTime: expireTime}, nil
}

func (r *monthCardRepo) GetMonthCardRward(ctx context.Context, req *pb.GetMonthCardRewardRequest) (*pb.GetMonthCardRewardReply, error) {
	r.logger.WithContext(ctx).Info("GetMonthCardRward, userId: ", req.UserId)

	cardInfo, err := r.getUserMonthCardInfo(ctx, req.UserId)
	if err != nil {
		return nil, constant.MONTH_CARD_REWARD_HAVE_RECEIVED_ERROR
	}
	rewardData := r.getTodayReward(cardInfo.currentDay)
	fmt.Println(rewardData)
	return &pb.GetMonthCardRewardReply{RewardId: rewardData.RewardId, RewardNum: rewardData.RewardNum, RewardName: rewardData.RewardName, RewardUrl: rewardData.RewardUrl}, nil
}

func (r *monthCardRepo) GetMonthCardList(ctx context.Context, req *pb.GetMonthCardListRequest) (*pb.GetMonthCardListReply, error) {

	r.logger.WithContext(ctx).Info("GetMonthCardList, userId: ", req.UserId)

	openStatus := r.isOpenMonthCard(ctx, req.UserId)
	if !openStatus {
		return nil, constant.MONTH_CARD_HAVE_OPENED_ERROR
	}

	cardInfo, err := r.getUserMonthCardInfo(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	fmt.Println(strings.Split(cardInfo.days, ","))

	list := r.getShowDays(cardInfo.currentDay, cardInfo.days)
	var cardData []*pb.GetMonthCardListReplyCardData
	for _, v := range list {
		ts := biz.FormatDays{}
		json.Unmarshal([]byte(v), &ts)
		cardData = append(cardData, &pb.GetMonthCardListReplyCardData{Day: ts.Day, Status: ts.Status, RewardId: ts.RewardId, RewardNum: ts.RewardNum, RewardName: ts.RewardName, RewardUrl: ts.RewardUrl})
	}
	return &pb.GetMonthCardListReply{EconomizeMoney: cardInfo.accrueMoney, SurplusDay: cardInfo.surplusDay, DayList: cardData}, nil
}

func (r *monthCardRepo) isOpenMonthCard(ctx context.Context, userId int64) bool {
	status := r.data.rdb.HGet(ctx, fmt.Sprintf("monthCard:userId:%d", userId), "userId")
	return status.Val() != ""
}

func (r *monthCardRepo) getUserMonthCardInfo(ctx context.Context, userId int64) (*monthCardInfo, error) {
	rm := MonthCard{}
	res := r.data.db.WithContext(ctx).Model(rm).Where("userid = ?", userId).Find(&rm)
	if res.Error != nil {
		return nil, res.Error
	}

	currentTime := r.getCurrentTime()
	diffTime := rm.ExpireTime - currentTime
	surplusDay := diffTime / 86400
	if surplusDay < 1 {
		surplusDay = 0
	} else {
		surplusDay = surplusDay - 1
	}
	return &monthCardInfo{currentDay: constant.DAY_OF_MONTH_CARD - surplusDay, accrueMoney: rm.SaveMoney, surplusDay: surplusDay, days: rm.Days}, nil
}

func (r *monthCardRepo) getTodayReward(currentDay int64) *biz.FormatDays {
	fmt.Println(currentDay)
	return nil
}

func (r *monthCardRepo) getShowDays(currentDay int64, days string) []string {

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
		var haveRewardDay bool
		if days != "" {
			RewardDays := strings.Split(days, ",")
			for _, v := range RewardDays {
				val, _ := strconv.ParseInt(v, 10, 64)
				if int64(i) == val {
					haveRewardDay = true
					break
				}
			}
		}
		ts := biz.FormatDays{}
		ts.Day = int64(i)
		if int64(i) > currentDay {
			ts.Status = constant.REWARD_STATUS_WAIT
		}
		if haveRewardDay {
			ts.Status = constant.REWARD_STATUS_RECEIVED
		}

		if currentDay > int64(i) && !haveRewardDay {
			ts.Status = constant.REWARD_STATUS_OVERTRUE
		}
		ts.RewardId = 1
		ts.RewardNum = 2
		ts.RewardName = ""
		ts.RewardUrl = RewardMap[i-1]
		jsonData, err := json.Marshal(ts)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(jsonData)
		//json.Unmarshal(jsonData, &messageMap)
		list = append(list, string(jsonData))
	}
	return list
}

func (r *monthCardRepo) getCurrentTime() int64 {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	timestamp := today.Unix()
	return timestamp
}
