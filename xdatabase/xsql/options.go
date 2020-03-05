package xsql

import "time"

type ConnectionOptions struct {
	DriverName string
	Dsn        string

	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifeTime time.Duration
}

func NewDefaultConnectionOptions() ConnectionOptions {
	return ConnectionOptions{
		MaxIdleConn:     20,
		MaxOpenConn:     50,
		ConnMaxLifeTime: time.Second * 50,
	}
}

// ConnectionOption is helper function to modify database pool options
type ConnectionOption func(options *ConnectionOptions)

func WithDriverName(name string) ConnectionOption {
	return func(options *ConnectionOptions) { options.DriverName = name }
}

func WithDsn(dsn string) ConnectionOption {
	return func(options *ConnectionOptions) { options.Dsn = dsn }
}

func WithMaxIdleConn(n int) ConnectionOption {
	return func(options *ConnectionOptions) { options.MaxIdleConn = n }
}

func WithMaxOpenConn(n int) ConnectionOption {
	return func(options *ConnectionOptions) { options.MaxOpenConn = n }
}

func WithConnMaxIdleTime(d time.Duration) ConnectionOption {
	return func(options *ConnectionOptions) { options.ConnMaxLifeTime = d }
}
