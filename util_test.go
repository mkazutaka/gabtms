package gabtms

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnmarshalDataString(t *testing.T) {
	actual := map[string]interface{}{}
	err := UnmarshalDataString(json.RawMessage(`"{\"foo\":\"A\",\"bar\":1}"`), &actual)
	if err != nil {
		t.Errorf("Expected error to be `nil`, got %v", err)
	}

	expected := map[string]interface{}{
		"foo": "A",
		"bar": 1.0,
	}
	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("func differs: (-got +want)\n%s", diff)
	}
}
