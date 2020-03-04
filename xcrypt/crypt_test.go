package xcrypt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAes(t *testing.T) {
	token := []byte("12312312904812903812985qs90127489121314523")
	iv := token
	encrypted, err := AesEncryptRaw([]byte("hello"), token, iv)
	require.Nil(t, err, "err must be nil")
	decrypted, err := AesDecryptRaw(encrypted, token, iv)
	require.Nil(t, err, "err must be nil")
	require.EqualValues(t, "hello", decrypted)
}
