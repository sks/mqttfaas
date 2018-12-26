package types

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

//HTTPRequest ...
type HTTPRequest struct {
	Topics      []string          `yaml:"topic"`
	HTTPPath    string            `yaml:"path"`
	OutputTopic string            `yaml:"output_topic"`
	ErrorTopic  string            `yaml:"error_topic"`
	Headers     map[string]string `yaml:"headers"`
	Query       map[string]string `yaml:"query"`
	ContentType string            `yaml:"content_type"`
}

func (h *HTTPRequest) getContentType() string {
	if h.ContentType == "" {
		return "application/octet-stream"
	}
	return h.ContentType
}

func (h *HTTPRequest) getQueryValues() string {
	urlValues := url.Values{}
	for k, v := range h.Query {
		urlValues.Add(k, v)
	}
	return urlValues.Encode()
}

func (h *HTTPRequest) getRequestURL(baseURL string) string {
	if strings.Contains(h.HTTPPath, "://") {
		return h.HTTPPath
	}
	return fmt.Sprintf("%s/%s", baseURL, h.HTTPPath)
}

//CreateRequest create a request out of the HTTPRequest definition
func (h *HTTPRequest) CreateRequest(baseURL string, data io.Reader) (*http.Request, error) {
	url := h.getRequestURL(baseURL)
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return nil, err
	}
	for k, v := range h.Headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Content-Type", h.getContentType())
	req.URL.RawQuery = h.getQueryValues()
	return req, nil
}

//ShouldTriggerFor returns true if the HTTPRequest has to be invoked
func (h *HTTPRequest) ShouldTriggerFor(incomingTopic string) bool {
	for i := range h.Topics {
		if isSubTopic(incomingTopic, h.Topics[i]) {
			return true
		}
	}
	return false
}
