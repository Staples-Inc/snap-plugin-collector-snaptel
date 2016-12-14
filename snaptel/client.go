package snaptel

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type SnaptelClient interface {
	GetTaskMetrics() ([]TaskData, error)
}

type snaptelClient struct {
	client *http.Client
	url    string
	apiKey string
}

func NewSnaptelClient(address string, https bool, insecure bool, key string) *snaptelClient {
	var scheme string
	if https {
		scheme = "https://"
	} else {
		scheme = "http://"
	}
	url := scheme + address

	client := &http.Client{}
	if insecure {
		client.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	}
	return &snaptelClient{
		client: client,
		apiKey: key,
		url:    url,
	}
}

func (s snaptelClient) GetTaskMetrics() ([]TaskData, error) {
	url := s.url
	req, err := http.NewRequest("GET", url+"/v1/tasks", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth("snap", s.apiKey)
	response, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Bad response code on request to snaptel api: %v", response.StatusCode)
	}
	decoder := json.NewDecoder(response.Body)
	tasks := &TaskResponse{}
	err = decoder.Decode(tasks)
	if err != nil {
		return nil, err
	}
	return tasks.Body.ScheduledTasks, nil
}
