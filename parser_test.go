package httparser

import (
	"testing"
)

func TestParse(t *testing.T) {
	method := "GET"
	path := "/main"
	proto := "HTTP/1.1"
	head := "Host"
	value := "host:port"
	body := "Hello"

	http := []byte(method + " " + path + " " + proto + "\r\n" + head + ":" + " " + value + "\r\n\r\n" + body)

	req, err := Parse(http)
	if err != nil {
		t.Errorf("the '%s' error was not expected", err.Error())
	} else {
		if req.Line.Method != method {
			t.Errorf("the '%s' method was expected", method)
		}

		if req.Line.Path != path {
			t.Errorf("the '%s' path was expected", path)
		}

		if req.Line.Proto != proto {
			t.Errorf("the '%s' protocol was expected", proto)
		}

		if v, ok := req.Headers[head]; !ok {
			t.Errorf("the '%s' header was expected", head)
		} else if v != value {
			t.Errorf("the '%s' header value was expected", value)
		}

		if req.Body != body {
			t.Errorf("the '%s' body was expected", body)
		}
	}
}

func TestLine(t *testing.T) {
	method := "GET"
	path := "/main"
	proto := "HTTP/1.1"

	http := []byte(method + " " + path + " " + proto + "\r\n")

	line, err := ParseLine(http)
	if err != nil {
		t.Errorf("the '%s' error was not expected", err.Error())
	} else {
		if line.Method != method {
			t.Errorf("the '%s' method was expected", method)
		}

		if line.Path != path {
			t.Errorf("the '%s' path was expected", path)
		}

		if line.Proto != proto {
			t.Errorf("the '%s' protocol was expected", proto)
		}

	}
}
