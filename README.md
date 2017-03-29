The `fmtq` command tool simply parses stdin and returns a double quoted string safely escaped with Go syntax. It's essentially a command-line equivalent to the print verb `%q` from Go's [fmt](https://golang.org/pkg/fmt/) package.

## Installation

`go get -u github.com/kevin-cantwell/fmtq`

## Usage

```
$ echo "foo
> \"bar\"
> baz" | fmtq
"foo\n\"bar\"\nbaz\n"
```
