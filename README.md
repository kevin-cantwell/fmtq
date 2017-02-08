The `fmtq` command tool simply parses stdin and returns a quoted string safely escaped with Go syntax. It's essentially a command-line equivalent to the print verb `%q` from Go's [fmt](https://golang.org/pkg/fmt/) package.

## Installation

`go get -u github.com/kevin-cantwell/fmtq`

## Usage

```
$ echo "foo
> bar
> baz" | fmtq
"foo\nbar\nbaz\n"
```

You can choose to return a double quoted string (default) or a single-quoted string (using the `-s` flag):

```
$ echo -n "This is a single quote ' and this is a double quote \"" | fmtq
"This is a single quote ' and this is a double quote \""
$ echo -n "This is a single quote ' and this is a double quote \"" | fmtq -s
'This is a single quote \' and this is a double quote "'
```