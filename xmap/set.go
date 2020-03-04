package xmap

import (
	"github.com/go-board/x-go/xcontainer/set"
)

func KeySet(m interface{}) set.AnyHashSet {
	return set.NewAnyHashSet(Keys(m))
}

func ValueSet(m interface{}) set.AnyHashSet {
	return set.NewAnyHashSet(Values(m))
}
