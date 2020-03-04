package xfs

import (
	"os"

	"github.com/spf13/afero"
)

type FileSystemExt interface {
	afero.Fs
	CopyFile(dst, src string) (int64, error)
	ReadFile(f string) ([]byte, error)
	WriteFile(f string, data []byte, perm os.FileMode) error
}

type FileExt interface {
	afero.File
}
