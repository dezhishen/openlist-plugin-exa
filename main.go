package main

import (
	"log"

	"github.com/OpenListTeam/OpenList/v4/pkg/plugin"
)

func main() {
	ps, err := plugin.LoadPlugins("./plugins")
	if err != nil {
		panic(err)
	}
	for _, p := range ps {
		log.Printf("Loaded plugin: %s \n", p.Config().Name)
	}
}
