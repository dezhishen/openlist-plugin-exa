package main

import (
	"github.com/OpenListTeam/OpenList/v4/pkg/plugin"
)

func main() {
	ps, err := plugin.LoadPlugins("./plugins")
	if err != nil {
		panic(err)
	}
	for _, p := range ps {
		for _, driver := range p.Drivers() {
			println("  Driver:", driver.Config().Name)
		}
	}
}
