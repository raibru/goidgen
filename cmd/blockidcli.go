package cmd

import (
	"fmt"
	"os"

	"github.com/raibru/goidgen/gen/blockid"
	"github.com/spf13/cobra"
)

var (
	printBlockidExamples bool
	blockidParam         blockid.GenerateParam
)

var blockidCmd = &cobra.Command{
	Use:   "blockid",
	Short: "Create block id values with count of bytes printed in hex notation and pre/post-fix",
	Long: `
Create different block id values handled in data generations. This is build by
n count of bytes printed in hex value notation. Using string parameter as pre/post fix around hex
value

Execute for further information

goidgen blockid --examples

`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handleBlockidParams(cmd, args); err != nil {
			//cmd.Help()
			fmt.Printf("\nblockid command parsing error:\n%v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(blockidCmd)
	blockidCmd.PersistentFlags().BoolVarP(&printBlockidExamples, "examples", "", false, "print blockid examples to stdout.")
	blockidCmd.PersistentFlags().StringVarP(&blockidParam.PrefixName, "prefix-name", "b", "", "prefix name set before hex values.")
	blockidCmd.PersistentFlags().StringVarP(&blockidParam.PostfixName, "postfix-name", "e", "", "postfixi name set after hex values.")
	blockidCmd.PersistentFlags().IntVarP(&blockidParam.NumByteCount, "num-byte-count", "n", 4, "number of bytes used as base hex value. Default is 4.")
	blockidCmd.PersistentFlags().IntVarP(&blockidParam.NumBlockids, "num-to-generate", "c", 1, "number of block ids to generate.")
	blockidCmd.PersistentFlags().StringVarP(&blockidParam.OutputFile, "output-file", "o", "", "output to file instead of stdout.")
	blockidCmd.PersistentFlags().BoolVarP(&blockidParam.ToUppercases, "uppercase", "u", false, "print id letters in uppercase. Default is lowercase.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func handleBlockidParams(cmd *cobra.Command, args []string) error {
	//if len(args) == 0 {
	//	cmd.Help()
	//	return nil
	//}

	if printBlockidExamples {
		fmt.Printf("\n%s\n\n", blockidExamples())
		return nil
	}

	err := blockid.Validate(&blockidParam)
	if err != nil {
		fmt.Printf("Failed validating blockid parameter: %v", err)
		return err
	}

	result, err := blockid.GenerateId(&blockidParam)
	if err != nil {
		fmt.Printf("Failed generating blockid: %v", err)
		return err
	}

	err = blockid.DumpId(result, &blockidParam)
	if err != nil {
		fmt.Printf("Failed dumping blockid: %v", err)
		return err
	}

	return nil
}

func blockidExamples() string {
	return `
Some examples how to use the Block ID generator.

Examples:

	TODO:rbr:2025-06-15: some examples here

`
}
