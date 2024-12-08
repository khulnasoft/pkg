package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/pkg/errors"
)

type KhulnasoftLabel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Created     string `json:"created"`
	Author      string `json:"author"`
}

type KhulnasoftLabels struct {
	KhulnasoftLabels []KhulnasoftLabel `json:"result"`
}

// GetKhulnasoftLabel - get a single Khulnasoft label
func (cli *Client) GetKhulnasoftLabel(name string) (*KhulnasoftLabel, error) {
	var err error
	var response KhulnasoftLabel
	apiPath := fmt.Sprintf("/api/v1/settings/labels/%s", name)
	err = cli.limiter.Wait(context.Background())
	if err != nil {
		return nil, err
	}
	resp, body, errs := cli.gorequest.Clone().Set("Authorization", "Bearer "+cli.token).Get(cli.url + apiPath).End()
	if errs != nil {
		return nil, errors.Wrap(getMergedError(errs), "failed getting Khulnasoft label")
	}
	if resp.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetLabel from %s%s, %v ", cli.url, apiPath, err)
			return nil, err
		}
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response Body")
			return nil, err
		}
		var errorResponse ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			log.Printf("Failed to Unmarshal response Body to ErrorResponse. Body: %v. error: %v", string(body), err)
			return nil, err
		}
		return nil, fmt.Errorf("failed getting Khulnasoft label. status: %v. error message: %v", resp.Status, errorResponse.Message)
	}
	if response.Name == "" {
		return nil, fmt.Errorf("khulnasoft label: %s not found 404", name)
	}
	return &response, err
}

// GetKhulnasoftLabels - get a list of khulnasoft labels
func (cli *Client) GetKhulnasoftLabels() (*KhulnasoftLabels, error) {
	var err error
	var response KhulnasoftLabels
	apiPath := fmt.Sprintf("/api/v2/settings/labels")
	err = cli.limiter.Wait(context.Background())
	if err != nil {
		return nil, err
	}
	resp, body, errs := cli.gorequest.Clone().Set("Authorization", "Bearer "+cli.token).Get(cli.url + apiPath).End()
	if errs != nil {
		return nil, errors.Wrap(getMergedError(errs), "failed getting Khulnasoft labels")
	}
	if resp.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetKhulnasoftLabels from %s%s, %v ", cli.url, apiPath, err)
			return nil, err
		}
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response Body")
			return nil, err
		}
		var errorResponse ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			log.Printf("Failed to Unmarshal response Body to ErrorResponse. Body: %v. error: %v", string(body), err)
			return nil, err
		}
		return nil, fmt.Errorf("failed getting Khulnasoft labels. status: %v. error message: %v", resp.Status, errorResponse.Message)
	}

	return &response, err

}

// CreateKhulnasoftLabel - creates single Khulnasoft Khulnasoft label
func (cli *Client) CreateKhulnasoftLabel(khulnasoftLabel *KhulnasoftLabel) error {
	payload, err := json.Marshal(khulnasoftLabel)
	if err != nil {
		return err
	}
	request := cli.gorequest
	apiPath := fmt.Sprintf("/api/v1/settings/labels")
	err = cli.limiter.Wait(context.Background())
	if err != nil {
		return err
	}
	resp, _, errs := request.Clone().Set("Authorization", "Bearer "+cli.token).Post(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		return errors.Wrap(getMergedError(errs), "failed creating Khulnasoft label.")
	}
	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response Body")
			return err
		}
		var errorResponse ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			log.Printf("Failed to Unmarshal response Body to ErrorResponse. Body: %v. error: %v", string(body), err)
			return err
		}
		return fmt.Errorf("failed creating Khulnasoft label. status: %v. error message: %v", resp.Status, errorResponse.Message)
	}
	return nil
}

// UpdateKhulnasoftLabel updates an existing Khulnasoft label
func (cli *Client) UpdateKhulnasoftLabel(khulnasoftLabel *KhulnasoftLabel) error {
	payload, err := json.Marshal(khulnasoftLabel)
	if err != nil {
		return err
	}
	request := cli.gorequest
	apiPath := fmt.Sprintf("/api/v1/settings/labels/%s", khulnasoftLabel.Name)
	err = cli.limiter.Wait(context.Background())
	if err != nil {
		return err
	}
	resp, _, errs := request.Clone().Set("Authorization", "Bearer "+cli.token).Put(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		return errors.Wrap(getMergedError(errs), "failed modifying Khulnasoft label")
	}
	if resp.StatusCode != 201 && resp.StatusCode != 204 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response Body")
			return err
		}
		var errorResponse ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			log.Printf("Failed to Unmarshal response Body to ErrorResponse. Body: %v. error: %v", string(body), err)
			return err
		}
		return fmt.Errorf("failed modifying Khulnasoft label. status: %v. error message: %v", resp.Status, errorResponse.Message)
	}
	return nil
}

// DeleteKhulnasoftLabel removes a Khulnasoft label
func (cli *Client) DeleteKhulnasoftLabel(name string) error {
	request := cli.gorequest
	apiPath := fmt.Sprintf("/api/v1/settings/labels/%s", name)
	err := cli.limiter.Wait(context.Background())
	if err != nil {
		return err
	}
	resp, _, errs := request.Clone().Set("Authorization", "Bearer "+cli.token).Delete(cli.url + apiPath).End()
	if errs != nil {
		return errors.Wrap(getMergedError(errs), "failed deleting Khulnasoft label")
	}
	if resp.StatusCode != 204 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response Body")
			return err
		}
		var errorResponse ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			log.Printf("Failed to Unmarshal response Body to ErrorResponse. Body: %v. error: %v", string(body), err)
			return err
		}
		return fmt.Errorf("failed deleting Khulnasoft label, status: %v. error message: %v", resp.Status, errorResponse.Message)
	}
	return nil
}
