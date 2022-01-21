package arpcx

import (
	"github.com/vmihailenco/msgpack/v5"
)

type MsgpackCodec struct{}

func (m *MsgpackCodec) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (m *MsgpackCodec) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}

func (m *MsgpackCodec) String() string {
	return "msgpack"
}
