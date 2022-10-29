# geezer

## Usage

```sh
$ echo "Foo(bar=Bar(baz=Baz{n=1,m=2},qux=Qux(name=qqq,value=[1,2,3])))" | geezer
Foo(
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
)
```

### Options

```
  -n int
        indent width (default 2)
```

## License

MIT
