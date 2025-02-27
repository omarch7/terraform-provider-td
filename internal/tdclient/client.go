package tdclient

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	APIKey     string
}

func NewClient(host, api_key *string) (*Client, error) {
	return &Client{
		HostURL:    *host,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		APIKey:     *api_key,
	}, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "TD1 "+c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", resp.StatusCode, body)
	}

	return body, err
}
