# io Package

The `io` package in Go provides basic interfaces and functions for `I/O` operations, such as reading and writing data.

It is often used for working with `streams`, including `files`, network connections, and `in-memory buffers`.

## Key Types and Functions in the io Package

### Interfaces

- `Reader`: An interface that wraps the basic Read method.
- `Writer`: An interface that wraps the basic Write method.
- `Closer`: An interface that wraps the Close method.

### Common Functions

- `Copy`: `io.Copy(dst, src)` copies data from a source to a destination.
- `ReadAll`: `io.ReadAll(r)` reads from a Reader until EOF and returns the data as a byte slice.
