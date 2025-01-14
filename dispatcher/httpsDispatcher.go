package dispatcher

import (
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"
	"time"
)

type HTTPSDispatcher struct {
	Client *http.Client
}

// NewHTTPDispatcher initializes the HTTPDispatcher with a customizable timeout and TLS configuration.
func NewHTTPSDispatcher(timeout time.Duration, insecureSkipVerify bool) *HTTPSDispatcher {
	// Create a custom transport with TLS configuration for HTTPS calls
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecureSkipVerify, // Set to true if you want to skip TLS verification (not recommended for production)
		},
	}

	return &HTTPSDispatcher{
		Client: &http.Client{
			Timeout:   timeout,
			Transport: transport,
		},
	}
}

// CallBackend makes an HTTPS request to the specified backend with method, headers, and query parameters.
func (d *HTTPSDispatcher) CallBackend(method HTTPMethod, baseURL string, headers map[string]string, queryParams map[string]string) (*http.Response, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	q := parsedURL.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	parsedURL.RawQuery = q.Encode()

	req, err := http.NewRequest(string(method), parsedURL.String(), nil)
	if err != nil {
		return nil, err
	}

	// Add headers to the request
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Make the HTTPS request
	resp, err := d.Client.Do(req)
	if err != nil {
		return nil, err
	}

	// Check for non-OK status code
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to call backend")
	}

	return resp, nil
}
