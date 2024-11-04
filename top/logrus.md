# Logrus

`Logrus` is a structured logger for Go (golang) that is very popular for logging in applications.

It supports various `log levels`, `structured logging`, and `output formats`, making it suitable for both development and production environments.

```go
package main

import (
    "github.com/sirupsen/logrus"
)

func main() {
    // Create a new logger instance
    logger := logrus.New()

    // Set the log level (optional, default is InfoLevel)
    logger.SetLevel(logrus.InfoLevel)

    // Log messages of different levels
    logger.Info("This is an info message")
    logger.Warn("This is a warning message")
    logger.Error("This is an error message")

    // Log a structured message
    logger.WithFields(logrus.Fields{
        "username": "john_doe",
        "attempt":  1,
    }).Info("User login attempt")
}
```
