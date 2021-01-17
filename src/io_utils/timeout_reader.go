package io_utils

import (
	"errors"
	"io"
	"time"
)

type TimeoutReader struct {
	reader  io.ReadCloser
	timeout time.Duration
}

func NewTimeoutReader(reader io.ReadCloser, timeout time.Duration) io.ReadCloser {
	ret := new(TimeoutReader)
	ret.reader = reader
	ret.timeout = timeout
	return ret
}

func (self *TimeoutReader) Read(buf []byte) (n int, err error) {
	ch := make(chan bool)
	n = 0
	err = nil
	go func() {
		n, err = self.reader.Read(buf)
		ch <- true
	}()
	select {
	case <-ch:
		return
	case <-time.After(self.timeout):
		return 0, errors.New("Timeout expired")
	}
}

func (self *TimeoutReader) Close() error {
	return self.reader.Close() 
} 