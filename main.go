package main

import (
	"github.com/Staples-Inc/snap-plugin-collector-snaptel/snaptel"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

const (
	pluginName    = "snaptel"
	pluginVersion = 1
)

func main() {
	plugin.StartCollector(snaptel.CollectorSnaptel{}, pluginName, pluginVersion, plugin.RoutingStrategy(plugin.StickyRouter))
}
