# GO-INLINE-HTML!

Inlines HTML files into Go variables.

Replaces the content of []byte variables in Go files with their HTML counterpart, based on a filename-to-variable-name relationship. Support multilines, does not compress HTML.

## Usage.
```
$ go run . -help

Usage of go-embedd-html:
  -dest string
        destination Go file that contains variables to be filled with HTML
  -source string
        folder with HTML files, named as [variable].html

$ go run . -dest=./example/main.txt -source=./example

0
```

## Before.
```
package main

import "fmt"

var (
    example1 = []byte{}
    example2 = []byte{}
)

func main() {
    fmt.Println(string(example1))
    fmt.Println(string(example2))
}
```
## After.
```
package main

import "fmt"

var (
    example1 = []byte(`<html>

<body>
    <h1>Hello World!</h1>
</body>

</html>`)
    example2 = []byte(`<html>

<body>
    <h1>Goodbye World!</h1>
</body>

</html>`)
)

func main() {
    fmt.Println(string(example1))
    fmt.Println(string(example2))
}
```