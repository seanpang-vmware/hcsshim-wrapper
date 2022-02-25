/* Copyright © 2022 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: Apache-2.0 */

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	. "github.com/seanpang-vmware/hcsshim-wrapper/pkg/hcn"
)

var ()

func init() {

}

func NewHcnAttachCommand(args []string) *HcnAttachCommand {
	hc := &HcnAttachCommand{
		Fs: flag.NewFlagSet("HCNAttach", flag.ContinueOnError),
	}

	hc.Init(args)

	return hc
}

func NewHcnGetEndpointByIDCommand(args []string) *GetEndpointByIDCommand {
	hc := &GetEndpointByIDCommand{
		Fs: flag.NewFlagSet("HCNAttach", flag.ContinueOnError),
	}

	hc.Init(args)

	return hc
}

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func root(args []string) error {
	if len(args) < 1 {
		return errors.New("No subcommand found in command")
	}

	cmds := []Runner{
		NewHcnAttachCommand(args[1:]),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			fmt.Println("Get hcn command", cmd.Name)
			output = cmd.Run()
			if output != nil {
				b, err := json.Marshal(output)
				if err != nil {
					fmt.Println(string(b))
				}
			}
			return nil
		}
	}
	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
