package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"kusionstack.io/kubectl-kcl/pkg/options"
)

// NewRunCmd returns the run command.
func NewRunCmd() *cobra.Command {
	o := options.NewRunOptions()
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run KCL codes.",
		RunE: func(*cobra.Command, []string) error {
			fmt.Println(o.FilePath)
			return nil
		},
		SilenceUsage: true,
	}

	f := cmd.Flags()
	f.StringVar(&o.FilePath, "file", "", "input kcl file or path to pass to helm kcl template")

	return cmd
}
