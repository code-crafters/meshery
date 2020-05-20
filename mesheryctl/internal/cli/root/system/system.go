// Copyright 2019 The Meshery Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package system

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

const (
	url     = "http://localhost:9081"
	fileURL = "https://raw.githubusercontent.com/layer5io/meshery/master/docker-compose.yaml"
)

var systemDetails = `
Manage the state and configuration of Meshery server, adapters, and client..

Usage:
  mesheryctl system [command]

Available Commands:
  cleanup     Clean up Meshery
  help        Help for system commands
  logs        Print logs
  start       Start Meshery
  status      Check Meshery status
  stop        Stop Meshery
  update      Pull new Meshery images from Docker Hub

Flags:
  -h, --help            help for system commands

Use "mesheryctl system [command] --help" for more information about a command.
`

var (
	availableSubcommands = []*cobra.Command{}
)

// SystemCmd represents Meshery Lifecycle Management cli commands
var SystemCmd = &cobra.Command{
	Use:   "system",
	Short: "Meshery Lifecycle Management",
	Long:  `Manage the state and configuration of Meshery server, adapters, and client.`,
	//Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			log.Print(systemDetails)
			return nil
		}

		for _, subcommand := range availableSubcommands {
			if args[0] == subcommand.Name() {
				return nil
			}
		}

		return errors.New("sub-command not found : " + "\"" + args[0] + "\"")
	},
}

func init() {
	availableSubcommands = []*cobra.Command{
		cleanupCmd,
		logsCmd,
		startCmd,
		stopCmd,
		statusCmd,
		updateCmd,
	}
	SystemCmd.AddCommand(availableSubcommands...)
}