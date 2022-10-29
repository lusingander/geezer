package geezer

import (
	"bytes"
	"strings"
	"testing"
)

func TestExec(t *testing.T) {
	tests := []struct {
		w    int
		s    string
		want string
	}{
		{
			w:    2,
			s:    "abcdefg",
			want: "abcdefg",
		},
		{
			w: 2,
			s: "Foo(value=1)",
			want: `Foo(
  value=1
)`,
		},
		{
			w: 8,
			s: "Foo(value=1)",
			want: `Foo(
        value=1
)`,
		},
		{
			w: 2,
			s: "Foo(value=1,name=foobar)",
			want: `Foo(
  value=1,
  name=foobar
)`,
		},
		{
			w: 2,
			s: "Foo(value=1,name=foobar,bar=Bar(id=1))",
			want: `Foo(
  value=1,
  name=foobar,
  bar=Bar(
    id=1
  )
)`,
		},
		{
			w: 2,
			s: "Foo(bar=Bar(baz=Baz{n=1,m=2},qux=Qux(name=qqq,value=[1,2,3])))",
			want: `Foo(
  bar=Bar(
    baz=Baz{
      n=1,
      m=2
    },
    qux=Qux(
      name=qqq,
      value=[
        1,
        2,
        3
      ]
    )
  )
)`,
		},
	}
	for _, test := range tests {
		got := runExec(t, test.s, test.w)
		if got != test.want {
			t.Errorf("s=%v, got=%v, want=%v", test.s, got, test.want)
		}
	}
}

func runExec(t *testing.T, s string, indentWidth int) string {
	r := strings.NewReader(s)
	w := &bytes.Buffer{}
	err := Exec(r, w, indentWidth)
	if err != nil {
		t.Error(err)
	}
	return w.String()
}
