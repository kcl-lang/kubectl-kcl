package cmd

import (
	"github.com/spf13/cobra"
	"kusionstack.io/kubectl-kcl/pkg/options"
)

// NewApplyCmd returns the run command.
func NewApplyCmd() *cobra.Command {
	o := options.NewApplyOptions()
	cmd := &cobra.Command{
		Use:   "apply KCL codes to kubernetes runtime",
		Short: "Apply manifests defined by KCL.",
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
