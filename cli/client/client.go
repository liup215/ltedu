package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client is the HTTP client for the edu backend API.
type Client struct {
	cfg *Config
	hc  *http.Client
}

// NewClient creates a new API client using the current configuration.
func NewClient() *Client {
	return &Client{
		cfg: LoadConfig(),
		hc:  &http.Client{},
	}
}

// NewClientWithConfig creates a new API client with the provided configuration.
func NewClientWithConfig(cfg *Config) *Client {
	return &Client{cfg: cfg, hc: &http.Client{}}
}

// APIResponse is the standard response envelope from the backend.
type APIResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

// post sends an authenticated POST request and returns the parsed APIResponse.
func (c *Client) post(path string, body interface{}) (*APIResponse, error) {
	if c.cfg.BaseURL == "" {
		return nil, fmt.Errorf("base URL is not configured; set %s or run 'edu-cli config set-url'", EnvBaseURL)
	}
	if c.cfg.Token == "" {
		return nil, fmt.Errorf("authentication token not configured; run 'edu-cli config set-token <token>' or set %s", EnvToken)
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.cfg.BaseURL+"/api"+path, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.cfg.Token)

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(raw, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &apiResp, nil
}

// PostAndDecode sends a POST request and decodes the data field into result.
func (c *Client) PostAndDecode(path string, body interface{}, result interface{}) error {
	apiResp, err := c.post(path, body)
	if err != nil {
		return err
	}
	if apiResp.Code != 0 {
		return fmt.Errorf("API error: %s", apiResp.Message)
	}
	if result == nil || len(apiResp.Data) == 0 {
		return nil
	}
	return json.Unmarshal(apiResp.Data, result)
}

// get sends an authenticated GET request and returns the parsed APIResponse.
func (c *Client) get(path string) (*APIResponse, error) {
	if c.cfg.BaseURL == "" {
		return nil, fmt.Errorf("base URL is not configured; set %s or run 'edu-cli config set-url'", EnvBaseURL)
	}
	if c.cfg.Token == "" {
		return nil, fmt.Errorf("authentication token not configured; run 'edu-cli config set-token <token>' or set %s", EnvToken)
	}

	req, err := http.NewRequest(http.MethodGet, c.cfg.BaseURL+"/api"+path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.cfg.Token)

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(raw, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &apiResp, nil
}

// GetAndDecode sends a GET request and decodes the data field into result.
func (c *Client) GetAndDecode(path string, result interface{}) error {
	apiResp, err := c.get(path)
	if err != nil {
		return err
	}
	if apiResp.Code != 0 {
		return fmt.Errorf("API error: %s", apiResp.Message)
	}
	if result == nil || len(apiResp.Data) == 0 {
		return nil
	}
	return json.Unmarshal(apiResp.Data, result)
}

// SetToken updates the token in the loaded config and persists it.
func (c *Client) SetToken(token string) error {
	c.cfg.Token = token
	return SaveConfig(c.cfg)
}

// GetConfig returns the current configuration (read-only).
func (c *Client) GetConfig() *Config {
	return c.cfg
}
