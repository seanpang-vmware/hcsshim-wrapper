package hcn

import (
	"flag"
	"fmt"

	"github.com/Microsoft/hcsshim/hcn"
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

func (h *HcnAttachCommand) Run() error {
	fmt.Println("Running hcn attach command ", h.Name())
	fmt.Printf("Read epid %s, epname %s, sbid %s", h.EndpointId, h.EndpointName, h.SandboxID)

	attachedEpIds, err := hcn.GetNamespaceEndpointIds(sandbox)
	if err != nil {
		return nil, err
	}
	return nil
}
