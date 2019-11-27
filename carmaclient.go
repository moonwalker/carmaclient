package carmaclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/moonwalker/logger"
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
	carmaResp.err = nil
	logData := make(map[string]interface{})

	url := fmt.Sprintf("%s/%d/%s", c.carmaRESTURL, c.carmaOrganization, endpoint)
	logData["query"] = endpoint
	logData["URL"] = url
	logData["Body"] = body

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	b := new(bytes.Buffer)
	if body != nil {
		json.NewEncoder(b).Encode(body)
	}

	req, err := http.NewRequest(method, url, b)
	if err != nil {
		carmaResp.err = err
		if c.log != nil {
			logData["error"] = err
			c.log.Info("Carma Request", logData)
		}
		return
	}

	req.Header[carmaAuthHeader] = []string{c.carmaApiToken}
	req.Header["Accept"] = []string{"application/json"}
	req.Header["Content-Type"] = []string{"application/json"}

	resp, err := client.Do(req)
	if err != nil {
		carmaResp.err = err
		if c.log != nil {
			logData["error"] = err
			c.log.Info("Carma Request", logData)
		}
		return
	}

	defer resp.Body.Close()

	carmaResp.statusCode = resp.StatusCode

	respBody, err := ioutil.ReadAll(resp.Body)
	carmaResp.err = err
	carmaResp.data = respBody
	if carmaResp.err != nil {
		c.catchError(&carmaResp)
	}

	var jsonResp interface{}
	err = json.Unmarshal(respBody, jsonResp)
	if err == nil {
		logData["Response"] = respBody
	}
	logData["StatusCode"] = resp.StatusCode

	if c.log != nil {
		c.log.Info("Carma Request", logData)
	}

	return
}
