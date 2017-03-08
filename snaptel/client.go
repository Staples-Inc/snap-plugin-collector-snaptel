/*
Copyright 2016 Staples, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package snaptel

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 5
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

	//client := &http.Client{}

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	if insecure {
		client.Transport = &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		}
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
	if err != nil && response == nil {
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
