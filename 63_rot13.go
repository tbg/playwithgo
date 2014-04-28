package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(p []byte) (int, error) {
	var n int
	var err error
	n, err = r13.r.Read(p)
	for i, v := range p {
		p[i] = (v + 13)
		switch {
		case v <= 'L':
			p[i] = v + 13
		case v < 97:
			p[i] = v - 13
		case v < 108:
			p[i] = v + 13
		case true:
			p[i] = v - 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
