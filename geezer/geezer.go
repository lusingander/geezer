package geezer

import (
	"bufio"
	"io"
	"strings"
)

const indent = 2

type indenter struct {
	n int
}

func (i indenter) get() string {
	return strings.Repeat(" ", i.n*indent)
}

func Exec(r io.Reader, w io.Writer) error {
	br := bufio.NewReader(r)
	bw := bufio.NewWriter(w)

	ind := indenter{}

	for {
		r, _, err := br.ReadRune()

		if err == io.EOF {
			return bw.Flush()
		}
		if err != nil {
			return err
		}

		if r == ')' || r == ']' || r == '}' {
			ind.n--
			bw.WriteRune('\n')
			bw.WriteString(ind.get())
		}

		bw.WriteRune(r)

		if r == ',' {
			bw.WriteRune('\n')
			bw.WriteString(ind.get())
		}

		if r == '(' || r == '[' || r == '{' {
			ind.n++
			bw.WriteRune('\n')
			bw.WriteString(ind.get())
		}
	}
}
