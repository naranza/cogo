# Cogo

**Cogo** (Config for Go) is a lightweight Go package for loading structured configuration from plain text files.  
It supports parsing config files written in a `key type value` format and mapping the values directly into a Go struct.

Example use-case: Loading application settings from `.cogo` files in a strongly-typed way.

## Features

- Easy-to-read config file syntax
- Supports `int`, `bool`, `float`, and `string` types
- Maps config entries to Go struct fields via reflection
- Helpful error reporting for invalid lines and types

## Config File Format

Each line must follow the format:

```
key type value

```

- `key` maps to a field name in the struct (case-insensitive)
- `type` must be one of: `int`, `bool`, `float`, `float64`, or `string`
- `value` is the value to assign (supports multi-word strings)

Lines starting with `#` are treated as comments.

Example `config.cogo`:

```
# Application settings

Port int 8080
Debug bool true
Rate float 0.75
Name string MyApp
```

## Usage

```go
package main

import (
  "fmt"
  "log"
  "gitlab.com/naranza/cogo"
)

type Config struct {
  Port  int
  Debug bool
  Rate  float64
  Name  string
}

func main() {
  var cfg Config
  err := cogo.LoadConfig("config.cogo", &cfg)
  if err != nil {
      log.Fatal(err)
  }

  fmt.Printf("Loaded config: %+v\n", cfg)
}
```

## Testing

To run unit tests, ensure you have the Go toolchain installed and run:

```bash
go test ./...
```

Your tests are in `gogo_test.go` and expect test config files to exist in a `files/` subdirectory.

## Project Structure

```
/cogo
  ├── cogo.go           // Core implementation
  ├── tests/            // Unit tests and config files
```

---

## License

AGPL-3.0
Copyright © 2025
Andrea Davanzo and contributors
