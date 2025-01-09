package print

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

type (
	Writer struct {
		header []string
		opts   writerOps
		tw     *tabwriter.Writer
	}

	writerOps struct {
		output io.Writer
	}

	WriterOption func(w *writerOps)
)

// New returns a new instance of the Writer struct.
func New(opts ...WriterOption) Writer {
	o := writerOps{
		output: os.Stdout,
	}

	for _, opt := range opts {
		opt(&o)
	}

	return Writer{
		header: []string{},
		opts:   o,
		tw:     tabwriter.NewWriter(o.output, 10, 1, 5, ' ', 0),
	}
}

// WithOutput sets the output field of the writerOps struct.
func WithOutput(output io.Writer) WriterOption {
	return func(w *writerOps) {
		w.output = output
	}
}

// SetHeader sets the header fields for the Writer.
// It takes a variadic parameter fields of type any, representing the header fields.
// If no fields are provided, the function returns without making any changes.
// Each field in the fields parameter is converted to a string and stored in the header slice of the Writer.
// The converted fields are also appended to the underlying text writer, tw.
// If any field in the fields parameter is not a string, the function panics with an error message.
// The header fields are converted to uppercase before being stored in the header slice.
// The formatted header fields are then written to the text writer using the format function.
func (w *Writer) SetHeader(fields ...any) {
	if len(fields) == 0 {
		return
	}

	for i, field := range fields {
		v, ok := field.(string)
		if !ok {
			panic("Header fields must be string")
		}
		fields[i] = strings.ToUpper(v)
		w.header = append(w.header, v)
	}

	fmt.Fprintf(w.tw, format(fields...), fields...)
}

// AddFields adds fields to the writer.
// It takes a variadic parameter fields of type any.
// If the length of fields is zero, it returns immediately.
// If the length of fields is not equal to the length of the writer's header, it panics with the message "Fields must be the same length as header".
// It formats each field using the fieldFormat function.
// It then formats the fields using the format function and writes them to the writer's tw.
// Finally, it writes the formatted fields to the writer's tw using fmt.Fprintf.
func (w Writer) AddFields(fields ...any) {
	if len(fields) == 0 {
		return
	}

	if len(fields) != len(w.header) {
		panic("Fields must be the same length as header")
	}

	for i, field := range fields {
		fields[i] = fieldFormat(field)
	}

	fmt.Fprintf(w.tw, format(fields...), fields...)
}

// Print a line
func (w Writer) PrintTable() {
	w.tw.Flush()
}
