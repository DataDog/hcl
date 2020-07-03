package hcl

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSONMarshalling(t *testing.T) {
	expected := `{
  "__true_bool_comments__": [
    "Some comment on true_bool."
  ],
  "true_bool": true,
  "false_bool": false,
  "str": "string",
  "float": 1.234,
  "list": [
    1,
    2,
    3
  ],
  "map": {
    "a": 1,
    "b": "str"
  },
  "block": {
    "__comments__": [
      "A block."
    ],
    "label": {
      "empty_list": [],
      "empty_map": {}
    }
  }
}`
	ast, err := ParseString(`
			// Some comment on true_bool.
			true_bool = true
			false_bool = false
			str = "string"
			float = 1.234
			list = [1, 2, 3]
			map = {
				"a": 1,
				b: "str"
			}
			// A block.
			block "label" {
				empty_list = []
				empty_map = {}
			}
		`)
	require.NoError(t, err)
	actual, err := json.MarshalIndent(ast, "", "  ")
	require.NoError(t, err)
	// fmt.Println(string(actual))
	require.Equal(t, expected, string(actual))
}
