package utils

import (
	"bytes"
	"encoding/json"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"strings"
)

var lexer = lexers.Get("json")
var formatter = formatters.Get("terminal")
var style = styles.Get("solarized-dark")

// PrettyPrintJSON takes a struct and returns a pretty-printed JSON string.
// It attempts to marshal the input into JSON format and then indents the result for pretty-ness.
// If marshaling or indentation fails, it returns an error.
//
// Parameters:
//   - response: The data structure to be marshaled and pretty-printed.
//
// Returns:
//   - string: A pretty-printed JSON string if successful.
func PrettyPrintJSON(response interface{}) (string, error) {
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, jsonBytes, "", "    ")
	if err != nil {
		return "", err
	}

	return prettyJSON.String(), nil
}

// HighlightJSON takes a JSON string and applies syntax highlighting using the Chroma library.
// It returns the highlighted JSON string if successful.
// If highlighting fails, an error is returned along with an empty string.
//
// Parameters:
//   - input: The JSON string to be highlighted.
//
// Returns:
//   - string: The highlighted JSON string, or an empty string if an error occurs.
//   - error: An error object.
func HighlightJSON(input string) (string, error) {
	iterator, err := lexer.Tokenise(nil, input)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	err = formatter.Format(&b, style, iterator)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
