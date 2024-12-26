package dispatcher

import (
	"errors"
	"net/http"
	"net/url"
	"time"
)

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
	PATCH  HTTPMethod = "PATCH"
)

type HTTPDispatcher struct {
	Client *http.Client
}

func NewHTTPDispatcher(timeout time.Duration) *HTTPDispatcher {
	return &HTTPDispatcher{
		Client: &http.Client{Timeout: timeout},
	}
}

func (d *HTTPDispatcher) CallBackend(method HTTPMethod, baseURL string, headers map[string]string, queryParams map[string]string) (*http.Response, error) {
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

	// Make the request
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
