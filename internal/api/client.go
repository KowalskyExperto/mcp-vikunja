package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	BaseURL    *url.URL
	Token      string
	HTTPClient *http.Client
}

func NewClient(rawURL, token string) (*Client, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("Invalid API URL: %w", err)
	}
	return &Client{

		BaseURL: parsedURL,
		Token:   token,
		HTTPClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}, nil
}

func (c *Client) GetProjects() (Projects, error) {
	fullURL := c.BaseURL.JoinPath("projects").String()
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error executing request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error API Vikunja: Code %d", resp.StatusCode)
	}
	var projects Projects
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return projects, nil
}

func (c *Client) GetTasks() (Tasks, error) {
	fullURL := c.BaseURL.JoinPath("tasks").String()
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error executing request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error API Vikunja: Code %d", resp.StatusCode)
	}
	var tasks Tasks
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return tasks, nil
}

func (c *Client) GetTasksByProject(projectID int) (Tasks, error) {
	fullURL := c.BaseURL.JoinPath("projects/", strconv.Itoa(projectID), "/tasks").String()
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error executing request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error API Vikunja: Code %d", resp.StatusCode)
	}
	var tasks Tasks
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return tasks, nil
}

func (c *Client) SearchTasks(search string) (Tasks, error) {
	fullURL := c.BaseURL.JoinPath("tasks").String()
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %w", err)
	}
	if search != "" {
		q := req.URL.Query()
		q.Set("s", search)
		req.URL.RawQuery = q.Encode()
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error executing request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error API Vikunja: Code %d", resp.StatusCode)
	}
	var tasks Tasks
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return tasks, nil
}
