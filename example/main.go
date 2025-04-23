package main

import (
	"fmt"
	"github.com/9triver/cfn-client-go"
	"github.com/9triver/cfn/proto/data"
)

func main() {
	newClient, err := client.NewClient("localhost:9090")
	//newClient, err := client.NewClient("localhost:8667")
	if err != nil {
		panic(err)
	}
	replay, err := newClient.DeployPyFunc(&data.AppendPyFunc{
		Name:          "",
		Params:        nil,
		Venv:          "",
		Requirements:  []string{"pandas==2.2.1", "numpy==1.26.4"},
		PickledObject: nil,
		Language:      0,
		Resource: &data.Resource{
			CPU: &data.CPU{
				Cores: "5000m",
			},
			Memory: &data.Memory{
				Capacity: "2Gi",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Deployed PyFunc:", replay)
	defer newClient.Close()
}
