package xseq

import (
	"github.com/bwmarrin/snowflake"

	"github.com/go-board/x-go/types"
)

var node *snowflake.Node

func NewGenerator(n int64) *snowflake.Node {
	node, _ = snowflake.NewNode(n)
	return node
}

func NextID() types.ID {
	return types.ID(node.Generate())
}
