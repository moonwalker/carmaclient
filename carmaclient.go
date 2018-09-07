package carmaclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"bytes"
	"encoding/json"
	"github.com/moonwalker/logger"
	"errors"
)

const (
	timeout         = 30 * time.Second
	carmaAuthHeader = "X-Carma-Authentication-Token"
)

type CarmaClient struct {
	HTTPClient *http.Client

	carmaApiToken      string
	carmaOrganization  int64
	carmaRESTURL       string
	carmaControllerURL string

	log logger.Logger

	common     service
	Triggers   *TriggersService
	Lists      *ListsService
	Properties *PropertiesService
}

type service struct {
	client *CarmaClient
}

type carmaResponse struct {
	statusCode int
	data       []byte
	err        error
	msg        string
}

type carmaError struct {
	Error string `json:"error"`
}

func NewCarmaClient(organization int64, token, restURL, controllerURL string, log logger.Logger) (client *CarmaClient, err error) {
	client = &CarmaClient{
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
		carmaApiToken:      token,
		carmaOrganization:  organization,
		carmaRESTURL:       restURL,
		carmaControllerURL: controllerURL,
		log:                log,
	}

	client.common.client = client

	client.Triggers = (*TriggersService)(&client.common)
	client.Lists = (*ListsService)(&client.common)
	client.Properties = (*PropertiesService)(&client.common)

	return
}

func (c CarmaClient) catchError(carmaResp *carmaResponse) {
	carmaError := &carmaError{}
	err := json.Unmarshal(carmaResp.data, carmaError)
	if err != nil {
		return
	}

	carmaResp.err = errors.New(carmaError.Error)
}

func (c CarmaClient) carmaRequest(endpoint string, method string, body interface{}) (carmaResp carmaResponse) {
	logRequest := make(map[string]interface{})
	logResponse := make(map[string]interface{})

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	url := fmt.Sprintf("%s/%d/%s", c.carmaRESTURL, c.carmaOrganization, endpoint)

	b := new(bytes.Buffer)
	if body != nil {
		json.NewEncoder(b).Encode(body)
	}

	logData, err := json.Marshal(body)
	if err == nil {
		logRequest["Body"] = string(logData)
	}

	req, err := http.NewRequest(method, url, b)
	if err != nil {
		carmaResp.err = err
		return
	}

	req.Header[carmaAuthHeader] = []string{c.carmaApiToken}
	req.Header["Accept"] = []string{"application/json"}
	req.Header["Content-Type"] = []string{"application/json"}

	logResponse["query"] = endpoint
	logRequest["query"] = endpoint

	logRequest["URL"] = url

	if c.log != nil {
		c.log.Info("Carma Request", logRequest)
	}

	resp, err := client.Do(req)
	if err != nil {
		carmaResp.err = err
		return
	}

	defer resp.Body.Close()

	carmaResp.statusCode = resp.StatusCode

	respBody, err := ioutil.ReadAll(resp.Body)
	carmaResp.err = err
	carmaResp.data = respBody
	c.catchError(&carmaResp)

	logResponse["Response"] = string(respBody)

	logResponse["StatusCode"] = resp.StatusCode
	logResponse["URL"] = url

	if c.log != nil {
		c.log.Info("Carma Response", logResponse)
	}

	return
}
