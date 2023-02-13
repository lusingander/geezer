# geezer

## Usage

```sh
$ echo "Foo(bar=Bar(baz={n=1,m=2},qux=Qux(name=qqq,value=[1,2,3])))" | geezer
Foo(
  bar = Bar(
    baz = {
      n = 1,
      m = 2
    },
    qux = Qux(
      name = qqq,
      value = [
        1,
        2,
        3
      ]
    )
  )
)
```

### Options

```
  -n int
        indent width (default 2)
  -s string
        characters with spaces before and after (default "=")
```

For exapmle:

```
$ echo "foo:{bar:{n:1,m:2}}" | go run main.go -n 1 -s ":" 
foo : {
 bar : {
  n : 1,
  m : 2
 }
}
```

## License

MIT
