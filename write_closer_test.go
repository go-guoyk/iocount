package iocount

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWriteCloser(t *testing.T) {
	var err error
	var o *TestReadWriteCloser
	var r WriteCloser
	var buf []byte
	var n int

	o = NewTestReadWriteCloser()
	r = NewWriteCloser(o)

	assert.Equal(t, o, r.WriteCloser())

	buf = []byte("hello")

	n, err = r.Write(buf)
	assert.Equal(t, len(buf), n)
	assert.NoError(t, err)
	n, err = r.Write(buf)
	assert.Equal(t, len(buf), n)
	assert.NoError(t, err)

	assert.Equal(t, int64(len(buf))*2, r.WriteCount())

	err = r.Close()
	assert.NoError(t, err)
	assert.True(t, o.closed)

	/////

	o = NewTestReadWriteCloser()
	o.shouldFail = true
	r = NewWriteCloser(o)

	assert.Equal(t, o, r.WriteCloser())

	buf = []byte("hello")

	n, err = r.Write(buf)
	assert.Equal(t, 0, n)
	assert.Equal(t, errTest, err)
	n, err = r.Write(buf)
	assert.Equal(t, 0, n)
	assert.Equal(t, errTest, err)

	assert.Equal(t, int64(0), r.WriteCount())

	err = r.Close()
	assert.Equal(t, errTest, err)
	assert.True(t, o.closed)
}
