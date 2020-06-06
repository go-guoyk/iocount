package iocount

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReadWriter(t *testing.T) {
	var err error
	var o *TestReadWriteCloser
	var r ReadWriter
	var buf []byte
	var n int

	o = NewTestReadWriteCloser()
	r = NewReadWriter(o)

	assert.Equal(t, o, r.ReadWriter())

	buf = []byte("hello")

	n, err = r.Write(buf)
	assert.Equal(t, len(buf), n)
	assert.NoError(t, err)
	n, err = r.Write(buf)
	assert.Equal(t, len(buf), n)
	assert.NoError(t, err)

	n, err = r.Read(buf)
	assert.Equal(t, len(buf), n)
	assert.NoError(t, err)
	n, err = r.Read(buf)
	assert.Equal(t, len(buf), n)
	assert.NoError(t, err)

	assert.Equal(t, int64(len(buf))*2, r.ReadCount())
	assert.Equal(t, int64(len(buf))*2, r.WriteCount())

	/////

	o = NewTestReadWriteCloser()
	o.shouldFail = true
	r = NewReadWriter(o)

	assert.Equal(t, o, r.ReadWriter())

	buf = []byte("hello")

	n, err = r.Write(buf)
	assert.Equal(t, 0, n)
	assert.Equal(t, errTest, err)
	n, err = r.Write(buf)
	assert.Equal(t, 0, n)
	assert.Equal(t, errTest, err)

	n, err = r.Read(buf)
	assert.Equal(t, 0, n)
	assert.Equal(t, errTest, err)
	n, err = r.Read(buf)
	assert.Equal(t, 0, n)
	assert.Equal(t, errTest, err)

	assert.Equal(t, int64(0), r.WriteCount())
	assert.Equal(t, int64(0), r.ReadCount())
}
