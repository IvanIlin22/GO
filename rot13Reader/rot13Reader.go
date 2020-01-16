package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (p *rot13Reader) Read(b []byte) (int, error){
	count, err := p.r.Read(b)
	for i := 0; i < len(b); i++ {
		if (b[i] >= 'a' && b[i] <= 'm') || (b[i] >= 'A' && b[i] <= 'M') {
			b[i] += 13 
		} else if (b[i] >= 'n' && b[i] <= 'z') || (b[i] >= 'N' && b[i] <= 'Z') {
			b[i] -= 13
		}
	}
	return count, err	
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}