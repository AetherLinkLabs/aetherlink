package autodrive

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client is a client for the autodrive server
type Client struct {
	ctx      context.Context
	endpoint string
	token    string
}

// NewClient creates a new autodrive client
func NewClient(ctx context.Context, endpoint, token string) *Client {
	return &Client{
		ctx:      ctx,
		endpoint: endpoint,
		token:    token,
	}
}

func (c *Client) get(url string) (io.Reader, error) {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-auth-provider", "apikey")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resson, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("request failed with status %d", resp.StatusCode)
		}

		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, resson)
	}

	return resp.Body, nil
}

func (c *Client) post(url string, payload io.Reader, res any, contentType string) error {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodPost, url, payload)
	if err != nil {
		return err
	}

	req.Header.Set("x-auth-provider", "apikey")
	req.Header.Set("Content-Type", "application/json")
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resson, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("request failed with status %d", resp.StatusCode)
		}

		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, resson)
	}

	if res == nil {
		return nil
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(responseBody, res)
}
