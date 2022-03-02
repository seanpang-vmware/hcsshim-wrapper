package hcn

import (
	"errors"

	"github.com/Microsoft/hcsshim/hcn"
)

func GetEndpointsByID(sandbox string) ([]string, error) {
	eps, err := hcn.GetNamespaceEndpointIds(sandbox)
	if err != nil {
		return nil, err
	}
	return eps, nil
}

func isEndpointAttached(epId string, sandbox string) (bool, error) {
	attachedEpIds, err := GetEndpointsByID(sandbox)
	if err != nil {
		return false, err
	}
	for _, existingEP := range attachedEpIds {
		if existingEP == epId {
			return true, nil
		}
	}
	return false, nil
}

func attachEndpoint(epId string, sandbox string) error {
	hcnEp, err := hcn.GetEndpointByID(epId)
	if err != nil {
		return errors.New("Can not find endpoint by netns")
	}
	if err := hcnEp.NamespaceAttach(sandbox); err != nil {
		return err
	}
	return nil

}
