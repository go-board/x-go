package xstrings

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strings"
)

var randReader = rand.Reader

const RandomSet = "0123456789ABCDEF"

// Random creates new string based on strSet. It uses crypto/rand as the
// random number generator. error is the one returned by rand.Int
func Random(strSet string, length int) (string, error) {
	if length == 0 || strSet == "" {
		return "", nil
	}
	set := []rune(strSet)
	bigLen := big.NewInt(int64(len(set)))

	res := make([]rune, length)
	for i := 0; i < length; i++ {
		n, err := rand.Int(randReader, bigLen)
		if err != nil {
			return "", err
		}
		res[i] = set[n.Int64()]
	}

	return string(res), nil
}

// FastRandom creates new string with one read, this version is fast than the random version.
func FastRandom(n int) (string, error) {
	i := n/2 + 1
	buf := make([]byte, i)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	str := hex.EncodeToString(buf)
	if n%2 == 1 {
		str = str[:len(str)-1]
	}
	return strings.ToUpper(str), nil
}
