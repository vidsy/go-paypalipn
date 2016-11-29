package paypalipn_test

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"testing"

	"github.com/vidsy/go-paypalipn"
)

type (
	MockClient struct {
		mockDo func(*http.Request) (*http.Response, error)
	}
)

func (mc MockClient) Do(req *http.Request) (*http.Response, error) {
	if mc.mockDo != nil {
		return mc.mockDo(req)
	}

	return NewMockResponse("")
}

func NewMockResponse(response string) (*http.Response, error) {
	return &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		StatusCode: 200,
	}, nil
}

func TestReadBodyIntoString(t *testing.T) {
	t.Run("ReadBodyIntoString()", func(t *testing.T) {
		t.Run("ReturnsStringFromReadCloser", func(t *testing.T) {
			reader := ioutil.NopCloser(bytes.NewBufferString("SOME=DATA"))
			data, _ := paypalipn.ReadBodyIntoString(reader)

			if data != "SOME=DATA" {
				t.Fatalf("Expected data to be 'SOME=DATA', got: '%s'", data)
			}
		})
	})
}
