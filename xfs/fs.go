package xfs

import (
	"io"
	"os"

	"github.com/spf13/afero"
)

// FileSystem is the filesystem interface.
type FileSystem = afero.Fs

// File represents a file in the filesystem.
type File = afero.File

var osFs = afero.NewOsFs()

// PwdFs create a new file system base on current work directory
func PwdFs() (FileSystem, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return afero.NewBasePathFs(afero.NewOsFs(), pwd), nil
}

// MemoryFs create an in-memory file system, this is useful for static file serving in a http server
func MemoryFs() FileSystem {
	return afero.NewMemMapFs()
}

func OsWriteFile(filename string, data []byte) error {
	return WriteFile(osFs, filename, data)
}

func WriteFile(fs FileSystem, filename string, data []byte) error {
	return afero.WriteFile(fs, filename, data, 0644)
}

// OsCopyFile will copy file content from src to dst
// TODO: if in *nix, use sendfile instead
func OsCopyFile(dst, src string, mode os.FileMode) (int64, error) {
	return CopyFile(osFs, dst, osFs, src, mode)
}

func CopyFile(dstFs FileSystem, dst string, srcFs FileSystem, src string, mode os.FileMode) (int64, error) {
	srcFile, err := srcFs.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()
	dstFile, err := dstFs.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_RDWR, mode)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

// OsExists returns true if the path is existing
func OsExists(path string) (bool, error) {
	return Exists(osFs, path)
}

// Exists returns true if the path is existing
func Exists(fs FileSystem, path string) (bool, error) {
	_, err := fs.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
