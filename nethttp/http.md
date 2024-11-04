```go
import (
    "fmt"
    "net/http"
)
```

- `fmt`: This package is used for formatted `I/O` operations, such as printing to the console.
- `net/http`: This package provides HTTP `client` and `server` functionalities.

**Handler Function:**

```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}
```

### Parameters:

- `w http.ResponseWriter`: This interface is used to construct an HTTP response. It allows you to write data back to the client.
- `r *http.Request`: This pointer represents the incoming HTTP request. It contains information about the `request`, such as headers, method, and body.

### Functionality:

The fmt.Fprintf function writes "Hello, World!" to the response writer, which sends it back to the `client`.

**Main Function:**

```go
func main() {
    http.HandleFunc("/", helloHandler)
```

- `http.HandleFunc("/", helloHandler)`: This registers the helloHandler function to handle requests to the root URL (/). When a request is made to this path, helloHandler will be invoked.
