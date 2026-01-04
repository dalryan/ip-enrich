package output

import (
	"encoding/json"
	"fmt"
	"io"
)

// JSONFormatter outputs results as JSON.
type JSONFormatter struct {
	writer io.Writer
	pretty bool
}

// NewJSONFormatter creates a new JSON formatter.
// If pretty is true, output is indented for readability.
func NewJSONFormatter(w io.Writer, pretty bool) *JSONFormatter {
	return &JSONFormatter{
		writer: w,
		pretty: pretty,
	}
}

// Format writes the report as JSON.
func (f *JSONFormatter) Format(report *Report) error {
	var data []byte
	var err error

	if f.pretty {
		data, err = json.MarshalIndent(report, "", "  ")
	} else {
		data, err = json.Marshal(report)
	}

	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	_, err = f.writer.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	if f.pretty {
		_, _ = f.writer.Write([]byte("\n"))
	}

	return nil
}
