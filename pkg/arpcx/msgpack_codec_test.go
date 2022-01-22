package arpcx

import (
	"testing"

	"gotest.tools/assert"
)

func TestMarshalAndUnarshal(t *testing.T) {
	type Per struct {
		Pro string `msgpack:"pro" json:"pro"`
		Vv  []int  `msgpack:"vv" json:"vv"`
	}
	var p Per
	p.Pro = "zzz"
	p.Vv = []int{1, 2, 3, 40000000000000000}

	m := &MsgpackCodec{}

	bt, err := m.Marshal(p)

	assert.Assert(t, err == nil)

	var p1 Per
	err = m.Unmarshal(bt, &p1)

	assert.Assert(t, err == nil)

	assert.Assert(t, p1.Pro == "zzz")
}

func BenchmarkMsgpack(b *testing.B) {
	var p struct {
		Pro string `msgpack:"pro" json:"pro"`
		Vv  []int  `msgpack:"vv" json:"vv"`
	}
	p.Pro = "zzz"
	p.Vv = []int{1, 2, 3, 40000000000000000}

	m := &MsgpackCodec{}
	bt, _ := m.Marshal(p)

	// bt, _ := json.Marshal(p)

	for n := 0; n < b.N; n++ {
		if err := m.Unmarshal(bt, &p); err != nil {
			b.Error(err)
		}
		// json.Unmarshal(bt, &p)
	}
}
