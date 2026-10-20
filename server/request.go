package server

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	s "strings"
)

type Request struct {
	method string
	path   string
	header map[string]string
	body   []byte
}

func Parse(r *bufio.Reader) (*Request, error) {
	startLine, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	startLine = s.TrimSpace(startLine)

	parts := s.Split(startLine, " ")

	if len(parts) < 2 {
		return nil, fmt.Errorf("malformed request line %q", startLine)
	}

	method, path := parts[0], parts[1]

	header := make(map[string]string)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = s.TrimSpace(line)
		if line == "" {
			break
		}

		k, v, ok := s.Cut(line, ":")

		if !ok {
			return nil, fmt.Errorf("malformed header: %q", line)
		}

		header[s.ToLower(s.TrimSpace(k))] = s.TrimSpace(v)
	}

	var body []byte
	if contentLengthStr, ok := header["content-length"]; ok {
		contentLength, err := strconv.Atoi(contentLengthStr)
		if err != nil {
			return nil, fmt.Errorf("invalid Content-Length: %w", err)
		}
		if contentLength > 0 {
			body = make([]byte, contentLength)

			_, err = io.ReadFull(r, body)

			if err != nil {
				return nil, fmt.Errorf("reading body: %w", err)
			}
		}
	}

	return &Request{
		method: method,
		path:   path,
		header: header,
		body:   body,
	}, nil
}
