package clients

import (
	"bytes"
	"context"
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
	return &Client{
		BaseURL:     baseURL,
		AccessToken: accessToken,
		httpClient:  &http.Client{},
		logger:      slog.Default().With("component", "jira-client"),
	}
}

func (c *Client) Me(
	ctx context.Context,
) (*MeAPIResponse, error) {
	endpoint := fmt.Sprintf("%s/myself", c.BaseURL)

	reqURL, parseErr := url.Parse(endpoint)
	if parseErr != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", parseErr)
	}

	var responseData MeAPIResponse
	err := c.request(ctx, http.MethodGet, reqURL, nil, &responseData)

	if err != nil {
		c.logger.ErrorContext(ctx, "request failed", slog.Any("error", err))

		return nil, err
	}

	return &responseData, nil
}

func (c *Client) GetProjects(
	ctx context.Context,
) (*GetProjectsAPIResponse, error) {
	endpoint := fmt.Sprintf("%s/project", c.BaseURL)

	reqURL, parseErr := url.Parse(endpoint)
	if parseErr != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", parseErr)
	}

	var entries []Project
	err := c.request(ctx, http.MethodGet, reqURL, nil, &entries)

	if err != nil {
		c.logger.ErrorContext(ctx, "request failed", slog.Any("error", err))

		return nil, err
	}

	return &GetProjectsAPIResponse{Data: entries}, nil
}

func (c *Client) CreateProject(
	ctx context.Context,
	input CreateProjectInput,
) (*CreateProjectAPIResponse, error) {
	endpoint := fmt.Sprintf("%s/project", c.BaseURL)

	project := &CreateProject{
		LeadAccountID:      input.AccountID,
		AssigneeType:       input.AssigneeType,
		Name:               input.Name,
		Description:        input.Description,
		Key:                input.TaskPrefixKey,
		ProjectTypeKey:     input.ProjectTypeKey,
		ProjectTemplateKey: input.ProjectTemplateKey,
		PermissionScheme:   0,
	}

	reqURL, parseErr := url.Parse(endpoint)
	if parseErr != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", parseErr)
	}

	var responseData CreateProjectAPIResponse
	err := c.request(ctx, http.MethodPost, reqURL, &project, &responseData)

	if err != nil {
		c.logger.ErrorContext(ctx, "request failed", slog.Any("error", err))

		return nil, err
	}

	return &responseData, nil
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

	req.Header.Set("Authorization", "Basic "+c.AccessToken)
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

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated &&
		resp.StatusCode != http.StatusNoContent {
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
