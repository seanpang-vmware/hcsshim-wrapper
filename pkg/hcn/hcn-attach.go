import (
	"fmt"
)

type HcnAttachCommand struct {
	fs   *flag.FlagSet
	name string
}

func (h *HcnAttachCommand) Name() string {
	return h.fs.Name()
}

func (h *HcnAttachCommand) Init(args []string) error {
	return h.fs.Parse(args)
}

func (h *HcnAttachCommand) Run() error {
	fmt.Println("Running hcn attach command ", g.name)
	return nil
}

