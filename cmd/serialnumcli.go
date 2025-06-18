package cmd

import (
	"fmt"
	"os"

	"github.com/raibru/goidgen/gen/serialnumid"
	"github.com/spf13/cobra"
)

var (
	printSerialNumIdExamples bool
	serialNumIdParam         serialnumid.GenerateParam
)

var serialNumIdCmd = &cobra.Command{
	Use:   "serial-num-id",
	Short: "Create serial number id values",
	Long: `
Create serial number id values handled in data generations. 
The number value is temporaly persisted in an file with last generated number so
goidgen use multible times in scripts or with other tools

Execute for further information

goidgen serial-num-id --examples

`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handleSerialNumIdParams(cmd, args); err != nil {
			//cmd.Help()
			fmt.Printf("\nserialNumId command parsing error:\n%v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serialNumIdCmd)
	serialNumIdCmd.PersistentFlags().BoolVarP(&printSerialNumIdExamples, "examples", "", false, "print serialNumId examples to stdout.")
	serialNumIdCmd.PersistentFlags().IntVarP(&serialNumIdParam.StartNum, "start-number", "s", 1, "number start serial number counting")
	serialNumIdCmd.PersistentFlags().IntVarP(&serialNumIdParam.NumCount, "number-count", "n", 1, "number of generated serial numbers")
	serialNumIdCmd.PersistentFlags().BoolVarP(&serialNumIdParam.CleanNum, "clean-serial-number", "c", false, "cleanup of counter in temp-file")
	serialNumIdCmd.PersistentFlags().StringVarP(&serialNumIdParam.TmpDataFile, "temp-file", "t", "./sernum.dat", "location and filename containing serial number.")
	serialNumIdCmd.PersistentFlags().StringVarP(&serialNumIdParam.OutputFile, "output-file", "o", "", "output to file instead of stdout.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func handleSerialNumIdParams(cmd *cobra.Command, args []string) error {
	//if len(args) == 0 {
	//	cmd.Help()
	//	return nil
	//}

	if printSerialNumIdExamples {
		fmt.Printf("\n%s\n\n", serialNumIdExamples())
		return nil
	}

	if serialNumIdParam.CleanNum {
		if err := os.Remove(serialNumIdParam.TmpDataFile); err != nil {
			fmt.Printf("Failed to cleanup serial number id file %s", serialNumIdParam.TmpDataFile)
			return err
		}
		return nil
	}

	err := serialnumid.Validate(&serialNumIdParam)
	if err != nil {
		fmt.Printf("Failed validating serial number id parameter")
		return err
	}

	err = serialnumid.ReadSerialNumber(&serialNumIdParam)
	if err != nil {
		fmt.Println("Failed parsing persited serial number from file")
		return err
	}

	result, err := serialnumid.GenerateId(&serialNumIdParam)
	if err != nil {
		fmt.Printf("Failed generating serial number id")
		return err
	}

	err = serialnumid.DumpId(result, &serialNumIdParam)
	if err != nil {
		fmt.Printf("Failed dumping serial number id")
		return err
	}

	err = serialnumid.WriteSerialNumber(&serialNumIdParam)
	if err != nil {
		fmt.Println("Failed saving current serial number into file")
		return err
	}

	return nil
}

func serialNumIdExamples() string {
	return `
Some examples how to use the Serial Number ID generator.

Examples:

	TODO:rbr:2025-06-15: some examples here

`
}
