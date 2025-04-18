package main

import (
	"fmt"
	"github.com/9triver/cfn-client-go"
	"github.com/9triver/cfn/proto/data"
)

func main() {
	newClient, err := client.NewClient("localhost:8080")
	if err != nil {
		panic(err)
	}
	replay, err := newClient.DeployPyFunc(&data.AppendPyFunc{
		Name:          "",
		Params:        nil,
		Venv:          "",
		Requirements:  nil,
		PickledObject: nil,
		Language:      0,
		Resource:      nil,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Deployed PyFunc:", replay)
}
