package xos

import (
	"os"
)

var hostname string

func Hostname() string {
	if hostname != "" {
		return hostname
	}
	hostname, _ = os.Hostname()
	return hostname
}
