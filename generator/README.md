# CloudAvenue Common Generator Library

This library provides custom data generators for CloudAvenue SDK testing and development, built on top of [gofakeit](https://github.com/brianvoe/gofakeit).  
It allows you to easily generate realistic test data for CloudAvenue objects, names, URNs, and more.

## Custom Generators

Below is a summary of all custom generators provided by this library, grouped by category:

### CloudAvenue Resource Generators

| Name               | Description                                            | Example                                   |
|--------------------|--------------------------------------------------------|-------------------------------------------|
| `edgegateway_name` | Generate a new EdgeGateway name matching conventions   | `tn01e02ocb0001234spt101`                 |
| `t0_name`          | Generate a new T0 router name matching conventions     | `pr01e02ocb0001234spt101`                 |
| `urn`              | Generate a new URN with a custom type (see params)     | `urn:vcloud:gateway:12345678-1234-4234-1234-123456789012` |

### String Format Generators

| Name         | Description                          | Example                 |
|--------------|--------------------------------------|-------------------------|
| `camelCase`  | Generate a random camelCase string   | `camelCaseExample`      |
| `PascalCase` | Generate a random PascalCase string  | `PascalCaseExample`     |
| `snake_case` | Generate a random snake_case string  | `snake_case_example`    |
| `UPPER_CASE` | Generate a random UPPER_CASE string  | `UPPER_CASE_EXAMPLE`    |
| `kebab-case` | Generate a random kebab-case string  | `kebab-case-example`    |

### Other Generators

| Name         | Description                          | Example                 |
|--------------|--------------------------------------|-------------------------|
| `href_uuid`        | Generate a random URL ending with a UUID               | `https://example.com/590c1440-9888-45b0-bd51-a817ee07c3f2` |

### URN Generator Parameters

| Name   | Description                                      | Example                        |
|--------|--------------------------------------------------|--------------------------------|
| `urn`  | Accepts parameter `urnType` (e.g., `vm`, `gateway`) to specify the URN type | `urn:vcloud:gateway:...`       |

For a complete list of available URN types, see the documentation here:  
[https://pkg.go.dev/github.com/orange-cloudavenue/common-go/urn#pkg-variables](https://pkg.go.dev/github.com/orange-cloudavenue/common-go/urn#pkg-variables)

## Usage

You can use these generators with the `Struct`, `Template`, `Generate`, and `Regex` functions re-exported from gofakeit:

```go
import "github.com/orange-cloudavenue/common-go/generator"

type MyStruct struct {
    Name string `fake:"{edgegateway_name}"`
    URN  string `fake:"{urn:edgeGateway}"`
    HREF string `fake:"{href_uuid}"`
}

generator.Struct(&myStruct)
```

Resulting Data:

```go
myStruct := MyStruct{
    Name: "tn01e02ocb0001234spt101",
    URN:  "urn:vcloud:gateway:12345678-1234-4234-1234-123456789012",
    HREF: "https://example.com/590c1440-9888-45b0-bd51-a817ee07c3f2",
}
```

## Other Generators

For all other built-in generators (names, addresses, numbers, etc.), refer to the official [gofakeit documentation](https://github.com/brianvoe/gofakeit).
