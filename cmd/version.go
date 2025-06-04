package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// var section
var (
	major = "0"
	minor = "1"
	patch = "0"
	//buildTag  = "-"
	//buildDate = "-"
	buildInfo = "-"
	appName   = "goidgen - go ID generator"
	author    = "raibru <github.com/raibru>"
	license   = "MIT License"
)

// PrintVersion prints the tool versions string
func PrintVersion(w io.Writer) {
	fmt.Fprintf(w, "v%s.%s.%s\n", major, minor, patch)
}

// PrintVersion prints the tool versions string
func PrintFullVersion(w io.Writer) {
	fmt.Fprintf(w, "%s - v%s.%s.%s\n", appName, major, minor, patch)
	fmt.Fprintf(w, "  Build  : %s\n", buildInfo)
	fmt.Fprintf(w, "  Author : %s\n", author)
	fmt.Fprintf(w, "  License: %s\n\n", license)
}

// edenCmd represents the eden dump command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "full version info of goidgen",
	Long: `

goidgen version
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handleVersionParams(cmd, args); err != nil {
			//cmd.Help()
			fmt.Printf("\ngoidgen parsing error:\n%v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func handleVersionParams(cmd *cobra.Command, args []string) error {
	// Go lint behavior
	_ = cmd
	_ = args
	PrintFullVersion(os.Stdout)
	return nil
}
