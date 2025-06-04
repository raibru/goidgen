package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	prtVersion bool
)

var rootCmd = &cobra.Command{
	Use:   "goidgen",
	Short: "Go ID Generator",
	Long: `
Create diffrent IDs used in data. See subcommands.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handleGlobalParams(cmd, args); err != nil {
			cmd.Help()
			fmt.Println("\nRoot command parsing error: ", err)
			os.Exit(1)
		}
	},
}

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Execute root cmd has error", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&prtVersion, "version", "v", false, "display goidgen simple version")
}

func handleGlobalParams(cmd *cobra.Command, args []string) error {
	// Go lint behavior
	_ = cmd
	_ = args

	if prtVersion {
		PrintVersion(os.Stdout)
	} else {
		cmd.Help()
	}

	return nil
}
