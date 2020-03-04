package xnet

import (
	"net"
	"time"
)

type timeoutListener struct {
	*net.TCPListener
	timeout time.Duration
}

func (l timeoutListener) Accept() (net.Conn, error) {
	conn, err := l.TCPListener.Accept()
	if err != nil {
		return nil, err
	}
	SetTimeout(conn, l.timeout)
	return conn, nil
}

func TimeoutListener(l net.Listener, timeout time.Duration) net.Listener {
	if tl, ok := l.(*net.TCPListener); ok {
		return timeoutListener{TCPListener: tl, timeout: timeout}
	}
	return l
}

func SetTimeout(c net.Conn, timeout time.Duration) {
	tc, ok := c.(*net.TCPConn)
	if !ok {
		return
	}
	deadline := time.Now().Add(timeout)
	tc.SetDeadline(deadline)
	tc.SetReadDeadline(deadline)
	tc.SetWriteDeadline(deadline)
}
