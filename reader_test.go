package iocount

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReader(t *testing.T) {
	var err error
	var o *TestReadWriteCloser
	var r Reader
	var buf []byte
	var n int

	o = NewTestReadWriteCloser()
	r = NewReader(o)

	assert.Equal(t, o, r.Reader())

	buf = []byte("hello")

	o.buf.Write(buf)
	o.buf.Write(buf)

	n, err = r.Read(buf)
	assert.Equal(t, len(buf), n)
	assert.NoError(t, err)
	n, err = r.Read(buf)
	assert.Equal(t, len(buf), n)
	assert.NoError(t, err)

	assert.Equal(t, int64(len(buf))*2, r.ReadCount())

	/////

	o = NewTestReadWriteCloser()
	o.shouldFail = true
	r = NewReader(o)

	assert.Equal(t, o, r.Reader())

	buf = []byte("hello")

	o.buf.Write(buf)
	o.buf.Write(buf)

	n, err = r.Read(buf)
	assert.Equal(t, 0, n)
	assert.Equal(t, errTest, err)
	n, err = r.Read(buf)
	assert.Equal(t, 0, n)
	assert.Equal(t, errTest, err)

	assert.Equal(t, int64(0), r.ReadCount())
}
