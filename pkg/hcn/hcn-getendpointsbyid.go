package hcn

import (
	"flag"
	"fmt"

	"github.com/Microsoft/hcsshim/hcn"
)

type HcnGetEndpointsBySandboxIdCommand struct {
	Fs        *flag.FlagSet
	SandboxID string
}

func (h *HcnGetEndpointsBySandboxIdCommand) Name() string {
	return "HCNGetNamespaceEndpoints"
}

func (h *HcnGetEndpointsBySandboxIdCommand) Init(args []string) error {
	h.Fs.StringVar(&h.SandboxID, "sandbox", "", "Sandbox id for endpoint to attach")
	h.Fs.Parse(args)
	fmt.Println("init sbid ", h.SandboxID)
	return nil
}

func (h *HcnGetEndpointsBySandboxIdCommand) Run() error {
	fmt.Println("Running hcn command ", h.Name())
	fmt.Printf("Read sbid %s", h.SandboxID)

	attachedEpIds, err := hcn.GetNamespaceEndpointIds(h.SandboxID)
	if err != nil {
		return nil, err
	}
	return attachedEpIds
}
