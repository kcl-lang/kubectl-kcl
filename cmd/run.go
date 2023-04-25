package cmd

import (
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
			err := o.Validate()
			if err != nil {
				return err
			}
			return o.Run()
		},
		SilenceUsage: true,
	}

	f := cmd.Flags()
	f.StringVarP(&o.InputPath, "filename", "f", "", "input kcl spec file to pass to kubectl kcl")
	f.StringVarP(&o.OutputPath, "output", "o", "", "output yaml path, default is stdout")

	return cmd
}
