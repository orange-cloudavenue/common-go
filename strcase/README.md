# strcase

This a fork of the scaleway `strcase` package. Many thanks for the original authors for their work.

The `strcase` package provides utilities for converting strings between different naming conventions commonly used in Go and other programming languages. It is especially useful for code generation, struct field manipulation, and working with APIs that require specific case formats.

## Features

- **CamelCase**: Convert strings to camelCase (e.g., `myVariableName`).
- **PascalCase**: Convert strings to PascalCase (e.g., `MyVariableName`).
- **Snake case**: Convert strings to snake_case (e.g., `my_variable_name`).
- **Kebab case**: Convert strings to kebab-case (e.g., `my-variable-name`).
- **UPPER CASE**: Convert strings to UPPER_CASE (e.g., `MY_VARIABLE_NAME`).
- **Go Name**: Convert strings to valid Go exported or unexported identifiers.
- **Number handling**: Properly handles numbers and mixed-case input.
- **Bash argument escaping**: Utilities for safely formatting strings as bash arguments.

## Usage

```go
import "github.com/orange-cloudavenue/common-go/strcase"

fmt.Println(strcase.ToCamel("my_variable_name"))      // myVariableName
fmt.Println(strcase.ToPascal("my_variable_name"))     // MyVariableName
fmt.Println(strcase.ToSnake("MyVariableName"))        // my_variable_name
fmt.Println(strcase.ToKebab("MyVariableName"))        // my-variable-name
fmt.Println(strcase.ToUpper("myVariableName"))        // MY_VARIABLE_NAME
fmt.Println(strcase.ToPublicGoName("my_field"))       // MyField
fmt.Println(strcase.ToPrivateGoName("my_field"))      // myField
```

## Advanced

- The package includes helpers for working with numbers in identifiers.
- Bash argument helpers are available for safe shell scripting.
