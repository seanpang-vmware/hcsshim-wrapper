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
	return hc
}

func NewHcnGetEndpointsBySandboxIdCommand(args []string) *HcnGetEndpointsBySandboxIdCommand {
	hc := &HcnGetEndpointsBySandboxIdCommand{
		Fs: flag.NewFlagSet("HCNGetNamespaceEndpoints", flag.ContinueOnError),
	}
	return hc
}

func NewHcnIsAttachedCommand(args []string) *HcnIsAttachedCommand {
	hc := &HcnIsAttachedCommand{
		Fs: flag.NewFlagSet("HCNIsAttached", flag.ContinueOnError),
	}
	return hc
}

type Runner interface {
	Init([]string) error
	Run() (string, error)
	Name() string
}

func root(args []string) error {
	if len(args) < 1 {
		return errors.New("Error: No subcommand found in command")
	}

	cmds := []Runner{
		NewHcnGetEndpointsBySandboxIdCommand(args[1:]),
		NewHcnAttachCommand(args[1:]),
		NewHcnIsAttachedCommand(args[1:]),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(args[1:])
			output, err := cmd.Run()
			if err != nil {
				return err
			}
			fmt.Println(output)
			return nil
		}
	}
	return fmt.Errorf("Error: Unknown subcommand: %s", subcommand)
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Printf("Error: failed to run command %s", err)
		fmt.Println(err)
		os.Exit(1)
	}

}
