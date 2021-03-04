package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = reader.r.Read(b)

	for i, data := range b {
		if data >= 'a' && data <= 'm' || data >= 'A' && data <= 'M' {
			b[i] += 13
			continue
		}
		if data >= 'n' && data <= 'z' || data >= 'N' && data <= 'Z' {
			b[i] -= 13
			continue
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
