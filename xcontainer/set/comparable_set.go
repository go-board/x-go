package set

import (
	"github.com/google/btree"

	"github.com/go-board/x-go/types"
)

type ComparableHashSet struct {
	store map[types.Comparable]struct{}
}

type ComparableBtreeSet struct {
	tree *btree.BTree
}
