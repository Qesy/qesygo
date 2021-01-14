package qesygo

import (
	"github.com/golang/protobuf/proto"
)

func Unmarshal(data []byte, pbStruct proto.Message) error {
	return proto.Unmarshal(data[4:], pbStruct)
}

func Marshal(ProtoId int32, pbStruct proto.Message) ([]byte, error) {
	data := []byte{}
	if msg, err := proto.Marshal(pbStruct); err != nil {
		return msg, err
	} else {
		PidByte := IntToBytes(ProtoId)
		data = append(data, PidByte...)
		data = append(data, msg...)
	}
	return data, nil
}
