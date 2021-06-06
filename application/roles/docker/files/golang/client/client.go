package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Comment struct {
	UserID int64  `json:"userId"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

const (
	baseURL = "https://jsonplaceholder.typicode.com"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) newRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

func (c *Client) GetComments(ctx context.Context) (comments []Comment, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, baseURL+"/comments", nil)
	if err != nil {
		return comments, fmt.Errorf("comments: invalid request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return comments, fmt.Errorf("comments: failed to make HTTP request: %w", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&comments)
	if err != nil {
		return comments, fmt.Errorf("comments: failed to parse HTTP response: %w", err)
	}

	return comments, nil
}
