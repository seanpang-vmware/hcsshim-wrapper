package hcn

import (
	"encoding/json"
	"flag"
	"fmt"
)

type HcnAttachCommand struct {
	Fs           *flag.FlagSet
	EndpointName string
	EndpointId   string
	SandboxID    string
}

func (h *HcnAttachCommand) Name() string {
	return "HCNAttach"
}

func (h *HcnAttachCommand) Init(args []string) error {
	h.Fs.StringVar(&h.EndpointId, "epId", "", "HNS endpoint id")
	h.Fs.StringVar(&h.EndpointName, "epName", "", "HNS endpoint name")
	h.Fs.StringVar(&h.SandboxID, "sandbox", "", "Sandbox id for endpoint to attach")
	h.Fs.Parse(args)
	return nil
}

func (h *HcnAttachCommand) Run() (string, error) {
	err := attachEndpoint(h.EndpointId, h.SandboxID)
	if err != nil {
		fmt.Println("Error: Failed to attach endpoint", err)
		return "", err
	}
	fmt.Println(json.Marshal(h.EndpointId))
	return h.EndpointId, nil
}
