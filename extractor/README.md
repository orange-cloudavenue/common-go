# Extractor Package

The `extractor` package provides utility functions for extracting UUIDs and URNs from strings using regular expressions.  
It is designed for use in CloudAvenue SDKs and tools, but can be used in any Go project that needs to reliably extract these identifiers from arbitrary text.

## Features

- **ExtractUUID**:  
  Extracts a single UUID (version 4) from a string.  
  Returns an error if no UUID is found or if multiple UUIDs are present.

- **ExtractURN**:  
  Extracts a single URN containing a UUID4 from a string.  
  The URN must follow the format: `urn:<namespace>:<name>:<uuid4>`.  
  Returns an error if no URN is found or if multiple URNs are present.

## Usage

```go
import "github.com/orange-cloudavenue/common-go/extractor"

uuid, err := extractor.ExtractUUID("prefix d3c42a20-96b9-4452-91dd-f71b71dfe314 suffix")
// uuid == "d3c42a20-96b9-4452-91dd-f71b71dfe314"

urn, err := extractor.ExtractURN("http://foo.bar/urn:vcloud:gateway:d3c42a20-96b9-4452-91dd-f71b71dfe314")
// urn == "urn:vcloud:gateway:d3c42a20-96b9-4452-91dd-f71b71dfe314"
```

## Error Handling

Both functions return an error if:

- No match is found in the input string.
- More than one match is found (to avoid ambiguity).

## Example

```go
result, err := extractor.ExtractUUID("some text with uuid d3c42a20-96b9-4452-91dd-f71b71dfe314")
if err != nil {
    // handle error
}
fmt.Println("UUID:", result)
```

## License

This package is distributed under the MPL-2.0 license. See the [LICENSE](../LICENSE) file for details.
