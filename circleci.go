package circleci

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
)

const (
	apiURL       = "https://circleci.com/api/v1"
	acceptHeader = "application/json" // For machine readable JSON
)

// Client manages requests to & responses from the CircleCi API
type Client struct {
	apiToken string
}

// NewClient returns a new Client which will use the provided API
// token for authentication
func NewClient(apiToken string) *Client {
	return &Client{apiToken: apiToken}
}

func (c *Client) get(path string, v interface{}) error {
	req, _ := c.newRequest("GET", path)

	return c.do(req, v)
}

// newRequest creates a new API request using the provided method
// and path, adding the required headers and authentication query
// parameters.
func (c *Client) newRequest(method, path string) (*http.Request, error) {
	path = apiURL + path
	// TODO, generalise logging
	log.Print(method, path)

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		log.Print(err)
	}

	// Add headers
	req.Header.Add("Accept", acceptHeader)
	req.Header.Add("User-Agent", "go-circleci/"+version)

	// Add api token query param
	v := url.Values{}
	v.Set("circle-token", c.apiToken)
	req.URL.RawQuery = v.Encode()

	return req, nil
}

// do sends an API request, checks the response, and decodes the
// response into the provided interface.
func (c *Client) do(req *http.Request, v interface{}) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Print("Bad request", req.URL)
		return errors.New("Bad request")
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			log.Print("Decode problem")
			return err
		}
	}

	return nil
}
