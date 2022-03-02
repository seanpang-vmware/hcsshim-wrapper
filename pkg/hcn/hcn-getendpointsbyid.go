package hcn

import (
	"encoding/json"
	"flag"
)

type HcnGetEndpointsBySandboxIdCommand struct {
	Fs      *flag.FlagSet
	Sandbox string
}

func (h *HcnGetEndpointsBySandboxIdCommand) Name() string {
	return "HCNGetNamespaceEndpoints"
}

func (h *HcnGetEndpointsBySandboxIdCommand) Init(args []string) error {
	h.Fs.StringVar(&h.Sandbox, "sandbox", "", "Container netns for endpoint to attach")
	h.Fs.Parse(args)
	return nil
}

func (h *HcnGetEndpointsBySandboxIdCommand) Run() (string, error) {
	attachedEpIds, err := GetEndpointsByID(h.Sandbox)
	if err != nil {
		return "", err
	}

	if len(attachedEpIds) > 0 {
		output, err := json.Marshal(attachedEpIds)
		if err != nil {
			return "", err
		}
		return string(output), err
	}

	return "", nil
}
