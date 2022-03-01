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

func NewHcnGetEndpointsBySandboxIdCommand(args []string) *HcnGetEndpointsBySandboxIdCommand {
	hc := &HcnGetEndpointsBySandboxIdCommand{
		Fs: flag.NewFlagSet("HCNGetNamespaceEndpoints", flag.ContinueOnError),
	}

	hc.Init(args)

	return hc
}

type Runner interface {
	Init([]string) error
	Run() (string, error)
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
			output, err := cmd.Run()
			if err != nil {
				fmt.Println(output)
			}
			return nil
		}
	}
	return fmt.Errorf("Error: Unknown subcommand: %s", subcommand)
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
