package stand

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"regexp"
	"testing"
)

type simpleChain struct{}

func (simpleChain) Proceed() error {
	return nil
}

var _ Loggerf = (*rememberLogs)(nil)

type log struct {
	format string
	args   []interface{}
}

func (l log) String() string {
	return fmt.Sprintf(l.format, l.args...)
}

type rememberLogs struct {
	debug []log
	info  []log
	warn  []log
	error []log
}

func (r *rememberLogs) Debugf(format string, args ...interface{}) {
	r.debug = append(r.debug, log{format, args})
}

func (r *rememberLogs) Infof(format string, args ...interface{}) {
	r.info = append(r.info, log{format, args})
}

func (r *rememberLogs) Warnf(format string, args ...interface{}) {
	r.warn = append(r.warn, log{format, args})
}

func (r *rememberLogs) Errorf(format string, args ...interface{}) {
	r.error = append(r.error, log{format, args})
}

func TestWrapMiddleware(t *testing.T) {
	handler := ConvertHttpHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
	count := 0
	middleware := HandlerFunc(func(writer http.ResponseWriter, request *http.Request, chain Chain) error {
		count++
		return chain.Proceed()
	})
	wrapped := WrapMiddleware(handler, middleware, middleware, middleware)
	wrapped.ServeHTTP(nil, nil)
	assert.Equal(t, 3, count)
}

func TestConvertHttpHandler(t *testing.T) {
	tcs := []struct {
		name        string
		panicArg    interface{}
		expectedMsg string
	}{
		{"string", "failed", "failed"},
		{"error", errors.New("failed"), "failed"},
		{"number", 1000, "1000"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			captured := ConvertHttpHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				panic(tc.panicArg)
			}))
			err := captured.ServeHTTP(nil, nil, simpleChain{})
			assert.EqualError(t, err, tc.expectedMsg)
		})
	}
}

func TestTimingMiddleware(t *testing.T) {
	logs := &rememberLogs{}
	wrapped := WrapMiddleware(HandlerFunc(func(writer http.ResponseWriter, request *http.Request, chain Chain) error {
		return chain.Proceed()
	}), TimingMiddleware(logs))
	requestUrl, parseErr := url.Parse("http://explod.io/foo")
	assert.NoError(t, parseErr)

	r := &http.Request{URL: requestUrl}
	wrapped.ServeHTTP(nil, r)

	assert.Equal(t, 2, len(logs.info))
	assert.Equal(t, ">>> http://explod.io/foo", logs.info[0].String())
	assert.Regexp(t, regexp.MustCompile("<<< http://explod.io/foo \\(\\d+\\w+\\)"), logs.info[1].String())
}

func TestErrorLoggingMiddleware(t *testing.T) {
	logs := &rememberLogs{}
	wrapped := WrapMiddleware(HandlerFunc(func(writer http.ResponseWriter, request *http.Request, chain Chain) error {
		return errors.New("expected")
	}), ErrorLoggingMiddleware(logs))
	requestUrl, parseErr := url.Parse("http://explod.io/foo")
	assert.NoError(t, parseErr)

	r := &http.Request{URL: requestUrl}
	wrapped.ServeHTTP(nil, r)

	assert.Equal(t, 1, len(logs.error))
	assert.Equal(t, "http://explod.io/foo: expected", logs.error[0].String())
}
