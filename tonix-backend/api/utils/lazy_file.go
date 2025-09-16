package utils

import (
	"io"
	"os"
	"sync"
)

type AutoclosingFile struct {
	mu     sync.Mutex
	file   *os.File
	closed bool
}

func CreateAutoclosingFile(file *os.File) *AutoclosingFile {
	return &AutoclosingFile{
		mu:     sync.Mutex{},
		file:   file,
		closed: false,
	}
}

func (l *AutoclosingFile) Read(b []byte) (int, error) {
	if l.closed {
		return 0, io.EOF
	}

	n, err := l.file.Read(b)

	if err == io.EOF {
		l.file.Close()
		l.closed = true
	}

	return n, err
}

func (l *AutoclosingFile) Seek(offset int64, whence int) (int64, error) {
	if l.closed {
		file, err := os.Open(l.file.Name())

		if err != nil {
			return 0, err
		}

		l.file = file
		l.closed = false
	}

	return l.file.Seek(offset, whence)
}

func (l *AutoclosingFile) Close() error {
	if l.closed {
		return nil
	}

	l.closed = true
	return l.file.Close()
}
