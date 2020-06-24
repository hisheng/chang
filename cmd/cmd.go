package cmd

import (
	"fmt"
	"github.com/hisheng/chang/cmd/croncmd"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chang",
	Short: "chang 服务",
	Long:  `chang 服务`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(crontabCmd)
	rootCmd.AddCommand(croncmd.BonusCmd)
}
