package xdebug

import (
	"os"
	"testing"
)

func TestDebug(t *testing.T) {
	Debug(1, 2, 3, "", 1+1i, 1_000)
}

func TestFDebug(t *testing.T) {
	FDebug(os.Stderr, 1, 2, 3, "", 1+1i, 1_000)
}

func TestSDebug(t *testing.T) {
	t.Log(SDebug(1, 2, 3, "", 1+1i, 1_000))
}
