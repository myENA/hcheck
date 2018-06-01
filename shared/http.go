package shared

import (
	"crypto/tls"
	"net/http"
	"time"

	// http transport/client wrappers
	"github.com/hashicorp/go-cleanhttp"
	"github.com/nathanejohnson/intransport"
)

// HTTPCode returns the http code of requested url
func HTTPCode(url string, timeout time.Duration, insecure bool) (int, error) {
	var client *http.Client
	var request *http.Request
	var response *http.Response
	var err error

	// build request
	if request, err = http.NewRequest("GET", url, nil); err != nil {
		return 0, err
	}

	if insecure {
		// init client using default clean transport
		// and disable certificate verification
		transport := cleanhttp.DefaultTransport()
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
		client = &http.Client{
			Transport: transport,
		}
	} else {
		// init client using intransport library to automatically
		// fetch intermediate certificates and validate the chain and
		// verify stapled OCSP responses if certificates are marked
		// as must staple
		client = intransport.NewInTransportHTTPClient(nil)
	}

	// set timeout
	client.Timeout = timeout

	// execute request
	if response, err = client.Do(request); err != nil {
		return 0, err
	}

	// we don't use the body
	response.Body.Close()

	// all okay - return code
	return response.StatusCode, nil
}
