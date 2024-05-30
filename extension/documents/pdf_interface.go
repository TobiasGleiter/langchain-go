package extension

type pdf interface {
	Open(filePath string) error
	IsEncrypted() bool
	NumPages() int
	Close() error
  }