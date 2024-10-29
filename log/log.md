# Key Features of the log Package

1. Basic Logging Functions

- `Print`, `Printf`, `Println`: Output log messages to the standard logger.
- `Fatal`, `Fatalf`, `Fatalln`: Log messages and then call os.Exit(1).
- `Panic`, `Panicf`, `Panicln`: Log messages and then call panic().

2. Custom Loggers

You can create custom loggers to direct log output to different writers, such as `files` or network connections.

3. Log Flags

You can customize the log output `format`, such as including `timestamps`, `file names`, and `line numbers`.
