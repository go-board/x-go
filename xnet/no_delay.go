package xnet

import (
	"net"
)

type noDelayListener struct {
	net.Listener
}

// NoDelayedListener wrap net.Listener accept only no_delay connection without Nagle algorithm,
// so connection can send packet as soon as possible.
func NoDelayedListener(ln net.Listener) net.Listener {
	return &noDelayListener{ln}
}

func (n *noDelayListener) Accept() (net.Conn, error) {
	conn, err := n.Listener.Accept()
	if err != nil {
		return nil, err
	}
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		if err := tcpConn.SetNoDelay(true); err != nil {
			return nil, err
		}
	}
	return conn, nil
}
