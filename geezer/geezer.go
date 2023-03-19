package geezer

import (
	"bufio"
	"io"
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

	skipWs := false

	for {
		r, _, err := br.ReadRune()

		if skipWs {
			if unicode.IsSpace(r) {
				continue
			}
			skipWs = false
		}

		if err == io.EOF {
			return bw.Flush()
		}
		if err != nil {
			return err
		}

		if isIn(r, closingBrackets) {
			ind.dec()
			bw.WriteRune('\n')
			bw.WriteString(ind.get())
		}

		if isIn(r, withSpaceRunes) {
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

		if isIn(r, openingBrackets) {
			ind.inc()
			bw.WriteRune('\n')
			bw.WriteString(ind.get())
		}
	}
}

func isIn(r rune, rs []rune) bool {
	for _, rr := range rs {
		if r == rr {
			return true
		}
	}
	return false
}
