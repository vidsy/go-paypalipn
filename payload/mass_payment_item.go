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

// ReasonString Returns the string representaion of an error
// response code.
func (m MassPaymentItem) ReasonString() string {
	reasonCodeMapping := map[string]string{
		"1001":  "Receiver's account is invalid",
		"1002":  "Sender has insufficient funds",
		"1003":  "User's country is not allowed",
		"1004":  "User's credit card is not in the list of allowed countries of the gaming merchant",
		"3004":  "Cannot pay self",
		"3014":  "Sender's account is locked or inactive",
		"3015":  "Receiver's account is locked or inactive",
		"3016":  "Either the sender or receiver exceeded the transaction limit",
		"3017":  "Spending limit exceeded",
		"3047":  "User is restricted",
		"3078":  "Negative balance",
		"3148":  "Receiver's address is in a non-receivable country or a PayPal zero country",
		"3535":  "Invalid currency",
		"3547":  "Sender's address is located in a restricted State (e.g., California)",
		"3558":  "Receiver's address is located in a restricted State (e.g., California)",
		"3769":  "Market closed and transaction is between 2 different countries",
		"4001":  "Internal error",
		"4002":  "Internal error",
		"8319":  "Zero amount",
		"8330":  "Receiving limit exceeded",
		"8331":  "Duplicate mass payment",
		"9302":  "Transaction was declined",
		"11711": "Per-transaction sending limit exceeded",
		"14159": "Transaction currency cannot be received by the recipient",
		"14550": "Currency compliance",
		"14761": "The mass payment was declined because the secondary user sending the mass payment has not been verified",
		"14764": "Regulatory review - Pending",
		"14765": "Regulatory review - Blocked",
		"14767": "Receiver is unregistered",
		"14768": "Receiver is unconfirmed",
		"14769": "Youth account recipient",
		"14800": "POS cumulative sending limit exceeded",
	}

	if message, exists := reasonCodeMapping[m.ReasonCode]; exists {
		return message
	}

	return "No message available for this error type"
}
