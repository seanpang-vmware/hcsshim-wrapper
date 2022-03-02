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
	Sandbox      string
}

func (h *HcnAttachCommand) Name() string {
	return "HCNAttach"
}

func (h *HcnAttachCommand) Init(args []string) error {
	h.Fs.StringVar(&h.EndpointId, "epId", "", "HNS endpoint id")
	h.Fs.StringVar(&h.EndpointName, "epName", "", "HNS endpoint name")
	h.Fs.StringVar(&h.Sandbox, "sandbox", "", "Contianer netns for endpoint to attach")
	h.Fs.Parse(args)
	return nil
}

func (h *HcnAttachCommand) Run() (string, error) {
	err := attachEndpoint(h.EndpointId, h.Sandbox)
	if err != nil {
		return "", err
	}
	fmt.Println(json.Marshal(h.EndpointId))
	return h.EndpointId, nil
}
