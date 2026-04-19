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

func (c *Client) CreateProject(project ProjectInput) (Project, error) {
	jsonData, err := json.Marshal(project)
	if err != nil {
		return Project{}, fmt.Errorf("Error serializing project: %w", err)
	}
	fullURL := c.BaseURL.JoinPath("projects").String()
	resp, err := c.doRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return Project{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return Project{}, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var created Project
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return Project{}, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return created, nil
}

func (c *Client) UpdateProject(projectID int, project ProjectInput) (Project, error) {
	jsonData, err := json.Marshal(project)
	if err != nil {
		return Project{}, fmt.Errorf("Error serializing project: %w", err)
	}
	fullURL := c.BaseURL.JoinPath("projects", strconv.Itoa(projectID)).String()
	resp, err := c.doRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return Project{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return Project{}, fmt.Errorf("project with ID %d not found", projectID)
		}
		return Project{}, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var updated Project
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return Project{}, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return updated, nil
}

func (c *Client) DeleteProject(projectID int) error {
	fullURL := c.BaseURL.JoinPath("projects", strconv.Itoa(projectID)).String()
	resp, err := c.doRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("project with ID %d not found", projectID)
		}
		return fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	return nil
}

func (c *Client) GetProject(projectID int) (Project, error) {
	fullURL := c.BaseURL.JoinPath("projects", strconv.Itoa(projectID)).String()
	resp, err := c.doRequest("GET", fullURL, nil)
	if err != nil {
		return Project{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return Project{}, fmt.Errorf("project with ID %d not found", projectID)
		}
		return Project{}, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var project Project
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return Project{}, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return project, nil
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

func (c *Client) ListProjectViews(projectID int) (ProjectViews, error) {
	fullURL := c.BaseURL.JoinPath("projects", strconv.Itoa(projectID), "views").String()
	resp, err := c.doRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("project with ID %d not found", projectID)
		}
		return nil, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var views ProjectViews
	if err := json.NewDecoder(resp.Body).Decode(&views); err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return views, nil
}

func (c *Client) ListKanbanBuckets(projectID, viewID int) (Buckets, error) {
	fullURL := c.BaseURL.JoinPath("projects", strconv.Itoa(projectID), "views", strconv.Itoa(viewID), "buckets").String()
	resp, err := c.doRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("view with ID %d not found", viewID)
		}
		return nil, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var buckets Buckets
	if err := json.NewDecoder(resp.Body).Decode(&buckets); err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return buckets, nil
}

func (c *Client) MoveTaskToBucket(projectID, viewID, bucketID, taskID int) error {
	body := map[string]int{"task_id": taskID}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("Error serializing request: %w", err)
	}
	fullURL := c.BaseURL.JoinPath(
		"projects", strconv.Itoa(projectID),
		"views", strconv.Itoa(viewID),
		"buckets", strconv.Itoa(bucketID),
		"tasks",
	).String()
	resp, err := c.doRequest("POST", fullURL, bytes.NewBuffer(jsonData))
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

func (c *Client) ListLabels() (Labels, error) {
	fullURL := c.BaseURL.JoinPath("labels").String()
	resp, err := c.doRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var labels Labels
	if err := json.NewDecoder(resp.Body).Decode(&labels); err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return labels, nil
}

func (c *Client) CreateLabel(input LabelInput) (Label, error) {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return Label{}, fmt.Errorf("Error serializing label: %w", err)
	}
	fullURL := c.BaseURL.JoinPath("labels").String()
	resp, err := c.doRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return Label{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return Label{}, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var label Label
	if err := json.NewDecoder(resp.Body).Decode(&label); err != nil {
		return Label{}, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return label, nil
}

func (c *Client) DeleteLabel(labelID int) error {
	fullURL := c.BaseURL.JoinPath("labels", strconv.Itoa(labelID)).String()
	resp, err := c.doRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("label with ID %d not found", labelID)
		}
		return fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	return nil
}

func (c *Client) AddLabelToTask(taskID, labelID int) (Label, error) {
	body := map[string]int{"label_id": labelID}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return Label{}, fmt.Errorf("Error serializing label: %w", err)
	}
	fullURL := c.BaseURL.JoinPath("tasks", strconv.Itoa(taskID), "labels").String()
	resp, err := c.doRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return Label{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return Label{}, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var label Label
	if err := json.NewDecoder(resp.Body).Decode(&label); err != nil {
		return Label{}, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return label, nil
}

func (c *Client) RemoveLabelFromTask(taskID, labelID int) error {
	fullURL := c.BaseURL.JoinPath("tasks", strconv.Itoa(taskID), "labels", strconv.Itoa(labelID)).String()
	resp, err := c.doRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("label with ID %d not found on task %d", labelID, taskID)
		}
		return fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	return nil
}

func (c *Client) ListTaskComments(taskID int) (TaskComments, error) {
	fullURL := c.BaseURL.JoinPath("tasks", strconv.Itoa(taskID), "comments").String()
	resp, err := c.doRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("task with ID %d not found", taskID)
		}
		return nil, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var comments TaskComments
	if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return comments, nil
}

func (c *Client) CreateTaskComment(taskID int, input TaskCommentInput) (TaskComment, error) {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return TaskComment{}, fmt.Errorf("Error serializing comment: %w", err)
	}
	fullURL := c.BaseURL.JoinPath("tasks", strconv.Itoa(taskID), "comments").String()
	resp, err := c.doRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return TaskComment{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return TaskComment{}, fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	var comment TaskComment
	if err := json.NewDecoder(resp.Body).Decode(&comment); err != nil {
		return TaskComment{}, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return comment, nil
}

func (c *Client) DeleteTaskComment(taskID, commentID int) error {
	fullURL := c.BaseURL.JoinPath("tasks", strconv.Itoa(taskID), "comments", strconv.Itoa(commentID)).String()
	resp, err := c.doRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("comment with ID %d not found", commentID)
		}
		return fmt.Errorf("Error API Vikunja: Code %d - Status: %v", resp.StatusCode, resp.Status)
	}
	return nil
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
