package constant

import (
	pb "monthCard/api/monthCard/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

const NET_WORK_ERR = 201
const MONTH_CARD_HAVE_OPENED = 202
const MONTH_CARD_NOT_OPENED = 203
const MONTH_CARD_REWARD_RECEIVED = 2006
const MONTH_CARD_REWARD_FAILED = 206

var (
	NET_WORK_ERROR                   = errors.New(NET_WORK_ERR, pb.ErrorCode_NET_WORK_ERROR.String(), "111")
	MONTH_CARD_NOT_OPEN              = errors.New(MONTH_CARD_NOT_OPENED, pb.ErrorCode_MONTH_CARD_NOT_OPEN.String(), "111")
	MONTH_CARD_HAVE_OPENED_ERROR     = errors.New(MONTH_CARD_HAVE_OPENED, pb.ErrorCode_MONTH_CARD_HAVE_OPENED_ERROR.String(), "111")
	MONTH_CARD_REWARD_RECEIVED_ERROR = errors.New(MONTH_CARD_REWARD_RECEIVED, pb.ErrorCode_MONTH_CARD_REWARD_RECEIVED_ERROR.String(), "111")
	MONTH_CARD_REWARD_FAILED_ERROR   = errors.New(MONTH_CARD_REWARD_FAILED, pb.ErrorCode_MONTH_CARD_REWARD_FAILED_ERROR.String(), "111")
)
