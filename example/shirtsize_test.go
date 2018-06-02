package example_test

import (
	"encoding"
	"testing"

	"github.com/alext/textenums/example"
)

func TestShirtSize(t *testing.T) {
	s := example.M
	if _, ok := interface{}(s).(encoding.TextMarshaler); !ok {
		t.Fatalf("Expected ShirtSize to satisfy encoding.TextMarshaler interface")
	}
	if _, ok := interface{}(&s).(encoding.TextUnmarshaler); !ok {
		t.Fatalf("Expected ShirtSize to satisfy encoding.TextUnmarshaler interface")
	}

	tests := map[example.ShirtSize]string{
		example.NA: `NA`,
		example.XS: `XS`,
		example.S:  `S`,
		example.M:  `M`,
		example.L:  `L`,
		example.XL: `XL`,
	}
	for shirtsize, str := range tests {
		actual, err := shirtsize.MarshalText()
		if err != nil {
			t.Errorf("Error marshalling %s: %s", str, err.Error())
			continue
		}
		if string(actual) != str {
			t.Errorf("Marshalling %s, want: %s, got: %s", str, str, actual)
		}

		var ss example.ShirtSize
		err = ss.UnmarshalText([]byte(str))
		if err != nil {
			t.Errorf("Error unmarshalling %s: %s", str, err.Error())
			continue
		}
		if ss != shirtsize {
			t.Errorf("Unmarshalling %s, want: %d, got: %d", str, shirtsize, ss)
		}
	}
}
