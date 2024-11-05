# Viper

Viper is a popular configuration management package in Go.
It is used for handling `configuration files`, environment `variables`, `flags`, and more, in a way that makes it easy to load and access configuration data in your Go applications.

With Viper, you can centralize and manage configuration in various formats like `JSON`, `TOML`, `YAML`, `HCL`, or even environment variables. It makes it simpler to read `configuration values dynamically`, and it supports `defaults`, `environment overrides`, and more.

## 1. Install Viper

You can install Viper using go get:

```sh
go get github.com/spf13/viper
```

## 2. Basic Usage Example

Let’s start by creating a simple Go application that reads a configuration from a file.

**Folder Structure:**

```lua
/project-root
    /config
        config.yaml
    main.go
    go.mod

```

```yaml
# config.yaml
server:
  host: "localhost"
  port: 8080

database:
  user: "root"
  password: "password"
  host: "localhost"
  port: 3306
```
