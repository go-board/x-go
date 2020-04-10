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

func TestBCrypt(t *testing.T) {
	encoded, err := BCryptHash([]byte("Hello,world"))
	require.Nil(t, err, "bcrypt hash must success")
	ok := BCryptValidate([]byte("Hello,world"), encoded)
	require.True(t, ok, "validate must be successful")
}

//BenchmarkBCrypt
//BenchmarkBCrypt/hash
//BenchmarkBCrypt/hash-12         	      19	  56413888 ns/op	    5277 B/op	      11 allocs/op
//BenchmarkBCrypt/validate
//BenchmarkBCrypt/validate-12     	      20	  56294324 ns/op	    5280 B/op	      15 allocs/op
func BenchmarkBCrypt(b *testing.B) {
	b.Run("hash", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = BCryptHash([]byte("Hello,world12345"))
		}
	})

	b.Run("validate", func(b *testing.B) {
		encoded, _ := BCryptHash([]byte("Hello,world"))
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			BCryptValidate([]byte("Hello,world12345"), encoded)
		}
	})
}
