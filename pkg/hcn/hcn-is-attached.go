package hcn

import (
	"flag"
	"strconv"
)

type HcnIsAttachedCommand struct {
	Fs         *flag.FlagSet
	EndpointId string
	Sandbox    string
}

func (h *HcnIsAttachedCommand) Name() string {
	return "HCNIsAttached"
}

func (h *HcnIsAttachedCommand) Init(args []string) error {
	h.Fs.StringVar(&h.EndpointId, "epId", "", "HNS endpoint id")
	h.Fs.StringVar(&h.Sandbox, "sandbox", "", "Container netns for endpoint to attach")
	h.Fs.Parse(args)
	return nil
}

func (h *HcnIsAttachedCommand) Run() (string, error) {
	attached, err := isEndpointAttached(h.EndpointId, h.Sandbox)
	if err != nil {
		return strconv.FormatBool(false), err
	}
	return strconv.FormatBool(attached), nil
}
