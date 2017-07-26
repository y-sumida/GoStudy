package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b[]byte) (n int, err error) {
	n, err = rot.r.Read(b)

	for i := 0; i < len(b); i++ {
		s := b[i]
		if (s >= 'A' && s <= 'M') || (s >= 'a' && s <= 'm') {
			s += 13
		} else if (s >= 'N' && s <= 'Z') || (s >= 'n' && s <= 'z') {
			s -= 13
		}
		b[i] = s
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
