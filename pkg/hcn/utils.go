package hcn

import "github.com/Microsoft/hcsshim/hcn"

func GetEndpointByID(epId string) ([]string, error) {
	eps, err := hcn.GetNamespaceEndpointIds(epId)
	if err != nil {
		return eps, nil
	}
	return nil, err
}
