package clients

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL     string
	AccessToken string
	httpClient  *http.Client
	logger      *slog.Logger
}

func NewClient(baseURL, accessToken string) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &Client{
		BaseURL:     baseURL,
		AccessToken: accessToken,
		httpClient:  &http.Client{Transport: tr},
		logger:      slog.Default().With("component", "jira-client"),
	}
}

func (c *Client) request(
	ctx context.Context,
	method string,
	endpoint *url.URL,
	body interface{},
	out interface{},
) error {
	c.logger.InfoContext(ctx, "Request", "url", endpoint.String())

	var requestBody io.Reader

	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}

		requestBody = bytes.NewBuffer(jsonData)

		c.logger.InfoContext(ctx, "Request", "body", string(jsonData))
	}

	req, requestErr := http.NewRequestWithContext(
		ctx,
		method,
		endpoint.String(),
		requestBody,
	)

	if requestErr != nil {
		return fmt.Errorf("failed to create request: %w", requestErr)
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, responseErr := c.httpClient.Do(req)
	if responseErr != nil {
		return fmt.Errorf("failed to perform request: %w", responseErr)
	}

	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			c.logger.ErrorContext(ctx, "Failed to close response body", "error", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d, body: %s", resp.StatusCode, string(body))
	}

	if out != nil {
		if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
			return fmt.Errorf("failed to decode response body: %w", err)
		}
	}

	return nil
}
