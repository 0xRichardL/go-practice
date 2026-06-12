package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Generator struct {
	From int
	To   int
	cur  int
}

func NewGenerator(from, to int) io.Reader {
	return &Generator{
		From: from,
		To:   to,
		cur:  from,
	}
}

func (g *Generator) Read(p []byte) (n int, err error) {
	if g.cur > g.To {
		return 0, io.EOF
	}
	for g.cur <= g.To {
		s := strconv.Itoa(g.cur) + "\n"
		if len(s) > len(p)-n {
			break
		}
		copy(p[n:], s)
		n += len(s)
		g.cur++
	}
	return n, nil
}

func double_transform(input io.Reader) io.Reader {
	r, w := io.Pipe()
	go func() {
		defer w.Close()
		scanner := bufio.NewScanner(input)
		// Scan 64KB (default max token size) at a time.
		for scanner.Scan() {
			// Read a line (no newline character included).
			line := scanner.Text()
			num, err := strconv.Atoi(line)
			if err != nil {
				w.CloseWithError(err)
				return
			}
			// Write the doubled number followed by a newline.
			fmt.Fprintln(w, num*2)
		}
		if scanner.Err() != nil {
			w.CloseWithError(scanner.Err())
		}
	}()
	return r
}

func main() {
	r := NewGenerator(1, 10)
	dr := double_transform(r)
	io.Copy(os.Stdout, dr)
}
