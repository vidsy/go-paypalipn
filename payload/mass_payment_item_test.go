package payload_test

import (
	"testing"

	"github.com/vidsy/go-paypalipn/payload"
)

func TestMassMpaymentItem(t *testing.T) {
	t.Run(".ReasonString()", func(t *testing.T) {
		t.Run("WithValidReasonCode", func(t *testing.T) {
			massPaymentItem := payload.MassPaymentItem{ReasonCode: "1001"}
			expectedMessage := "Receiver's account is invalid"

			if massPaymentItem.ReasonString() != expectedMessage {
				t.Fatalf("Expected .ReasonString() to be '%s', got: '%s'", expectedMessage, massPaymentItem.ReasonString())
			}
		})

		t.Run("WithInvalidResponseCode", func(t *testing.T) {
			massPaymentItem := payload.MassPaymentItem{ReasonCode: "1"}
			expectedMessage := "No message available for this error type"

			if massPaymentItem.ReasonString() != expectedMessage {
				t.Fatalf("Expected .ReasonString() to be '%s', got: '%s'", expectedMessage, massPaymentItem.ReasonString())
			}
		})
	})
}
