package payload_test

import (
	"reflect"
	"testing"

	"github.com/vidsy/go-paypalipn/payload"
)

type (
	TestStruct struct {
		Foo    string             `form_field:"foo"`
		Bar    []string           `form_field:"bar"`
		Nested []TestNestedStruct `form_field:"-"`
	}

	TestNestedStruct struct {
		Baz string  `form_field:"baz"`
		Qux float64 `form_field:"qux"`
	}
)

func TestPopulate(t *testing.T) {
	t.Run(".Populate()", func(t *testing.T) {
		t.Run("FloatError", func(t *testing.T) {
			data := `qux_0=1.23&qux_1=some_float`
			values, _ := payload.NewValuesFromFormData(data)

			testStruct := &TestStruct{}
			structType := reflect.TypeOf(testStruct)
			structValue := reflect.ValueOf(testStruct)

			populate := payload.Populate{}
			errors := populate.Process(structType, structValue, values, -1, "qux")

			if len(errors) != 1 {
				t.Fatalf("Expected len(errors) to be 1, got: %d", len(errors))
			}
		})

		t.Run("PopulatesStruct", func(t *testing.T) {
			data := `foo=test_foo&bar0=bar_0&bar1=bar_1&baz_0=baz_0&baz_1=baz_1&qux_0=1.23&qux_1=1.25`
			values, _ := payload.NewValuesFromFormData(data)

			testStruct := &TestStruct{}
			structType := reflect.TypeOf(testStruct)
			structValue := reflect.ValueOf(testStruct)

			populate := payload.Populate{}
			populate.Process(structType, structValue, values, -1, "bar")

			if testStruct.Foo != "test_foo" {
				t.Fatalf("Expected .Foo to be: 'test_foo', got: '%s'", testStruct.Foo)
			}

			if len(testStruct.Bar) != 2 {
				t.Fatalf("Expected len(.Bar) to be: 2, got: %d", len(testStruct.Bar))
			}

			if len(testStruct.Nested) != 2 {
				t.Fatalf("Expected len(.Nested) to be: 2, got: %d", len(testStruct.Nested))
			}

			if testStruct.Nested[1].Baz != "baz_1" {
				t.Fatalf("Expected .Nested[1].Baz to be 'baz_1', got: '%v'", testStruct.Nested[0].Baz)
			}

			if testStruct.Nested[1].Qux != 1.25 {
				t.Fatalf("Expected .Nested[1].Qux to be 'qux_1', got: %.2f", testStruct.Nested[0].Qux)
			}

		})
	})
}
