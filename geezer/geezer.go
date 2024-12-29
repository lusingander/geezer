package geezer

import (
	"bufio"
	"io"
	"slices"
	"strings"
	"unicode"
)

var (
	openingBrackets = []rune{'(', '[', '{'}
	closingBrackets = []rune{')', ']', '}'}
)

type indenter struct {
	n      int
	indent string
	w      int
}

func newIndenter(w int) *indenter {
	return &indenter{
		w: w,
	}
}

func (i indenter) get() string {
	return i.indent
}

func (i *indenter) inc() {
	i.n++
	i.update()
}

func (i *indenter) dec() {
	i.n--
	i.update()
}

func (i *indenter) update() {
	i.indent = strings.Repeat(" ", i.n*i.w)
}

func Exec(r io.Reader, w io.Writer, indentWidth int, withSpaceRunes []rune) error {
	br := bufio.NewReader(r)
	bw := bufio.NewWriter(w)

	ind := newIndenter(indentWidth)

	wsBuf := make([]rune, 0)
	skipWs := false

	for {
		r, _, err := br.ReadRune()

		if err == io.EOF {
			return bw.Flush()
		}
		if err != nil {
			return err
		}

		if unicode.IsSpace(r) {
			wsBuf = append(wsBuf, r)
			continue
		}

		if slices.Contains(closingBrackets, r) {
			ind.dec()
			bw.WriteRune('\n')
			bw.WriteString(ind.get())
			skipWs = true
		}

		if !skipWs {
			for _, ws := range wsBuf {
				bw.WriteRune(ws)
			}
		}
		wsBuf = make([]rune, 0)
		skipWs = false

		if slices.Contains(withSpaceRunes, r) {
			bw.WriteRune(' ')
			bw.WriteRune(r)
			bw.WriteRune(' ')
		} else {
			bw.WriteRune(r)
		}

		if r == ',' {
			bw.WriteRune('\n')
			bw.WriteString(ind.get())
			skipWs = true
		}

		if slices.Contains(openingBrackets, r) {
			ind.inc()
			bw.WriteRune('\n')
			bw.WriteString(ind.get())
			skipWs = true
		}
	}
}
