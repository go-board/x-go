package mdkey

import (
	"github.com/go-board/x-go/xstrings"
	"github.com/gofrs/uuid"
)

// RequestIDGenerator generate a unique request id in string format
// 生成字符串类型的唯一请求ID
type RequestIDGenerator interface {
	RequestID() string
}

// RandomIDGenerator is a global id generator for generate unique request id with random way
var RandomIDGenerator RequestIDGenerator = randomIdGenerator{}

// UUIDGenerator is a global id generator for generate unique request id with random uuid way
var UUIDGenerator RequestIDGenerator = uuidGenerator{}

type randomIdGenerator struct{}

func (randomIdGenerator) RequestID() string {
	id, _ := xstrings.FastRandom(32)
	return id
}

type uuidGenerator struct{}

func (uuidGenerator) RequestID() string { return uuid.Must(uuid.NewV4()).String() }
