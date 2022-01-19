package ip

import (
	"testing"

	"gotest.tools/assert"
)

func TestContains(t *testing.T) {
	ipBill := NewIpBill([]string{"127.0.0.1", "172.0.0.1/24"})

	assert.Assert(t, ipBill.Contains("127.0.0.1"))
	assert.Assert(t, !ipBill.Contains("127.0.0.3"))
	assert.Assert(t, ipBill.Contains("172.0.0.3"))
	assert.Assert(t, ipBill.Contains("172.0.0.4"))
}
