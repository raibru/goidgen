package cmd

import (
	"fmt"
	"os"

	"github.com/raibru/goidgen/gen/uuid"
	"github.com/spf13/cobra"
)

var (
	printUuidExamples bool
	uuidParam         uuid.GenerateParam
)

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Create diffrent UUID values handle in data",
	Long: `
Create different UUID values handled in data generations.

Execute for further information

goidgen uuid --examples

`,
	Run: func(cmd *cobra.Command, args []string) {
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
	uuidCmd.PersistentFlags().StringVarP(&uuidParam.UuidVersionFlag, "uuid-version", "V", "4", "use UUID version to generate:\n\t(1) time-based,\n\t(3) name-MD5,\n\t(4) random,\n\t(5) name-SHA1,\n\t(6) time-based sortable,\n\t(7) time-based sortable\n")
	uuidCmd.PersistentFlags().StringVarP(&uuidParam.NamespaceID, "namespace", "n", "", "namespace UUID for name-based UUIDs (version 3 or 5).\nMust be a valid UUID string.")
	uuidCmd.PersistentFlags().StringVarP(&uuidParam.NameData, "name-data", "N", "", "name data for name-based UUIDs (version 3 or 5).")
	uuidCmd.PersistentFlags().StringVarP(&uuidParam.OutputFile, "output-file", "o", "", "output to file instead of stdout.")
	uuidCmd.PersistentFlags().IntVarP(&uuidParam.NumUUIDs, "num-to-generate", "c", 1, "number of UUIDs to generate.")
	uuidCmd.PersistentFlags().BoolVarP(&uuidParam.ToUppercases, "uppercase", "u", false, "print uuid letters in uppercase. Default is lowercase")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func handleUuidParams(cmd *cobra.Command, args []string) error {
	//if len(args) == 0 {
	//	cmd.Help()
	//	return nil
	//}

	if printUuidExamples {
		fmt.Printf("\n%s\n\n", uuidExamples())
		return nil
	}

	err := uuid.Validate(&uuidParam)
	if err != nil {
		fmt.Printf("Failed validating uuid parameter: %v", err)
		return err
	}

	result, err := uuid.GenerateId(&uuidParam)
	if err != nil {
		fmt.Printf("Failed generating uuid: %v", err)
		return err
	}

	err = uuid.DumpId(result, &uuidParam)
	if err != nil {
		fmt.Printf("Failed dumping uuid: %v", err)
		return err
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
