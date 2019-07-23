package main

import (
	"github.com/imrancluster/th-common-payment/cmd"
)

const version = "v1.0.0"

func main() {
	cmd.Execute()
}

// main -> cmd.Executes()->RootCmd->server.go->init()->RootCmd.AddCommand->Run:serve->serve()
