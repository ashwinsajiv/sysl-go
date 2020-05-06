package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNopLogger_FlushLog(t *testing.T) {
	ctx, hook := NewTestContextWithLoggerHook()

	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	require.NoError(t, err)

	l, _ := NewRequestLogger(ctx, req, true)

	resp := http.Response{
		Status:           "",
		StatusCode:       200,
		Proto:            "",
		ProtoMajor:       0,
		ProtoMinor:       0,
		Header:           http.Header{},
		Body:             ioutil.NopCloser(&bytes.Buffer{}),
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
		Uncompressed:     false,
		Trailer:          nil,
		Request:          nil,
		TLS:              nil,
	}
	l.LogResponse(&resp)

	require.Empty(t, hook.Entries)

	// and should be no-ops
	//rw := httptest.NewRecorder()
	//require.Equal(t, rw, l.ResponseWriter(rw))

	l.FlushLog()
	require.NotEmpty(t, hook.Entries)
}

func TestLogger_FlushLog(t *testing.T) {
	ctx, hook := NewTestContextWithLoggerHook()

	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	require.NoError(t, err)

	l, _ := NewRequestLogger(ctx, req, true)

	resp := http.Response{
		Status:           "",
		StatusCode:       200,
		Proto:            "",
		ProtoMajor:       0,
		ProtoMinor:       0,
		Header:           http.Header{},
		Body:             ioutil.NopCloser(&bytes.Buffer{}),
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
		Uncompressed:     false,
		Trailer:          nil,
		Request:          nil,
		TLS:              nil,
	}
	l.LogResponse(&resp)

	require.Empty(t, hook.Entries)

	// and should do a new error if we call flush
	//hook.Reset()
	l.FlushLog()
	require.NotEmpty(t, hook.Entries)
	//require.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
}

func TestRequestLogger_NilBody(t *testing.T) {
	ctx, _ := NewTestContextWithLoggerHook()

	req, err := http.NewRequest("DELETE", "http://example.com/foo", nil)
	require.NoError(t, err)

	require.NotPanics(t, func() {
		NewRequestLogger(ctx, req, true)
	})
}

func TestRequestLogger_ResponseWriter(t *testing.T) {
	ctx, hook := NewTestContextWithLoggerHook()

	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	require.NoError(t, err)

	l, _ := NewRequestLogger(ctx, req, true)
	rw := l.ResponseWriter(httptest.NewRecorder())

	_, _ = rw.Write([]byte("hello"))
	l.FlushLog()
	require.Empty(t, hook.Entries)
}
