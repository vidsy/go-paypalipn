package paypalipn_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/vidsy/go-paypalipn"
)

func TestValidate(t *testing.T) {
	t.Run("Validate()", func(t *testing.T) {
		t.Run("ReturnTrueWithValidReponse", func(t *testing.T) {
			mockClient := &MockClient{
				mockDo: func(req *http.Request) (*http.Response, error) {
					return NewMockResponse("VERIFIED")
				},
			}

			valid, err := paypalipn.Validate("", paypalipn.Sandbox, mockClient)

			if err != nil {
				t.Fatalf("Expected error, to be nil got: %s", err)
			}

			if !valid {
				t.Fatalf("Expected valid to be true, got: %t", valid)
			}

		})

		t.Run("CorrectEndpointChosenForEnvironment", func(t *testing.T) {
			var envTestCases = []struct {
				env              string
				expectedEndpoint string
			}{
				{paypalipn.Sandbox, "https://ipnpb.sandbox.paypal.com/cgi-bin/webscr"},
				{paypalipn.Live, "https://ipnpb.paypal.com/cgi-bin/webscr"},
			}

			for _, testCase := range envTestCases {
				mockClient := &MockClient{
					mockDo: func(req *http.Request) (*http.Response, error) {
						if req.URL.String() != testCase.expectedEndpoint {
							t.Fatalf("Expected Request.URL to be '%s' in env '%s', got: '%s'", testCase.expectedEndpoint, testCase.env, req.URL.String())
						}

						return nil, errors.New("Client error")
					},
				}

				paypalipn.Validate("", testCase.env, mockClient)
			}
		})

		t.Run("ReturnsFalseOnClientError", func(t *testing.T) {
			mockClient := &MockClient{
				mockDo: func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("Client error")
				},
			}

			valid, err := paypalipn.Validate("", paypalipn.Sandbox, mockClient)

			if err == nil {
				t.Fatalf("Expected error, got: %s", err)
			}

			if valid {
				t.Fatalf("Expected valid to be false, got: %t", valid)
			}
		})

		t.Run("ReturnsFalseIfPayloadInvalid", func(t *testing.T) {
			mockClient := &MockClient{
				mockDo: func(req *http.Request) (*http.Response, error) {
					return NewMockResponse("INVALID")
				},
			}

			valid, err := paypalipn.Validate("", paypalipn.Sandbox, mockClient)

			if err == nil {
				t.Fatalf("Expected error, got: %s", err)
			}

			if valid {
				t.Fatalf("Expected valid to be false, got: %t", valid)
			}
		})

		t.Run("ReturnsFalseIfPayloadResponseNotExpected", func(t *testing.T) {
			mockClient := &MockClient{
				mockDo: func(req *http.Request) (*http.Response, error) {
					return NewMockResponse("uunexpected_response_body")
				},
			}

			valid, err := paypalipn.Validate("", paypalipn.Sandbox, mockClient)

			if err == nil {
				t.Fatalf("Expected error, got: %s", err)
			}

			if valid {
				t.Fatalf("Expected valid to be false, got: %t", valid)
			}
		})
	})
}
