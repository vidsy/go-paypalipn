package paypalipn_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/vidsy/go-paypalipn"
	"github.com/vidsy/go-paypalipn/payload"
)

type (
	MockLoader struct {
		mockLoad func(*payload.Values, payload.Processor) []error
		Data     *payload.Values
	}

	MockProcessor struct {
		mockProcess func(itemType reflect.Type, itemValue reflect.Value, formData *payload.Values, fieldIndex int, countField string) []error
	}
)

func (ml *MockLoader) Load(data *payload.Values, processor payload.Processor) []error {
	if ml.mockLoad != nil {
		return ml.mockLoad(data, processor)
	}

	ml.Data = data
	return nil
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
			mockClient := &MockClient{
				mockDo: func(req *http.Request) (*http.Response, error) {
					return NewMockResponse("VERIFIED")
				},
			}

			rawBody := ioutil.NopCloser(bytes.NewBufferString(`KEY=VALUE`))
			parser := paypalipn.NewParser(paypalipn.Sandbox, mockClient, &MockProcessor{})
			parser.Parse(rawBody, &mockLoader)

			if mockLoader.Data.Get("KEY") != "VALUE" {
				t.Fatalf("Expected Loader.Data.Get('KEY') to be '%s', got: '%s'", "VALUE", mockLoader.Data.Get("KEY"))
			}
		})
	})
}
