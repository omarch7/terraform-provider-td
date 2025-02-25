package tdclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetFolders() (*Folders, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/entities/by-folder/391382", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	folders := &Folders{}
	err = json.Unmarshal(body, folders)
	if err != nil {
		return nil, err
	}

	return folders, nil
}
