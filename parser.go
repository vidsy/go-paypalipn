package paypalipn

import "github.com/vidsy/go-paypalipn/payload"

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

// Parse takes a string body and a struct to populate.
func (p *Parser) Parse(body string, payloadItem Loader) error {
	valid, err := Validate(body, p.environment, p.client)
	if !valid {
		return err
	}

	data, err := payload.NewValuesFromFormData(body)
	if err != nil {
		return err
	}

	payloadItem.Load(data, p.processor)
	return nil
}
