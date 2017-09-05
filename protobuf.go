package QesyGo

import (
	"github.com/golang/protobuf/proto"
)

func Unmarshal(data []byte, pbStruct proto.Message) error {
	return proto.Unmarshal(data[4:], pbStruct)
}

func Marshal(ProtoId int32, pbStruct proto.Message) ([]byte, error) {
	return proto.Marshal(pbStruct)
}
