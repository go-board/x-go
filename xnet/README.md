# XNet
标准库实现了大量的net函数，这里作为补充，提供了`LimitedListener`和`KeepaliveListener`, 以及一个辅助函数`PrivateAddress`获取内网地址。


### [xhttp](./xhttp/README.md)

## LimitedListener
对最大同时连接数做了上限限制 `n`
```go
package xnet

type limitedListener struct {
	net.Listener
	sem       chan struct{}
	closeOnce sync.Once     // ensures the done chan is only closed once
	done      chan struct{} // no values sent; closed when Close is called
}

func LimitListener(l net.Listener, n int) net.Listener {
	return &limitedListener{
		Listener: l,
		sem:      make(chan struct{}, n),
		done:     make(chan struct{}),
	}
}
```

## KeepaliveListener
会定时发送keepalive message的tcpListener
```go
package xnet

type keepaliveListener struct {
	*net.TCPListener
	timeout time.Duration
}
// 创建一个新的keepalive的listener，keepalive间隔为keepalive，如果listener不是*TcpListener，则会原样返回。
func KeepAliveListener(l net.Listener, keepalive time.Duration) net.Listener
// 设置net.TcpConn keepalive的时间间隔
func SetKeepAlive(c net.Conn, keepalive time.Duration)
```

## TimeoutListener
会自动设置超时间的tcpListener
```go
package xnet

type timeoutListener struct {
	*net.TCPListener
	timeout time.Duration
}
// 创建一个新的具有超时时间的listener，超时时间为timeout，如果listener不是*TcpListener，则会原样返回。
func TimeoutListener(l net.Listener, timeout time.Duration) net.Listener
// 设置net.TcpConn的超时时间
func SetTimeout(c net.Conn, timeout time.Duration)
```
