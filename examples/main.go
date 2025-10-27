package main

import (
	"log"

	_local "github.com/OpenListTeam/OpenList/v4/drivers/local"
	opplugin "github.com/OpenListTeam/OpenList/v4/pkg/plugin"
	go_plugin "github.com/hashicorp/go-plugin"
)

// pluginImpl implements the opplugin.Plugin interface and returns the
// concrete driver instances that will run inside the plugin process.
type pluginImpl struct{}

func (p *pluginImpl) Info() opplugin.PluginInfo {
	return opplugin.PluginInfo{
		Name:        "local_test",
		Description: "local_test plugin for OpenList",
		Protocol:    "netrpc",
	}
}

func (p *pluginImpl) Drivers() opplugin.PluginDriver {
	// Return driver instances. The host will receive proxies for these drivers.
	// Here we return a single default Open115 driver instance.
	return opplugin.PluginDriver{&_local.Local{}}
}

func main() {
	// Serve the plugin over net/rpc using the PluginPlugin glue.
	err := func() error {
		go_plugin.Serve(&go_plugin.ServeConfig{
			HandshakeConfig: opplugin.HandshakeConfig,
			Plugins: map[string]go_plugin.Plugin{
				"main": &opplugin.PluginNetRpcPlugin{Impl: &pluginImpl{}},
			},
		})
		// go-plugin Serve never returns unless there is an error starting the server.
		return nil
	}()
	if err != nil {
		log.Fatalf("plugin serve failed: %v", err)
	}
}
