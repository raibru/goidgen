/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/spf13/cobra"
)

var (
	printUuidExamples bool
	uuidVersionFlag   string
	namespaceID       string
	nameData          string
	outputFile        string
	numUUIDs          int
)

// pusCmd represents the pus command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Create diffrent UUID values handle in data",
	Long: `
Create diffrent UUID values handled in data generations.

Execute for further information

goidgen uuid --examples

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("### in uuid run ...\n")
		if err := handleUuidParams(cmd, args); err != nil {
			//cmd.Help()
			fmt.Printf("\nuuid command parsing error:\n%v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
	uuidCmd.PersistentFlags().BoolVarP(&printUuidExamples, "examples", "", false, "print uuid examples to stdout")
	// uuidCmd.PersistentFlags().StringVarP(&uuidVersionFlag, "uuid-version", "v", "4", "UUID version to generate - Default is 4:\n\t- 1 (time-based),\n\t- 3 (name-MD5),\n\t- 4 (random),\n\t- 5 (name-SHA1),\n\t- 6 (time-based sortable),\n\t- 7 (time-based sortable).")
	uuidCmd.PersistentFlags().StringVarP(&uuidVersionFlag, "uuid-version", "V", "4", "UUID version to generate - Default is 4:\n\t- 1 (time-based),\n\t- 3 (name-MD5),\n\t- 4 (random),\n\t- 5 (name-SHA1),\n\t- 6 (time-based sortable),\n\t- 7 (time-based sortable).")
	uuidCmd.PersistentFlags().StringVarP(&namespaceID, "namespace", "n", "", "Namespace UUID for name-based UUIDs (version 3 or 5).\nMust be a valid UUID string.")
	uuidCmd.PersistentFlags().StringVarP(&nameData, "name-data", "N", "", "Name data for name-based UUIDs (version 3 or 5).")
	uuidCmd.PersistentFlags().StringVarP(&outputFile, "output-file", "o", "", "Output to file instead of stdout.")
	uuidCmd.PersistentFlags().IntVarP(&numUUIDs, "num-to-generate", "c", 1, "Number of UUIDs to generate.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func handleUuidParams(cmd *cobra.Command, args []string) error {
	fmt.Print("### in uuid handleUuidParams ...\n")
	if len(args) == 0 {
		cmd.Help()
		return nil
	}

	if printUuidExamples {
		fmt.Printf("\n%s\n\n", uuidExamples())
		return nil
	}

	return nil
}

func uuidExamples() string {
	return `
Some examples how to use the UUID generator.

Examples:

Generate a random UUID (v4)
>	goidgen uuid

Generate a time-based UUID (v1)
> goidgen uuid -v 1
> goidgen uuid -uuid-version 1

Generate a time-based sortable UUID (v6)
> goidgen uuid -v 6

Generate a time-based sortable UUID (v7)
> goidgen uuid -v 7                 # 

Name-based v3
> goidgen uuid -v 3 -n 00000000-0000-0000-0000-000000000000 -N example.com
> goidgen uuid -v 3 -n 00000000-0000-0000-0000-000000000000 -name-data example.com

Name-based v5 (DNS namespace)
> goidgen -v 5 -n 6ba7b810-9dad-11d1-80b4-00c04fd430c8 -N www.example.org
> goidgen -v 5 --namespace 6ba7b810-9dad-11d1-80b4-00c04fd430c8 -N www.example.org

Generate 5 random UUIDs
> goidgen -c 5
> goidgen --num-to-generate 5

Output UUID to a file
> goidgen	 -o output.txt
> goidgen	 -output-file output.txt

`
}
