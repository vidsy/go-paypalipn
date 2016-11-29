package paypalipn

import (
	"io"
	"io/ioutil"

	"github.com/vidsy/go-paypalipn/payload"
)

type (
	// Loader interface for structs that can be loaded
	// from give form data.
	Loader interface {
		Load(*payload.Values, payload.Processor)
	}
)

// Parse takes a http response and struct to populate
// based on post data in response.
func Parse(body io.ReadCloser, payloadItem Loader, processor payload.Processor) error {
	parsedBody, err := readBody(body)
	if err != nil {
		return err
	}

	data, err := payload.NewValuesFromFormData(parsedBody)
	if err != nil {
		return err
	}

	payloadItem.Load(data, processor)
	return nil
}

func readBody(rawBody io.ReadCloser) (string, error) {
	body, err := ioutil.ReadAll(io.LimitReader(rawBody, 1048576))
	if err != nil {
		return "", err
	}

	if err = rawBody.Close(); err != nil {
		return "", err
	}

	return string(body[:]), nil
}
