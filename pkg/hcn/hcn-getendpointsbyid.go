package hcn

import (
	"encoding/json"
	"flag"
	"fmt"

	utils "github.com/seanpang-vmware/hcsshim-wrapper/pkg/hcn/utils.go"
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

func (h *HcnGetEndpointsBySandboxIdCommand) Run() (string, error) {
	fmt.Println("Running hcn command ", h.Name())
	fmt.Printf("Read sbid %s", h.SandboxID)

	attachedEpIds, err := utils.GetEndpointByID(h.SandboxID)
	if err != nil {
		output, err := json.Marshal(attachedEpIds)
		if err != nil {
			return string(output), nil
		}
		return "", err
	}
	return "", err
}
