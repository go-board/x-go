package mdkey

import (
	"testing"
)

func TestRandom(t *testing.T) {
	id := RandomIDGenerator.RequestID()
	t.Logf("random request id is: %s", id)
}

func TestUUID(t *testing.T) {
	id := UUIDGenerator.RequestID()
	t.Logf("uuid request id is: %s", id)
}
