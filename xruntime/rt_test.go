package xruntime

import (
	"testing"
)

func TestCallerName(t *testing.T) {
	name := CallerName()
	t.Log(name)
}

func TestCaller(t *testing.T) {
	t.Log(Caller())
}

func TestBuildPath(t *testing.T) {
	t.Log(BuildPath())
}

func TestStack(t *testing.T) {
	t.Logf("%s\n", Stack(false))
}
