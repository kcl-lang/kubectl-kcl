package cmd

import (
	"github.com/spf13/cobra"
	"kcl-lang.io/kubectl-kcl/pkg/options"
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
	f.StringSliceVarP(&o.InputPaths, "filename", "f", nil, "input kcl spec file(s) to pass to kubectl kcl (repeatable, or comma-separated)")
	f.StringVarP(&o.OutputPath, "output", "o", "", "output yaml path, default is stdout")

	return cmd
}
