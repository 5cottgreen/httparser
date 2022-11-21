package httparser

import (
	"errors"
	"strings"
	"unsafe"
)

type Request struct {
	Method, Path string
	Proto        string
	Headers      map[string]string
	Body         string
}

// Parses the buffer as an HTTP request and returns data as a Request
func Parse(input []byte) (req Request, err error) {
	http := *(*string)(unsafe.Pointer(&input))
	l := len(http)
	pass := 0
	ok := false

	for i := pass; i < l; i++ {
		if http[i] == ' ' {
			req.Method = http[:i]
			pass = i + 1
			ok = true
			break
		}
	}

	if !ok {
		return Request{}, errors.New("missing method")
	}

	ok = false

	for i := pass; i < l; i++ {
		if http[i] == ' ' {
			req.Path = http[pass:i]
			pass = i + 1
			ok = true
			break
		}
	}

	if !ok {
		return Request{}, errors.New("missing path")
	}

	ok = false

	for i := pass; i < l; i++ {
		if http[i] == '\r' {
			req.Proto = http[pass:i]
			pass = i + 2
			ok = true
			break
		}
	}

	if !ok {
		return Request{}, errors.New("missing protocol")
	}

	ok = false

	req.Headers = make(map[string]string)

	for i := pass; i < l; i++ {
		if http[i] == '\r' && http[i+1] == '\n' {
			if http[i+2] == '\r' && http[i+3] == '\n' {
				req.Body = http[i+4:]
			} else {
				head := http[pass : i-2]
				pass = i + 2

				col := strings.IndexByte(head, ':')
				if col != -1 {
					req.Headers[head[:col]] = head[col+2:]
				} else {
					return Request{}, errors.New("missing header")
				}
			}
		}
	}

	return req, nil
}
