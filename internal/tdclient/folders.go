package tdclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"terraform-provider-td/internal/tdclient/models"
)

func (c *Client) GetFolders() (*models.Folders, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/entities/by-folder/391382", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	folders := &models.Folders{}
	err = json.Unmarshal(body, folders)
	if err != nil {
		return nil, err
	}

	return folders, nil
}

func (c *Client) CreateFolder(folder models.Folder) (*models.Folder, error) {
	reqBody, err := json.Marshal(folder)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/entities/folders", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newFolderResponse := &models.FolderReponse{}
	err = json.Unmarshal(body, newFolderResponse)
	if err != nil {
		return nil, err
	}

	return &newFolderResponse.Data, nil
}

func (c *Client) GetFolder(id string) (*models.Folder, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/entities/folders/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	folder := &models.FolderReponse{}
	err = json.Unmarshal(body, folder)
	if err != nil {
		return nil, err
	}

	return &folder.Data, nil
}

func (c *Client) UpdateFolder(folder models.Folder) (*models.Folder, error) {
	reqBody, err := json.Marshal(folder)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/entities/folders/%s", c.HostURL, folder.ID), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

    req.Header.Set("Content-Type", "application/json")

    body, err := c.doRequest(req)
    if err != nil {
        return nil, fmt.Errorf("error updating folder: %s", err)
    }

    updatedFolderResponse := &models.FolderReponse{}
    err = json.Unmarshal(body, updatedFolderResponse)
    if err != nil {
        return nil, err
    }

    return &updatedFolderResponse.Data, nil
}

func (c *Client) DeleteFolder(id string) error {
    req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/entities/folders/%s", c.HostURL, id), nil)
    if err != nil {
        return err
    }

    _, err = c.doRequest(req)
    if err != nil {
        return err
    }

    return nil
}
