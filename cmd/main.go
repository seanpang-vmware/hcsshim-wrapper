/* Copyright © 2022 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: Apache-2.0 */

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var ()

func init() {

}

func NewHcnAttachCommand() *HcnAttachCommand {
	hc := &HcnAttachCommand{
		fs : flag.NewFlagSet("hcn-attach", flag.ContinueOnError)
	}

	hc.fs.StringVar(&hc.name, "epName", "", "Name of the HCS endpoint")
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
		NewHcnAttachCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {

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
