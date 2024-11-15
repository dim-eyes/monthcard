package constant

import (
	pb "monthCard/api/monthCard/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

const NOT_FOUND = 201
const MONTH_CARD_HAVE_OPENED = 202
const MONTH_CARD_REWARD_HAVE_RECEIVED = 203

var (
	USER_NOT_FOUND                        = errors.New(NOT_FOUND, pb.ErrorCode_USER_NOT_FOUND.String(), "111")
	MONTH_CARD_HAVE_OPENED_ERROR          = errors.New(MONTH_CARD_HAVE_OPENED, pb.ErrorCode_MONTH_CARD_HAVE_OPENED_ERROR.String(), "111")
	MONTH_CARD_REWARD_HAVE_RECEIVED_ERROR = errors.New(MONTH_CARD_REWARD_HAVE_RECEIVED, pb.ErrorCode_MONTH_CARD_REWARD_HAVE_RECEIVED_ERROR.String(), "111")
)
