/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// xenStartCmd represents the start command
var xenStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Run shell command with arguments in 'start' action on 'xen' mode",
	Long: `Run shell command with arguments in 'start' action on 'xen' mode. For example:

eveadm xen start uuid`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		xenctx.containerUUID = arg
		err, args, envs := xenctx.xenStartToCmd()
		if err != nil {
			log.Fatalf("Error in obtain params in %s", cmd.Name())
		}
		xenctx.xenRuneWrapper(Timeout, args, envs, cmd.Name())
	},
}

func init() {
	xenCmd.AddCommand(xenStartCmd)
}
