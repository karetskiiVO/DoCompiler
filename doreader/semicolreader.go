package doreader

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"slices"
)

type SemicolReader struct {
	scanner *bufio.Scanner
	buf     []byte
}

var _ io.Reader = &SemicolReader{}

func NewSemicolReader(reader io.Reader) *SemicolReader {
	return &SemicolReader{
		scanner: bufio.NewScanner(NewCommentDeletor(reader)),
	}
}

var empty = map[rune]struct{}{
	'\n': {},
	'\r': {},
	' ':  {},
	'\t': {},
}

var skip = map[rune]struct{}{
	'{': {},
	'}': {},
	'(': {},
	'+': {},
	'-': {},
	'*': {},
	'/': {},
	'=': {},
	',': {},
	'.': {},
	'&': {},
	'|': {},
}

func (sr *SemicolReader) Read(p []byte) (int, error) {
	if sr.buf == nil {
		if !sr.scanner.Scan() {
			err := sr.scanner.Err()
			if err != nil {
				return 0, err
			} else {
				return 0, io.EOF
			}
		}

		buf := append(sr.scanner.Bytes(), '\n')
		idx := bytes.LastIndexFunc(buf, func(r rune) bool {
			_, ok := empty[r]
			return !ok
		})

		if idx != -1 {
			if _, ok := skip[rune(buf[idx])]; !ok {
				sr.buf = slices.Insert(buf, idx+1, ';')
			} else {
				fmt.Printf("except([%v] = '%c'(%v)): %v\n", idx, rune(buf[idx]), rune(buf[idx]), string(buf))
			}
		}

		if sr.buf == nil {
			sr.buf = buf
		}
	}

	if sr.buf != nil {
		n := copy(p, sr.buf)
		sr.buf = sr.buf[n:]

		if len(sr.buf) == 0 {
			sr.buf = nil
		}
		if len(p) > n {
			m, err := sr.Read(p[n:])
			return n + m, err
		}
		return n, nil
	}

	return 0, io.EOF
}
