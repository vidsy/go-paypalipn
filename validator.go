package paypalipn

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

const (
	// Live indicates the validation request should be made against the live PayPal environment.
	Live = "live"

	// Sandbox indicates the validation request should be made against the sandbox PayPal environment.
	Sandbox = "sandbox"

	sandboxIPNValidationEndpoint = "https://ipnpb.sandbox.paypal.com/cgi-bin/webscr"
	liveIPNValidationEndpoint    = "https://ipnpb.paypal.com/cgi-bin/webscr"
)

// Validate takes a PayPal IPN body, environment and http client and validates it
// based on the IPN spec: https://developer.paypal.com/docs/classic/ipn/integration-guide/IPNImplementation/#specs
func Validate(body string, environment string, httpClient TransportClient) (bool, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	request, _ := http.NewRequest(
		"POST",
		generateEndpoint(environment),
		bytes.NewBuffer([]byte(applyValidationParameter(body))),
	)

	response, err := httpClient.Do(request)
	if err != nil {
		return false, err
	}

	validationBody, err := ReadBodyIntoString(response.Body)
	if err != nil {
		return false, err
	}

	switch validationBody {
	case "VERIFIED":
		return true, nil
	case "INVALID":
		return false, errors.New(
			"IPN payload sent by PayPal didn't match payload received",
		)
	default:
		return false, fmt.Errorf(
			"Expected response to include: 'VERIFIED' or 'INVALID' but got: '%s'",
			validationBody,
		)
	}

}

func applyValidationParameter(body string) string {
	return fmt.Sprintf("%s&cmd=_notify-validate", body)
}

func generateEndpoint(environment string) string {
	if environment == Sandbox {
		return sandboxIPNValidationEndpoint
	}

	return liveIPNValidationEndpoint
}
