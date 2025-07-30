# Print package

## Description

`print` is a supercharged version of the [`tabwriter`](https://pkg.go.dev/text/tabwriter) package. It provides a `SetHeader` method to set a header for the table, and a `AddFields` method to add lines to the table. It also provides a `PrintTable` method to print the table to a writer.

`print` don't use external dependencies.

## Usage

A working example can be found in the `example` folder.

```go
package main

import "github.com/orange-cloudavenue/common-go/print"

func main() {
    p := print.New()
    p.SetHeader("String", "Int", "Bool", "Slice", "Map", "Struct", "Array")
    p.AddFields("I'm a string", 1, true, []string{"a", "b", "c"}, map[string]string{"a": "1", "b": "2", "c": "3"}, struct{ key1, key2, key3 string }{"a", "b", "c"}, [3]string{"a", "b", "c"})
    p.PrintTable()
}
```

## Output

```sh
STRING           INT       BOOL        SLICE     MAP             STRUCT                   ARRAY     
I'm a string     1         Enabled     a,b,c     a:1,b:2,c:3     key1:a,key2:b,key3:c     a,b,c   
```

## Custom `io.writer`

You can use a custom `io.Writer` to print the table.

```go
package main

import "github.com/orange-cloudavenue/common-go/print"

func main() {
 // Create io.writer
 w := new(bytes.Buffer)

 x := print.New(print.WithOutput(w))
 x.SetHeader("String", "Int", "Bool", "Slice", "Map", "Struct", "Array")
 x.AddFields("I'm a string", 1, true, []string{"a", "b", "c"}, map[string]string{"a": "1", "b": "2", "c": "3"}, struct{ key1, key2, key3 string }{"a", "b", "c"}, [3]string{"a", "b", "c"})
 x.PrintTable()

 // Print the output
 println(w.String())
}
```
