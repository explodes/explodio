package stand

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Handler interface {
	ServeHTTP(writer http.ResponseWriter, request *http.Request, chain Chain) error
}

var _ Handler = (*HandlerFunc)(nil)

type HandlerFunc func(writer http.ResponseWriter, request *http.Request, chain Chain) error

func (f HandlerFunc) ServeHTTP(writer http.ResponseWriter, request *http.Request, chain Chain) error {
	return f(writer, request, chain)
}

type Chain interface {
	Proceed() error
}

func WrapMiddleware(handler Handler, middleware ...Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		chain := chain{
			writer:  writer,
			request: request,
			index:   0,
			funcs:   append(middleware, handler),
			err:     nil,
		}
		chain.Proceed()
	})
}

type chain struct {
	writer  http.ResponseWriter
	request *http.Request
	index   int
	funcs   []Handler
	err     error
}

func (c *chain) Proceed() error {
	if c.err != nil {
		return c.err
	}
	if c.index == len(c.funcs) {
		return nil
	}
	c.index++
	c.err = c.funcs[c.index-1].ServeHTTP(c.writer, c.request, c)
	return c.err
}

func ConvertHttpHandler(handler http.Handler) Handler {
	return HandlerFunc(func(writer http.ResponseWriter, request *http.Request, chain Chain) error {
		var err error
		capture := func(writer http.ResponseWriter, request *http.Request) {
			defer func() {
				result := recover()
				if result == nil {
					return
				} else if e, ok := result.(error); ok {
					err = e
				} else if s, ok := result.(fmt.Stringer); ok {
					err = errors.New(s.String())
				} else {
					err = fmt.Errorf("%v", result)
				}
			}()
			handler.ServeHTTP(writer, request)
		}
		capture(writer, request)
		if err == nil {
			return chain.Proceed()
		}
		return err
	})
}

func TimingMiddleware(log Loggerf) Handler {
	return HandlerFunc(func(writer http.ResponseWriter, request *http.Request, chain Chain) error {
		log.Infof(">>> %s", request.URL)
		start := time.Now()
		err := chain.Proceed()
		end := time.Now()
		log.Infof("<<< %s (%v)", request.URL, end.Sub(start))
		return err
	})
}

func ErrorLoggingMiddleware(log Loggerf) Handler {
	return HandlerFunc(func(writer http.ResponseWriter, request *http.Request, chain Chain) error {
		err := chain.Proceed()
		if err != nil {
			log.Errorf("%s: %v", request.URL, err)
		}
		return err
	})
}
