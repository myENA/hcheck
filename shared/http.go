package shared

import (
	"crypto/tls"
	"net/http"
	"time"

	// hashicorp's handy http client wrapper
	"github.com/hashicorp/go-cleanhttp"
)

// HTTPCode returns the http code of requested url
func HTTPCode(url string, timeout time.Duration, insecure bool) (int, error) {
	var transport = cleanhttp.DefaultTransport()
	var client *http.Client
	var request *http.Request
	var response *http.Response
	var err error

	// build request
	if request, err = http.NewRequest("GET", url, nil); err != nil {
		return 0, err
	}

	// modify transport as needed
	if insecure {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	// build client with transport and timeout
	client = &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	// execute request
	if response, err = client.Do(request); err != nil {
		return 0, err
	}

	// we don't use the body
	response.Body.Close()

	// all okay - return code
	return response.StatusCode, nil
}
