package payload

import (
	"reflect"
	"time"
)

type (
	// MassPayment contains all the information
	// from a mass payment IPN response.
	MassPayment struct {
		PaymentDate   time.Time         `form_field:"payment_date"`
		PaymentStatus string            `form_field:"payment_status"`
		IPNTrackID    string            `form_field:"ipn_track_id"`
		PayerID       string            `form_field:"payer_id"`
		TXNType       string            `form_field:"txn_type"`
		Handling      []string          `form_field:"mc_handling"`
		Items         []MassPaymentItem `form_field:"-"`
	}
)

const (
	indexCountField = "status"
)

// Load takes the form data and loads it into the struct.
func (m *MassPayment) Load(formData *Values, processor Processor) []error {
	massPaymentType := reflect.TypeOf(m)
	massPaymentValue := reflect.ValueOf(m)
	return processor.Process(massPaymentType, massPaymentValue, formData, -1, indexCountField)
}
