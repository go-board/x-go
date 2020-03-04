package set

import (
	"golang.org/x/tools/container/intsets"
)

type IntSet struct {
	*intsets.Sparse
}
