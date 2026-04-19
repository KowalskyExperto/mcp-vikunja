package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func (c *Client) doRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error executing request: %w", err)
	}
	return resp, nil
}

func (c *Client) GetProjects() (Projects, error) {
	fullURL := c.BaseURL.JoinPath("projects").String()
	resp, err := c.doRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var projects Projects
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return projects, nil
}

func (c *Client) GetTasks() (Tasks, error) {
	fullURL := c.BaseURL.JoinPath("tasks").String()
	resp, err := c.doRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var tasks Tasks
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return tasks, nil
}

func (c *Client) GetTasksByProject(projectID int) (Tasks, error) {
	fullURL := c.BaseURL.JoinPath("projects", strconv.Itoa(projectID), "tasks").String()
	resp, err := c.doRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
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
		return nil, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var tasks Tasks
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return tasks, nil
}

func (c *Client) CreateTask(projectID int, task TaskInput) (Task, error) {
	jsonData, err := json.Marshal(task)
	if err != nil {
		return Task{}, fmt.Errorf("Error serializing task: %w", err)
	}
	fullURL := c.BaseURL.JoinPath("projects", strconv.Itoa(projectID), "tasks").String()
	resp, err := c.doRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return Task{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return Task{}, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var createdTask Task
	if err := json.NewDecoder(resp.Body).Decode(&createdTask); err != nil {
		return Task{}, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return createdTask, nil
}

func (c *Client) GetTask(taskID int) (TaskDetail, error) {
	fullURL := c.BaseURL.JoinPath("tasks", strconv.Itoa(taskID)).String()
	resp, err := c.doRequest("GET", fullURL, nil)
	if err != nil {
		return TaskDetail{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return TaskDetail{}, fmt.Errorf("task with ID %d not found", taskID)
		}
		return TaskDetail{}, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var task TaskDetail
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return TaskDetail{}, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return task, nil
}

func (c *Client) UpdateTask(taskID int, task TaskUpdate) (Task, error) {
	jsonData, err := json.Marshal(task)
	if err != nil {
		return Task{}, fmt.Errorf("Error serializing task: %w", err)
	}
	fullURL := c.BaseURL.JoinPath("tasks", strconv.Itoa(taskID)).String()
	resp, err := c.doRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return Task{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Task{}, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var updatedTask Task
	if err := json.NewDecoder(resp.Body).Decode(&updatedTask); err != nil {
		return Task{}, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return updatedTask, nil
}

func (c *Client) DeleteTask(taskID int) error {
	fullURL := c.BaseURL.JoinPath("tasks", strconv.Itoa(taskID)).String()
	resp, err := c.doRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("task with ID %d not found", taskID)
		}
		return fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	return nil
}
