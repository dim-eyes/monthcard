package encoder

import (
	"encoding/json"
	"net/http"

	"github.com/go-kratos/kratos/v2/encoding"
	"google.golang.org/protobuf/proto"
)

type Response struct {
	Code     int         `json:"code"`
	Reason   string      `json:"reason"`
	Message  string      `json:"message"`
	Metadata interface{} `json:"metadata"`
}

func RespEncoder(w http.ResponseWriter, r *http.Request, i interface{}) error {
	codec := encoding.GetCodec("json")
	messageMap := make(map[string]interface{})
	messageStr, _ := codec.Marshal(i.(proto.Message))
	_ = codec.Unmarshal(messageStr, &messageMap)
	// Notice 加上下面的逻辑，如果返回的结果中只有一个字段，那会将结果直接返回到message中
	// Notice 看实际情况，一般不会这样做的
	/*
			比如：返回的结果是："desc":{"name":"whw"}，只有desc一个字段
			那么 结果会是这样:
			{
		      "code": 200,
		      "reason": "",
		      "message": {"name": "whw"}
			}
	*/
	//if len(messageMap) == 1 {
	//	for _, v := range messageMap {
	//		i = v
	//	}
	//}

	resp := Response{
		Code: 200,
	}

	if msg, ok := messageMap["metadata"]; ok {
		i = msg
	}
	message, err := codec.Marshal(i)
	_ = json.Unmarshal(message, &resp.Metadata)
	if err != nil {
		return err
	}

	data, err := codec.Marshal(resp)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	w.WriteHeader(int(data[0]))
	return nil
}
