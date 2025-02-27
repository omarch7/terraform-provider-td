package tdclient 

import (
	"encoding/json"
	"fmt"
	"net/http"

    "terraform-provider-td/internal/tdclient/models"
)

func (c *Client) GetParentSegments() (*models.ParentSegments, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/entities/parent_segments", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	parentSegments := &models.ParentSegments{}
	err = json.Unmarshal(body, parentSegments)
	if err != nil {
		return nil, err
	}

	return parentSegments, nil
}
