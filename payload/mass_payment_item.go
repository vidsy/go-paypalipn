package payload

type (
	// MassPaymentItem contains details about a specific mass payment
	// from a mass payment IPN response.
	MassPaymentItem struct {
		ReasonCode    string  `form_field:"reason_code"`
		MassPayTxnID  string  `form_field:"masspay_txn_id"`
		Currency      string  `form_field:"mc_currency"`
		Fee           float64 `form_field:"mc_fee"`
		Gross         float64 `form_field:"mc_gross"`
		ReceiverEmail string  `form_field:"receiver_email"`
		Status        string  `form_field:"status"`
		UniqueID      string  `form_field:"unique_id"`
	}
)
