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
	"fmt"
	"vdoc/vdoc-system/service/web"

	"github.com/spf13/cobra"
)

var (
	sweb bool
	port int
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "internal service",
	Long:  `start service `,
	Run: func(cmd *cobra.Command, args []string) {
		if sweb {
			fmt.Println("Start web service")
			web.Start(port)
		}
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
	serviceCmd.Flags().BoolVarP(&sweb, "web", "w", true, "start web service")
	serviceCmd.Flags().IntVarP(&port, "port", "p", 8001, "start web service")
}
