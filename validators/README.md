# Validators Package

This package provides a set of custom validators for Go applications, built on top of the `go-playground/validator` library. These validators are designed to handle specific validation scenarios, such as validating IP ranges, TCP/UDP ports, HTTP status codes, and more.

## Features

- **Custom Validators**:
  - `ipv4_range`: Validates if a string represents a valid IPv4 range (e.g., `192.168.0.1-192.168.0.100`).
  - `tcp_udp_port`: Validates if a value is a valid TCP or UDP port. (e.g., `80`, `443`).
  - `tcp_udp_port_range`: Validates if a string represents a valid range of TCP/UDP ports. (e.g., `8000-8080`).
  - `http_status_code`: Validates if a value is a valid HTTP status code. (e.g., `200`, `404`).
  - `http_status_code_range`: Validates if a string represents a valid range of HTTP status codes. (e.g., `200-299`).
  - `urn`: Validates if a string is a valid URN.
  - `key_value`: Validates key-value pairs. (e.g., `key=value`).
  - `disallow_upper`: Ensures a string does not contain uppercase letters.
  - `disallow_space`: Ensures a string does not contain spaces.
  - `default`: Validates if a value is the default value for its type (e.g., `0` for `int`, `""` for `string`).

- **Easy Integration**: The package provides a `New()` function to create and configure a validator instance with all custom validations pre-registered.

## Installation

To use this package, add it to your project:

```bash
go get github.com/orange-cloudavenue/common-go/validators
```

## Usage

Here is an example of how to use the validators in your Go application:

```go
package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
    "github.com/orange-cloudavenue/common-go/validators"
)

type Example struct {
    IPRange string `validate:"ipv4_range"`
    Port    int    `validate:"tcp_udp_port"`
}

func main() {
    v := validators.New()

    example := Example{
        IPRange: "192.168.0.1-192.168.0.100",
        Port:    8080,
    }

    err := v.Struct(example)
    if err != nil {
        if _, ok := err.(*validator.InvalidValidationError); ok {
            fmt.Println("Validation error:", err)
        } else {
            for _, err := range err.(validator.ValidationErrors) {
                fmt.Printf("Field '%s' failed validation with tag '%s'\n", err.Field(), err.Tag())
            }
        }
    } else {
        fmt.Println("Validation passed!")
    }
}
```

## Contributing

We welcome contributions to this project! If you find a bug or have a feature request, please open an issue on GitHub. Pull requests are also welcome. Please ensure that your code adheres to the existing style and includes appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](../LICENSE) file for details.
