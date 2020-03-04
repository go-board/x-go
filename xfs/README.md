# XFS

文件系统的抽象，支持多种文件系统，本地，内存，Map，（阿里云）。

# FileSystem
```go
package xfs

import (
    "os"
    "time"
)

// Fs is the filesystem interface.
//
// Any simulated or real filesystem should implement this interface.
type Fs interface {
	// Create creates a file in the filesystem, returning the file and an
	// error, if any happens.
	Create(name string) (File, error)

	// Mkdir creates a directory in the filesystem, return an error if any
	// happens.
	Mkdir(name string, perm os.FileMode) error

	// MkdirAll creates a directory path and all parents that does not exist
	// yet.
	MkdirAll(path string, perm os.FileMode) error

	// Open opens a file, returning it or an error, if any happens.
	Open(name string) (File, error)

	// OpenFile opens a file using the given flags and the given mode.
	OpenFile(name string, flag int, perm os.FileMode) (File, error)

	// Remove removes a file identified by name, returning an error, if any
	// happens.
	Remove(name string) error

	// RemoveAll removes a directory path and any children it contains. It
	// does not fail if the path does not exist (return nil).
	RemoveAll(path string) error

	// Rename renames a file.
	Rename(oldname, newname string) error

	// Stat returns a FileInfo describing the named file, or an error, if any
	// happens.
	Stat(name string) (os.FileInfo, error)

	// The name of this FileSystem
	Name() string

	//Chmod changes the mode of the named file to mode.
	Chmod(name string, mode os.FileMode) error

	//Chtimes changes the access and modification times of the named file
	Chtimes(name string, atime time.Time, mtime time.Time) error
}
```
## File
```go
package xfs

import (
    "io"
    "os"
)

// File represents a file in the filesystem.
type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	io.WriterAt

	Name() string
	Readdir(count int) ([]os.FileInfo, error)
	Readdirnames(n int) ([]string, error)
	Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size int64) error
	WriteString(s string) (ret int, err error)
}
```
## PwdFs 当前目录作为根目录
`func PwdFs() (FileSystem, error)`
## MemoryFs 内存文件系统
`func MemoryFs() FileSystem`
## WriteFile 写文件
```go
package xfs
func WriteFile(fs FileSystem, filename string, data []byte) error
func OsWriteFile(filename string, data []byte) error
```
## CopyFile 拷贝文件
```go
package xfs
func CopyFile(dstFs FileSystem, dst string, srcFs FileSystem, src string, mode os.FileMode) (int64, error)
func OsCopyFile(dst, src string, mode os.FileMode) (int64, error)
```
## Exists 文件是否存在
```go
package xfs
func Exists(fs FileSystem, path string) (bool, error)
func OsExists(path string) (bool, error)
```
