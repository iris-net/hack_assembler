package cmd

import (
	"github.com/iris-net/hack_assembler/assembler"
	"github.com/spf13/cobra"
)

var (
	out string
	in  string
)

func init() {
	execCmd.Flags().StringVar(&in, "in", "i", "where to input an assembly file")
	execCmd.Flags().StringVar(&out, "out", "o", "where to output a hack file")

	rootCmd.AddCommand(execCmd)
}

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "A generator binary code of HACK from assembly code",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := assembler.NewAssembler(in)
		if err != nil {
			return err
		}

		err = a.Execute()
		if err != nil {
			return err
		}

		err = a.ExportBinaryCode(out)
		if err != nil {
			return err
		}

		return nil
	},
}
