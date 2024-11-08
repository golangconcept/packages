# os Package

The os package provides a platform-independent interface to operating system functionality, allowing you to work with `files`, `directories`, `environment` variables, and `process management`.

## Common Tasks with the `os` Package

### 1 Working with Files and Directories

- Creating a File
- Writing to a File
- Reading from a File
- Listing Files in a Directory

```go
func main() {
 files, err := ioutil.ReadDir(".") // Read current directory
 if err != nil {
     log.Fatal(err)
 }
 for _, file := range files {
     log.Println(file.Name())
 }
}
```

### 2 Working with Environment Variables

- Getting an Environment Variable
- Setting an Environment Variable

### 3 Process Management

- Exiting a Program
- Command-line Arguments

```

```
