package paypalipn_test

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/vidsy/go-paypalipn"
	"github.com/vidsy/go-paypalipn/payload"
)

type (
	MockLoader struct {
		mockLoad func(*payload.Values, payload.Processor)
		Data     *payload.Values
	}

	MockProcessor struct {
		mockProcess func(itemType reflect.Type, itemValue reflect.Value, formData *payload.Values, fieldIndex int, countField string) []error
	}
)

func (ml *MockLoader) Load(data *payload.Values, processor payload.Processor) {
	if ml.mockLoad != nil {
		ml.mockLoad(data, processor)
	}

	ml.Data = data
}

func (mp *MockProcessor) Process(itemType reflect.Type, itemValue reflect.Value, formData *payload.Values, fieldIndex int, countField string) []error {
	if mp.mockProcess != nil {
		return mp.mockProcess(itemType, itemValue, formData, fieldIndex, countField)
	}

	return []error{}
}

func TestProcessor(t *testing.T) {
	t.Run("Parse()", func(t *testing.T) {
		t.Run("ParsesBodyIntoLoader", func(t *testing.T) {
			mockLoader := MockLoader{}
			rawBody := ioutil.NopCloser(bytes.NewBufferString(`KEY=VALUE`))

			paypalipn.Parse(rawBody, &mockLoader, &MockProcessor{})
			if mockLoader.Data.Get("KEY") != "VALUE" {
				t.Fatalf("Expected Loader.Data.Get('KEY') to be '%s', got: '%s'", "VALUE", mockLoader.Data.Get("KEY"))
			}
		})
	})
}
