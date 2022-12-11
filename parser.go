package httparser

import (
	"errors"
	"strings"
	"unsafe"
)

type Request struct {
	Line    Line
	Headers map[string]string
	Body    string
}

type Line struct {
	Method, Path string
	Proto        string
}

// Parse parses the line of HTTP request and returns a Line
func ParseLine(input []byte) (line Line, err error) {
	http := *(*string)(unsafe.Pointer(&input))
	l := len(http)
	pass := 0
	ok := false

	for i := pass; i < l; i++ {
		if http[i] == ' ' {
			line.Method = http[:i]
			pass = i + 1
			ok = true
			break
		}
	}

	if !ok {
		return Line{}, errors.New("missing method")
	}

	ok = false

	for i := pass; i < l; i++ {
		if http[i] == ' ' {
			line.Path = http[pass:i]
			pass = i + 1
			ok = true
			break
		}
	}

	if !ok {
		return Line{}, errors.New("missing path")
	}

	ok = false

	for i := pass; i < l; i++ {
		if http[i] == '\r' {
			line.Proto = http[pass:i]
			pass = i + 2
			ok = true
			break
		}
	}

	if !ok {
		return Line{}, errors.New("missing protocol")
	}

	return line, nil
}

// Parse parses the HTTP request and returns a Request
func Parse(input []byte) (req Request, err error) {
	http := *(*string)(unsafe.Pointer(&input))
	l := len(http)
	pass := 0
	ok := false

	for i := pass; i < l; i++ {
		if http[i] == ' ' {
			req.Line.Method = http[:i]
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
			req.Line.Path = http[pass:i]
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
			req.Line.Proto = http[pass:i]
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
