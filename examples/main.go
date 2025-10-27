package main

import (
	"log"

	local "github.com/OpenListTeam/OpenList/v4/drivers/local"
	opplugin "github.com/OpenListTeam/OpenList/v4/pkg/plugin"
	go_plugin "github.com/hashicorp/go-plugin"
)

// customImpl implements the opplugin.Plugin interface and returns the
// concrete driver instances that will run inside the plugin process.
type customImpl struct{}

func (p *customImpl) Info() opplugin.PluginInfo {
	return opplugin.PluginInfo{
		Name:        "local_test",
		Description: "local_test plugin for OpenList",
		Drivers:     []string{"local"},
	}
}

func main() {
	// Serve the plugin over net/rpc using the PluginPlugin glue.
	err := func() error {
		go_plugin.Serve(&go_plugin.ServeConfig{
			HandshakeConfig: opplugin.HandshakeConfig,
			Plugins: map[string]go_plugin.Plugin{
				"main": &opplugin.RPCMainPlugin{
					Impl: &customImpl{},
				},
				"local": &opplugin.RPCDriverPlugin{
					Impl: &local.Local{},
				},
			},
		})
		// go-plugin Serve never returns unless there is an error starting the server.
		return nil
	}()
	if err != nil {
		log.Fatalf("plugin serve failed: %v", err)
	}
}
