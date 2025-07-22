package save

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func (c *ApiCLient) Get(enpoint string) ([]byte, error) {
	url := c.BaseURL + enpoint
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error while executing GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("failed request: %s, status code: %d", string(body), resp.StatusCode)
	}

	return body, nil
}

func (c *ApiCLient) Post(enpoint string, body []byte) ([]byte, int, error) {
	url := c.BaseURL + enpoint
	resp, err := c.Client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, 0, fmt.Errorf("error while executing POST request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, resp.StatusCode, fmt.Errorf("failed request: %s, status code: %d", string(responseBody), resp.StatusCode)
	}

	return responseBody, resp.StatusCode, nil
}
