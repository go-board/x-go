package xctx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type User struct {
	Name string
}

func TestTypeName(t *testing.T) {
	typeName := TypeName(User{})
	t.Log(typeName)
	ptrTypeName := TypeName(&User{})
	t.Log(ptrTypeName)
}

func TestRead(t *testing.T) {
	ctx := NewTyped(context.Background())
	ctx.With(User{Name: "lixiaohui"})
	user, ok := ctx.ReadTyped(User{})
	require.True(t, ok)
	realUser, ok := user.(User)
	require.True(t, ok)
	require.Equal(t, "lixiaohui", realUser.Name)
}
