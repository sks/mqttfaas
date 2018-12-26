package httpinvoker

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sks/mqttfaas/internal/retrier"
	"github.com/sks/mqttfaas/internal/types"
)

//HTTPInvoker ...
type HTTPInvoker struct {
	apiPrefix string
}

//New ...
func New(apiPrefix string) *HTTPInvoker {
	return &HTTPInvoker{
		apiPrefix,
	}
}

//Call call the given function
func (h *HTTPInvoker) Call(httpRequest *types.HTTPRequest, topic string, data []byte) (string, []byte) {
	var resp []byte
	var err error
	err = retrier.Call(3, time.Second, func() error {
		resp, err = h.makeRequest(httpRequest, topic, data)
		return err
	})
	if err != nil {
		return httpRequest.ErrorTopic, resp
	}
	return httpRequest.OutputTopic, resp
}

func (h *HTTPInvoker) makeRequest(httpRequest *types.HTTPRequest, topic string, data []byte) ([]byte, error) {
	requestBody := bytes.NewReader(data)
	req, err := httpRequest.CreateRequest(h.apiPrefix, requestBody)
	if err != nil {
		return nil, err
	}
	//Add custom header to indicate the topic
	req.Header.Add("X-Databus-Topic", topic)
	log.Printf("Invoking the function %q based on message on topic %q\n", req.URL, topic)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//error in case of non 2xx status code
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, errors.New(string(body))
	}
	return body, nil
}
