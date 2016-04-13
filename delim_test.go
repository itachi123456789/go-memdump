package structbuffer

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func join(bufs ...[]byte) []byte {
	return bytes.Join(bufs, nil)
}

func TestDelimetedReader_Simple(t *testing.T) {
	data := join([]byte("abc"), delim, []byte("defggg"), delim)
	r := newDelimitedReader(bytes.NewReader(data))

	buf := make([]byte, 100)

	n, err := r.Read(buf)
	require.NoError(t, err)
	assert.Equal(t, 3, n)
	assert.Equal(t, "abc", string(buf[:n]))

	n, err = r.Read(buf)
	assert.Equal(t, 0, n)
	assert.Equal(t, io.EOF, err)

	hasNext := r.Next()
	assert.True(t, hasNext)

	n, err = r.Read(buf)
	require.NoError(t, err)
	assert.Equal(t, 6, n)
	assert.Equal(t, "defggg", string(buf[:n]))

	n, err = r.Read(buf)
	assert.Equal(t, 0, n)
	assert.Equal(t, io.EOF, err)

	hasNext = r.Next()
	assert.False(t, hasNext)
}
