package arpcx

import (
	"net"
	"time"

	"github.com/lesismal/arpc"
)

func CreateArpcClientPool(listenAddr string, size int) (*arpc.ClientPool, error) {
	return arpc.NewClientPool(func() (net.Conn, error) {
		return net.DialTimeout("tcp", listenAddr, time.Second*5)
	}, size)
}
