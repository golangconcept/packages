# fmt Package

The `fmt package` in Go provides formatted `I/O` functions, similar to C's `printf` and `scanf`. It's widely used for formatting `strings`, `reading input`, and `printing output`.

## Key Functions in the fmt Package

### Printing Functions

- `Print`: Prints to standard output.
- `Printf`: Formats and prints according to a format specifier.
- `Println`: Prints to standard output with a newline.

### Scanning Functions

- `Scan`: Reads from standard input.
- `Scanf`: Reads formatted input from standard input.
- `Scanln`: Reads from standard input until a newline.

### String Formatting

- `Sprintf`: Formats and returns a string.
- `Sscanf`: Reads formatted input from a string.

## Formatting Strings

You can format strings using various verbs:

- `%s`: string
- `%d`: integer
- `%f`: float
- `%v`: default format

```go
package main

import (
    "fmt"
)

func main() {
    var name string
    var age int

    fmt.Print("Enter your name: ")
    fmt.Scan(&name) // Reads a string

    fmt.Print("Enter your age: ")
    fmt.Scanf("%d", &age) // Reads an integer

    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}
```
