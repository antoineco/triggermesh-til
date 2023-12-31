/*
Copyright 2021 TriggerMesh Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"io"
	"os"

	"til/cli"
)

func main() {
	if err := run(os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "Error running command: %s\n", err)
		os.Exit(1)
	}
}

// run executes the command.
func run(args []string, stdout, stderr io.Writer) error {
	c := cli.New(args, usage,
		cli.StdWriter(stdout),
		cli.ErrWriter(stderr),

		cli.Subcommand(cmdGenerate, new(GenerateCommand)),
		cli.Subcommand(cmdValidate, new(ValidateCommand)),
		cli.Subcommand(cmdGraph, new(GraphCommand)),
	)

	return c.Run()
}
