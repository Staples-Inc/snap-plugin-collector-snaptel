package snaptel

import "github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

// CollectorSnaptel for snaptel
type CollectorSnaptel struct {
	client      SnaptelClient
	initialized bool
}

//CollectMetrics will be called by Snap when a task that collects one of the metrics returned from this plugins
func (c CollectorSnaptel) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	if c.initialized == false {
		address, err := mts[0].Config.GetString("address")
		if err != nil {
			return nil, err
		}
		key, err := mts[0].Config.GetString("api_key")
		if err != nil {
			return nil, err
		}
		scheme, err := mts[0].Config.GetBool("https")
		if err != nil {
			return nil, err
		}
		insecure, err := mts[0].Config.GetBool("insecure")
		if err != nil {
			return nil, err
		}
		c.client = NewSnaptelClient(address, scheme, insecure, key)
		c.initialized = true
	}
	taskMetrics, err := c.client.GetTaskMetrics()
	if err != nil {
		return nil, err
	}
	statusMap := map[string]int{
		"Running":  0,
		"Stopped":  0,
		"Disabled": 0,
	}
	metrics := []plugin.Metric{}
	for t := range taskMetrics {
		statusMap[taskMetrics[t].State]++
		metrics = append(metrics, taskMetrics[t].returnMetrics()...)
	}
	for status, value := range statusMap {
		metrics = append(metrics, plugin.Metric{
			Namespace: plugin.Namespace{
				plugin.NamespaceElement{Value: "staples"},
				plugin.NamespaceElement{Value: "snaptel"},
				plugin.NamespaceElement{Value: "tasks"},
				plugin.NamespaceElement{Value: status},
			},
			Data: value,
		})
	}
	return interpolate(mts, metrics)
}

func interpolate(requestedMts []plugin.Metric, availableMts []plugin.Metric) ([]plugin.Metric, error) {
	mts := []plugin.Metric{}
	for _, metric := range requestedMts {
		if metric.Namespace.Strings()[2] == "task" {
			if metric.Namespace.Strings()[3] == "*" {
				for _, aMetric := range availableMts {
					if aMetric.Namespace.Strings()[2] == "task" {
						if aMetric.Namespace.Strings()[4] == metric.Namespace.Strings()[4] {
							mts = append(mts, plugin.Metric{
								Namespace: aMetric.Namespace,
								Data:      aMetric.Data,
								Config:    metric.Config,
								Tags:      metric.Tags,
							})
						}
					}
				}
			} else {
				for _, aMetric := range availableMts {
					if aMetric.Namespace.Strings()[3] == metric.Namespace.Strings()[3] {
						if aMetric.Namespace.Strings()[4] == metric.Namespace.Strings()[4] {
							mts = append(mts, plugin.Metric{
								Namespace: aMetric.Namespace,
								Data:      aMetric.Data,
								Config:    metric.Config,
								Tags:      metric.Tags,
							})
						}
					}
				}
			}
		}
		if metric.Namespace.Strings()[2] == "tasks" {
			for _, aMetric := range availableMts {
				if aMetric.Namespace.Strings()[3] == metric.Namespace.Strings()[3] {
					mts = append(mts, plugin.Metric{
						Namespace: aMetric.Namespace,
						Data:      aMetric.Data,
						Config:    metric.Config,
						Tags:      metric.Tags,
					})
				}
			}
		}
	}
	return mts, nil
}

//GetMetricTypes will be called when your plugin is loaded in order to populate the metric catalog
func (CollectorSnaptel) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	return getAvailableMetrics(), nil
}

// GetConfigPolicy returns the configPolicy
func (CollectorSnaptel) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	policy.AddNewStringRule([]string{"staples", "snaptel"}, "address", true, plugin.SetDefaultString("localhost:8181"))
	policy.AddNewStringRule([]string{"staples", "snaptel"}, "api_key", false, plugin.SetDefaultString(""))
	policy.AddNewBoolRule([]string{"staples", "snaptel"}, "https", true, plugin.SetDefaultBool(false))
	policy.AddNewBoolRule([]string{"staples", "snaptel"}, "insecure", true, plugin.SetDefaultBool(false))
	return *policy, nil
}
