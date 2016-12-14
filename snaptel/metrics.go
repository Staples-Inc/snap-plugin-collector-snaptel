package snaptel

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

var (
	stateCodeMap = map[string]int{
		"Running":  1,
		"Stopped":  2,
		"Disabled": 3,
	}
)

type TaskResponse struct {
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Type    string `json:"type"`
		Version int    `json:"version"`
	} `json:"meta"`
	Body struct {
		ScheduledTasks []TaskData `json:"ScheduledTasks"`
	} `json:"body"`
}

// TaskData is how tasks are deserialized when parsed from the snap api endpoint
type TaskData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	State     string `json:"task_state"`
	HitCount  int    `json:"hit_count"`
	FailCount int    `json:"fail_count"`
}

func (t TaskData) returnMetrics() []plugin.Metric {
	metrics := []plugin.Metric{
		plugin.Metric{
			Namespace: plugin.Namespace{
				plugin.NamespaceElement{Value: "staples"},
				plugin.NamespaceElement{Value: "snaptel"},
				plugin.NamespaceElement{Value: "task"},
				plugin.NamespaceElement{Value: t.Name, Name: "task_name", Description: "Name of the running task"},
				plugin.NamespaceElement{Value: "state"},
			},
			Data: t.State,
		},
		plugin.Metric{
			Namespace: plugin.Namespace{
				plugin.NamespaceElement{Value: "staples"},
				plugin.NamespaceElement{Value: "snaptel"},
				plugin.NamespaceElement{Value: "task"},
				plugin.NamespaceElement{Value: t.Name, Name: "task_name", Description: "Name of the running task"},
				plugin.NamespaceElement{Value: "statecode"},
			},
			Data: t.getStatusCode(),
		},
		plugin.Metric{
			Namespace: plugin.Namespace{
				plugin.NamespaceElement{Value: "staples"},
				plugin.NamespaceElement{Value: "snaptel"},
				plugin.NamespaceElement{Value: "task"},
				plugin.NamespaceElement{Value: t.Name, Name: "task_name", Description: "Name of the running task"},
				plugin.NamespaceElement{Value: "hit_count"},
			},
			Data: t.HitCount,
		},
		plugin.Metric{
			Namespace: plugin.Namespace{
				plugin.NamespaceElement{Value: "staples"},
				plugin.NamespaceElement{Value: "snaptel"},
				plugin.NamespaceElement{Value: "task"},
				plugin.NamespaceElement{Value: t.Name, Name: "task_name", Description: "Name of the running task"},
				plugin.NamespaceElement{Value: "fail_count"},
			},
			Data: t.FailCount,
		},
	}
	return metrics
}

func (t TaskData) getStatusCode() int {
	code, _ := stateCodeMap[t.State]
	return code
}

func getAvailableMetrics() []plugin.Metric {
	return []plugin.Metric{
		plugin.Metric{
			Namespace: plugin.Namespace{
				plugin.NamespaceElement{Value: "staples"},
				plugin.NamespaceElement{Value: "snaptel"},
				plugin.NamespaceElement{Value: "task"},
				plugin.NamespaceElement{Value: "*", Name: "task_name", Description: "Name of the running task"},
				plugin.NamespaceElement{Value: "state"},
			},
		},
		plugin.Metric{
			Namespace: plugin.Namespace{
				plugin.NamespaceElement{Value: "staples"},
				plugin.NamespaceElement{Value: "snaptel"},
				plugin.NamespaceElement{Value: "task"},
				plugin.NamespaceElement{Value: "*", Name: "task_name", Description: "Name of the running task"},
				plugin.NamespaceElement{Value: "statecode"}, // 1 :: running, 2 :: stopped, 3 :: disabled
			},
		},
		plugin.Metric{
			Namespace: plugin.Namespace{
				plugin.NamespaceElement{Value: "staples"},
				plugin.NamespaceElement{Value: "snaptel"},
				plugin.NamespaceElement{Value: "task"},
				plugin.NamespaceElement{Value: "*", Name: "task_name", Description: "Name of the running task"},
				plugin.NamespaceElement{Value: "hit_count"},
			},
		},
		plugin.Metric{
			Namespace: plugin.Namespace{
				plugin.NamespaceElement{Value: "staples"},
				plugin.NamespaceElement{Value: "snaptel"},
				plugin.NamespaceElement{Value: "task"},
				plugin.NamespaceElement{Value: "*", Name: "task_name", Description: "Name of the running task"},
				plugin.NamespaceElement{Value: "fail_count"},
			},
		},
		plugin.Metric{
			Namespace: plugin.Namespace{
				plugin.NamespaceElement{Value: "staples"},
				plugin.NamespaceElement{Value: "snaptel"},
				plugin.NamespaceElement{Value: "tasks"},
				plugin.NamespaceElement{Value: "Running"},
			},
		},
		plugin.Metric{
			Namespace: plugin.Namespace{
				plugin.NamespaceElement{Value: "staples"},
				plugin.NamespaceElement{Value: "snaptel"},
				plugin.NamespaceElement{Value: "tasks"},
				plugin.NamespaceElement{Value: "Disabled"},
			},
		},
		plugin.Metric{
			Namespace: plugin.Namespace{
				plugin.NamespaceElement{Value: "staples"},
				plugin.NamespaceElement{Value: "snaptel"},
				plugin.NamespaceElement{Value: "tasks"},
				plugin.NamespaceElement{Value: "Stopped"},
			},
		},
	}
}
