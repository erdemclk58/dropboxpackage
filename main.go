package main

import (
	"robomotion-go-plugin/nodes"

	"bitbucket.org/mosteknoloji/robomotion-go-lib/runtime"
)

func main() {
	runtime.RegisterNodes(
		&nodes.Connect{},
		&nodes.Copy{},
		&nodes.CreateFolder{},
		&nodes.Delete{},
		&nodes.Disconnect{},
		&nodes.DownloadFile{},
		&nodes.Move{},
		&nodes.Stat{},
		&nodes.WriteFile{},
	)

	runtime.Start()
}
