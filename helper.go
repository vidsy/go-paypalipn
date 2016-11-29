package paypalipn

import (
	"io"
	"io/ioutil"
)

// ReadBodyIntoString reads an io.ReadCloser, converts it to a byte array
// and returns the resulting string.
func ReadBodyIntoString(rawBody io.ReadCloser) (string, error) {
	body, err := ioutil.ReadAll(io.LimitReader(rawBody, 1048576))
	if err != nil {
		return "", err
	}

	if err = rawBody.Close(); err != nil {
		return "", err
	}

	return string(body[:]), nil
}
