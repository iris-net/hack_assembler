package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "assembler",
	Short: "An assembler for HACK platform",
}

func Execute() error {
	rootCmd.SetOut(os.Stdout)

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
