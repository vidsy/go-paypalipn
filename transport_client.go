package paypalipn

import (
	"net/http"
)

type (
	// TransportClient interface for client providing HTTP transport
	// functionality.
	TransportClient interface {
		Do(req *http.Request) (resp *http.Response, err error)
	}
)
