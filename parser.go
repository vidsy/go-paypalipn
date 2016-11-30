package paypalipn

import (
	"io"

	"github.com/vidsy/go-paypalipn/payload"
)

type (
	// Parser takes an environment, transportClient and processor
	// and parses an IPN paypal populating the passed in struct.
	Parser struct {
		environment string
		client      TransportClient
		processor   payload.Processor
	}

	// Loader interface for structs that can be loaded
	// from give form data.
	Loader interface {
		Load(*payload.Values, payload.Processor) []error
	}
)

// NewParser creates a new Parser.
func NewParser(environment string, client TransportClient, processor payload.Processor) *Parser {
	return &Parser{environment, client, processor}
}

// Parse takes a http response and struct to populate
// based on post data in response.
func (p *Parser) Parse(body io.ReadCloser, payloadItem Loader) error {
	parsedBody, err := ReadBodyIntoString(body)
	if err != nil {
		return err
	}

	valid, err := Validate(parsedBody, p.environment, p.client)
	if !valid {
		return err
	}

	data, err := payload.NewValuesFromFormData(parsedBody)
	if err != nil {
		return err
	}

	payloadItem.Load(data, p.processor)
	return nil
}
