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
	fmt.Printf("init epid %s, epname %s, sbid %s", h.EndpointId, h.EndpointName, h.SandboxID)
	return nil
}

func (h *HcnAttachCommand) Run() (string, error) {
	fmt.Println("Running hcn attach command ", h.Name())
	fmt.Printf("Read epid %s, epname %s, sbid %s", h.EndpointId, h.EndpointName, h.SandboxID)

	err := attachEndpoint(h.EndpointId, h.SandboxID)
	if err != nil {
		fmt.Println("Failed to attach endpoint")
		return "", err
	}
	fmt.Println(json.Marshal(h.EndpointId))
	return h.EndpointId, nil
}
