package xnet

import (
	"net"
	"time"
)

// keepaliveListener wraps *net.TCPListener.
type keepaliveListener struct {
	*net.TCPListener
	timeout time.Duration
}

// Accept returns a TCP keep-alive enabled connection.
func (l keepaliveListener) Accept() (c net.Conn, err error) {
	conn, err := l.AcceptTCP()
	if err != nil {
		return
	}
	SetKeepAlive(conn, l.timeout)
	return conn, nil
}

// KeepAliveListener returns listener if l is a
// *net.TCPListener.  Otherwise, l is returned without change.
func KeepAliveListener(l net.Listener, keepalive time.Duration) net.Listener {
	if tl, ok := l.(*net.TCPListener); ok {
		return keepaliveListener{TCPListener: tl, timeout: keepalive}
	}
	return l
}

// SetKeepAlive enables TCP keep-alive if c is a *net.TCPConn.
func SetKeepAlive(c net.Conn, keepalive time.Duration) {
	tc, ok := c.(*net.TCPConn)
	if !ok {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(keepalive)
}
