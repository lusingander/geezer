package geezer

import (
	"bytes"
	"strings"
	"testing"
)

func TestExec(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "abcdefg",
			want: "abcdefg",
		},
		{
			s: "Foo(value=1)",
			want: `Foo(
  value=1
)`,
		},
		{
			s: "Foo(value=1,name=foobar)",
			want: `Foo(
  value=1,
  name=foobar
)`,
		},
		{
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
		got := runExec(t, test.s)
		if got != test.want {
			t.Errorf("s=%v, got=%v, want=%v", test.s, got, test.want)
		}
	}
}

func runExec(t *testing.T, s string) string {
	r := strings.NewReader(s)
	w := &bytes.Buffer{}
	err := Exec(r, w)
	if err != nil {
		t.Error(err)
	}
	return w.String()
}
