package xseq

import (
	"github.com/go-board/x-go/xstrings"
	"github.com/gofrs/uuid"
)

var globalId string

func init() {
	var err error
	globalId, err = xstrings.FastRandom(32)
	if err != nil {
		globalId = UUID()
	}
}

func GlobalID() string {
	return globalId
}

func UUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
