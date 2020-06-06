package iocount

import (
	"bytes"
	"errors"
)

var (
	errTest = errors.New("test")
)

type TestReadWriteCloser struct {
	buf        *bytes.Buffer
	shouldFail bool
	closed     bool
}

func (t *TestReadWriteCloser) Read(p []byte) (n int, err error) {
	if t.shouldFail {
		return 0, errTest
	}
	return t.buf.Read(p)
}

func (t *TestReadWriteCloser) Write(p []byte) (n int, err error) {
	if t.shouldFail {
		return 0, errTest
	}
	return t.buf.Write(p)
}

func (t *TestReadWriteCloser) Close() error {
	t.closed = true
	if t.shouldFail {
		return errTest
	}
	return nil
}

func NewTestReadWriteCloser() *TestReadWriteCloser {
	return &TestReadWriteCloser{buf: &bytes.Buffer{}}
}
