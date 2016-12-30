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
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/jarcoal/httpmock"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	exampleTasks = `{
  "meta": {
    "code": 200,
    "message": "Scheduled tasks retrieved",
    "type": "scheduled_task_list_returned",
    "version": 1
  },
  "body": {
    "ScheduledTasks": [
      {
        "id": "f573affa-9326-44a8-a64c-7a0d803d5121",
        "name": "exampleTask1",
        "deadline": "5s",
        "creation_timestamp": 1448004968,
        "last_run_timestamp": 1448004989,
        "hit_count": 10,
        "fail_count": 15,
        "task_state": "Running"
      },
      {
        "id": "f573affa-9326-44a8-a64c-7a0d803d5121",
        "name": "exampleTask2",
        "deadline": "5s",
        "creation_timestamp": 1448004968,
        "last_run_timestamp": 1448004989,
        "hit_count": 1000,
        "task_state": "Stopped"
      },
      {
        "id": "f573affa-9326-44a8-a64c-7a0d803d5121",
        "name": "exampleTask3",
        "deadline": "5s",
        "creation_timestamp": 1448004968,
        "last_run_timestamp": 1448004989,
        "hit_count": 0,
        "fail_count": 15,
        "task_state": "Disabled"
      }
    ]
  }
}`

	exampleTaskValues = map[string]interface{}{
		"staples.snaptel.task.exampleTask1.state":      "Running",
		"staples.snaptel.task.exampleTask2.state":      "Stopped",
		"staples.snaptel.task.exampleTask3.state":      "Disabled",
		"staples.snaptel.task.exampleTask1.statecode":  1,
		"staples.snaptel.task.exampleTask2.statecode":  2,
		"staples.snaptel.task.exampleTask3.statecode":  3,
		"staples.snaptel.task.exampleTask1.hit_count":  10,
		"staples.snaptel.task.exampleTask2.hit_count":  1000,
		"staples.snaptel.task.exampleTask3.hit_count":  0,
		"staples.snaptel.task.exampleTask1.fail_count": 15,
		"staples.snaptel.task.exampleTask2.fail_count": 0,
		"staples.snaptel.task.exampleTask3.fail_count": 15,
		"staples.snaptel.tasks.Running":                1,
		"staples.snaptel.tasks.Disabled":               1,
		"staples.snaptel.tasks.Stopped":                1,
	}
)

func TestSnaptelPlugin(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "http://localhost:8181/v1/tasks",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, exampleTasks)
			return resp, nil
		},
	)
	httpmock.RegisterResponder("GET", "http://localhost:8182/v1/tasks",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, "")
			return resp, nil
		},
	)
	Convey("Get Metrics Types", t, func() {
		SnaptelCollector := CollectorSnaptel{}
		config := plugin.Config{}
		Convey("So should return 7 available metrics", func() {
			metrics, err := SnaptelCollector.GetMetricTypes(config)
			So(len(metrics), ShouldResemble, 7)
			So(err, ShouldBeNil)
		})
	})
	Convey("Collect Metrics from Mock", t, func() {
		SnaptelCollector := CollectorSnaptel{}
		config := plugin.Config{
			"address":  "localhost:8181",
			"api_key":  "",
			"https":    false,
			"insecure": false,
		}
		metrics, _ := SnaptelCollector.GetMetricTypes(config)
		for m := range metrics {
			metrics[m].Config = config
		}
		Convey("Snaptel should return collected metrics with the correct values", func() {
			collectedMetrics, err := SnaptelCollector.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			for _, m := range collectedMetrics {
				fmt.Printf("\n      Metric: %v", strings.Join(m.Namespace.Strings(), "."))
				val, metricExists := exampleTaskValues[strings.Join(m.Namespace.Strings(), ".")]
				So(metricExists, ShouldBeTrue)
				So(m.Data, ShouldEqual, val)
			}
		})
	})
	Convey("Collect Metrics from bad endpoint", t, func() {
		SnaptelCollector := CollectorSnaptel{}
		config := plugin.Config{
			"address":  "localhost:8182",
			"api_key":  "",
			"https":    false,
			"insecure": false,
		}
		metrics, _ := SnaptelCollector.GetMetricTypes(config)
		for m := range metrics {
			metrics[m].Config = config
		}
		Convey("Snaptel should return an error", func() {
			collectedMetrics, err := SnaptelCollector.CollectMetrics(metrics)
			So(err, ShouldNotBeNil)
			So(collectedMetrics, ShouldBeNil)
		})
	})
}
