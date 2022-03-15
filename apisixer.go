// Copyright 2022 kwanhur
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package main

import (
	"log"

	"github.com/spf13/cobra"
)

var commit string

var sixer = &cobra.Command{
	Use:   "apisixer",
	Short: "Apache APISIX release candidate verifier",
	Run:   sixerRun,
}

func init() {
	sixer.AddCommand(versionCmd, verboseCmd)
	sixer.AddCommand(dashboardCmd)
}

func init() {
	globals := sixer.PersistentFlags()
	BindRCFlag(globals)

	BindVerFlags(sixer.Flags())
}

func sixerPreRun(cmd *cobra.Command, args []string) {
	if ReleaseCandidate == "" {
		log.Fatalln("Please specify release candidate version first.")
	}
}

func sixerRun(cmd *cobra.Command, args []string) {
	if Version {
		showVersion()
	}

	if Verbose {
		showVerbose()
	}
}

func main() {
	if err := sixer.Execute(); err != nil {
		log.Fatalln("apisixer run failed:", err)
	}
}
