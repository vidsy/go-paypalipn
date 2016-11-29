package payload_test

import (
	"reflect"
	"testing"

	"github.com/vidsy/go-paypalipn/payload"
)

type (
	MockProcessor struct {
		mockProcess func(reflect.Type, reflect.Value, *payload.Values, int, string) []error
	}
)

func (ml *MockProcessor) Process(itemType reflect.Type, itemValue reflect.Value, formData *payload.Values, fieldIndex int, countField string) []error {
	if ml.mockProcess != nil {
		return ml.mockProcess(itemType, itemValue, formData, fieldIndex, countField)
	}
	return []error{}
}

func TestMassPayment(t *testing.T) {
	t.Run(".Load()", func(t *testing.T) {
		t.Run("CallsProcessor", func(t *testing.T) {
			mockProcessor := &MockProcessor{
				mockProcess: func(itemType reflect.Type, itemValue reflect.Value, formData *payload.Values, fieldIndex int, countField string) []error {
					if itemType.Elem().Kind().String() != "struct" {
						t.Fatalf("Expected type to be 'struct', got: '%s'", itemType.Elem().Kind().String())
					}

					if itemValue.String() != "<*payload.MassPayment Value>" {
						t.Fatalf("Expected value to be '<*payload.MassPayment Value>', got: '%s'", itemValue.String())
					}

					return nil
				},
			}

			massPayment := payload.MassPayment{}
			massPayment.Load(&payload.Values{}, mockProcessor)
		})
	})
}
