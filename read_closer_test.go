package iocount

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReadCloser(t *testing.T) {
	var err error
	var o *TestReadWriteCloser
	var r ReadCloser
	var buf []byte
	var n int

	o = NewTestReadWriteCloser()
	r = NewReadCloser(o)

	assert.Equal(t, o, r.ReadCloser())

	buf = []byte("hello")

	o.buf.Write(buf)
	o.buf.Write(buf)
	o.buf.Write(buf)

	n, err = r.Read(buf)
	assert.Equal(t, len(buf), n)
	assert.NoError(t, err)
	n, err = r.Read(buf)
	assert.Equal(t, len(buf), n)
	assert.NoError(t, err)

	assert.Equal(t, int64(len(buf))*2, r.ReadCount())

	err = r.Close()
	assert.NoError(t, err)
	assert.True(t, o.closed)

	/////

	o = NewTestReadWriteCloser()
	o.shouldFail = true
	r = NewReadCloser(o)

	assert.Equal(t, o, r.ReadCloser())

	buf = []byte("hello")

	o.buf.Write(buf)
	o.buf.Write(buf)
	o.buf.Write(buf)

	n, err = r.Read(buf)
	assert.Equal(t, 0, n)
	assert.Equal(t, errTest, err)
	n, err = r.Read(buf)
	assert.Equal(t, 0, n)
	assert.Equal(t, errTest, err)

	assert.Equal(t, int64(0), r.ReadCount())

	err = r.Close()
	assert.Equal(t, errTest, err)
	assert.True(t, o.closed)
}
