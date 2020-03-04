package xfs

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPwdFs(t *testing.T) {
	fs, err := PwdFs()
	require.Nil(t, err, "err must be nil")
	f, err := fs.Open("/fs_test.go")
	require.Nil(t, err, "open file should success")
	data, err := ioutil.ReadAll(f)
	require.Nil(t, err, "read file should success")
	t.Log(string(data))
}

func TestOsCopyFile(t *testing.T) {
	_, err := OsCopyFile("testdata/b.txt", "testdata/a.txt", 0644)
	require.Nil(t, err, "copy file should success")
}
