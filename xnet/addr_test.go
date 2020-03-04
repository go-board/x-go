package xnet

import (
	"testing"
)

func TestPrivateAddress(t *testing.T) {
	p, err := PrivateAddress()
	t.Log(p, err)
}
