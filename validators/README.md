# Validators Package

This package provides a set of custom validators for Go applications, built on top of the `go-playground/validator` library. These validators are designed to handle specific validation scenarios, such as validating IP ranges, TCP/UDP ports, HTTP status codes, URNs, key-value pairs, string case formats, and more.

## Features

### Network Validators

| Name               | Description                                                        | Parameters | Example                        |
|--------------------|--------------------------------------------------------------------|------------|--------------------------------|
| `ipv4_range`       | Validates if a string represents a valid IPv4 range                |     ➖       | `192.168.0.1-192.168.0.100`    |
| `tcp_udp_port`     | Validates if a value is a valid TCP or UDP port                    |      ➖      | `80`, `443`                    |
| `tcp_udp_port_range` | Validates if a string represents a valid range of TCP/UDP ports   |     ➖       | `8000-8080`                    |

### HTTP Validators

| Name                 | Description                                                      | Parameters | Example                        |
|----------------------|------------------------------------------------------------------|------------|--------------------------------|
| `http_status_code`   | Validates if a value is a valid HTTP status code                 |      ➖      | `200`, `404`                   |
| `http_status_code_range` | Validates if a string represents a valid range of HTTP status codes |     ➖       | `200-299`                |

### CloudAvenue Resource Validators

| Name               | Description                                                        | Parameters | Example                        |
|--------------------|--------------------------------------------------------------------|------------|--------------------------------|
| `urn=typeOfURN`    | Validates if a string is a valid URN. For a complete list of available URN types, see the documentation here: [https://pkg.go.dev/github.com/orange-cloudavenue/common-go/urn#pkg-variables](https://pkg.go.dev/github.com/orange-cloudavenue/common-go/urn#pkg-variables) | `typeOfURN` | `urn:vcloud:gateway:...`       |
| `resource_name=resourceKey` | Validates if a string is a valid CAV resource name for the given resource key | `resourceKey` | `tn01e02ocb0001234spt101` (for `edgegateway`), `prvrf01eocb0001234allsp01` (for `t0_name`) For a complete list of resource keys, see the documentation here: [https://pkg.go.dev/github.com/orange-cloudavenue/common-go/regex#pkg-variables](https://pkg.go.dev/github.com/orange-cloudavenue/common-go/regex#pkg-variables) |

### Key/Value Validators

| Name           | Description                                        | Parameters | Example                |
|----------------|----------------------------------------------------|------------|------------------------|
| `str_key_value`| Validates key-value pairs                          |      ➖      | `key=value`            |

### String Format Validators

#### Format Validators

Le validateur `format` permet de valider des formats de chaînes spécifiques. Actuellement, le format supporté est :

| Name           | Description                                         | Parameters | Example              |
|----------------|-----------------------------------------------------|------------|----------------------|
| `format=email` | Valide qu'une chaîne est une adresse email valide   |     ➖      | `user@example.com`   |

**Exemple d'utilisation :**

```go
type User struct {
    Email string `validate:"format=email"`
}

func main() {
    user := User{Email: "user@example.com"}
    err := validators.New().Struct(&user)
    if err != nil {
        panic(err)
    }
    fmt.Println("Email valide !")
}
```

#### Case Validators

| Name             | Description                                        | Parameters | Example                |
|------------------|----------------------------------------------------|------------|------------------------|
| `disallow_upper` | Ensures a string does not contain uppercase letters|     ➖       | `test`                 |
| `disallow_space` | Ensures a string does not contain spaces           |     ➖       | `testtest`             |
| `case`           | Checks if the field is in a specific case (see below)| `camelCase`, `snake_case`, `PascalCase`, `UPPER_CASE`, `kebab-case` | `camelCaseExample`   |

### Conditional Validators

| Name                | Description                                        | Parameters | Example                |
|---------------------|----------------------------------------------------|------------|------------------------|
| `required_if_null`  | Requires a field if another field is null          |     `fieldName`       |      `required_if_null=ID`                  |
| `excluded_if_null`  | Excludes a field if another field is null          |       `fieldName`     |      `excluded_if_null=ID`                  |

### Default Value Setter

| Name     | Description                                        | Example                |
|----------|----------------------------------------------------|------------------------|
| `default`| Sets the field to the default value for its type if it is empty | `0` for `int`, `""` for `string` |

Example usage of the `default` setter:

```go
type Example struct {
    Enabled bool `default:"true"`
    Count   int  `default:"10"`
}

func main() {
    example := Example{}
    if err := validators.New().Struct(&example); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Enabled: %v, Count: %d\n", example.Enabled, example.Count)
}
```

## Easy Integration

The package provides a `New()` function to create and configure a validator instance with all custom validations pre-registered.

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

    "github.com/orange-cloudavenue/common-go/validators"
)

type Example struct {
    IPRange string `validate:"ipv4_range"`
    Port    int    `validate:"tcp_udp_port"`
    Enabled bool   `default:"true"`
}

func main() {
    example := Example{
        IPRange: "192.168.0.1-192.168.0.100",
        Port:    8080,
    }

    err := validators.New().Struct(&example)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Enabled: %v, IP Range: %s, Port: %d\n", example.Enabled, example.IPRange, example.Port)
}
```

[Play example here](https://goplay.tools/snippet/YSwC4jAQelU).

## Contributing

We welcome contributions to this project! If you find a bug or have a feature request, please open an issue on GitHub. Pull requests are also welcome. Please ensure that your code adheres to the existing style and includes appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](../LICENSE) file for details.
