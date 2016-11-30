package payload_test

import (
	"testing"

	"github.com/vidsy/go-paypalipn/payload"
)

func TestValues(t *testing.T) {
	t.Run("NewFromFromData()", func(t *testing.T) {
		t.Run("ParsesFormData", func(t *testing.T) {
			formData := "foo=bar"
			values, _ := payload.NewValuesFromFormData(formData)

			if values == nil {
				t.Fatalf("Expected ParsesFormData to return Values struct, got: '%s'", values)
			}
		})
	})

	t.Run(".Get()", func(t *testing.T) {
		t.Run("WithSingleValueFields", func(t *testing.T) {
			formData := "foo=bar&baz=qux"
			values, _ := payload.NewValuesFromFormData(formData)

			if values.Get("baz") != "qux" {
				t.Fatalf("Expected .Get('baz') to be 'qux', got: '%s'", values.Get("baz"))
			}
		})

		t.Run("WithMultipleValuesOfSameName", func(t *testing.T) {
			formData := "foo=bar&foo0=a&foo1=b&foo_0=c&foo_1=d"
			values, _ := payload.NewValuesFromFormData(formData)

			if values.Get("foo") != "bar" {
				t.Fatalf("Expected .Get('foo') to be 'bar', got: '%s'", values.Get("baz"))
			}
		})
	})

	t.Run(".GetValues()", func(t *testing.T) {
		t.Run("InUnderScoreFormat", func(t *testing.T) {
			formData := "foo_1=bar&foo_2=qux"
			values, _ := payload.NewValuesFromFormData(formData)

			fields := values.GetValues("foo", 2)

			if len(fields) != 2 {
				t.Fatalf(`Expected .GetValues("foo") to be 2, got: %d`, len(fields))
			}

			if fields[1] != "qux" {
				t.Fatalf(`Expected .GetValues("foo") to be 'qux', got: %v`, fields[1])
			}
		})

		t.Run("NumericSuffix", func(t *testing.T) {
			formData := "foo=test&foo1=bar&foo2=qux"
			values, _ := payload.NewValuesFromFormData(formData)

			fields := values.GetValues("foo", 2)

			if len(fields) != 2 {
				t.Fatalf(`Expected .GetValues("foo") to be 2, got: %d`, len(fields))
			}

			if fields[1] != "qux" {
				t.Fatalf(`Expected .GetValues("foo") to be 'qux', got: %v`, fields[1])
			}
		})
	})

	t.Run(".GetValueAtIndex()", func(t *testing.T) {
		t.Run("InUnderScoreFormat", func(t *testing.T) {
			formData := "foo=bop&foo_1=bar&foo_2=qux"
			values, _ := payload.NewValuesFromFormData(formData)
			value := values.GetValueAtIndex("foo", 1)

			if value != "qux" {
				t.Fatalf("Expected 'qux', got: '%s'", value)
			}

		})

		t.Run("NumericSuffix", func(t *testing.T) {
			formData := "foo=test&foo1=bar&foo2=qux"
			values, _ := payload.NewValuesFromFormData(formData)
			value := values.GetValueAtIndex("foo", 1)

			if value != "qux" {
				t.Fatalf("Expected 'qux', got: '%s'", value)
			}
		})
	})

	t.Run(".ItemCount()", func(t *testing.T) {
		t.Run("InUnderScoreFormat", func(t *testing.T) {
			formData := "foo=bop&foo_1=bar&foo_2=qux"
			values, _ := payload.NewValuesFromFormData(formData)
			count := values.ItemCount("foo")
			if count != 2 {
				t.Fatalf("Expected .ItemCount() to be 2, got: %d", count)
			}

		})

		t.Run("NumericSuffix", func(t *testing.T) {
			formData := "foo=test&foo1=bar&foo2=qux"
			values, _ := payload.NewValuesFromFormData(formData)
			count := values.ItemCount("foo")
			if count != 2 {
				t.Fatalf("Expected .ItemCount() to be 2, got: %d", count)
			}
		})

	})
}
