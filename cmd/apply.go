package cmd

import (
	"github.com/spf13/cobra"
	"kcl-lang.io/kubectl-kcl/pkg/options"
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
	f.StringVarP(&o.InputPath, "file", "f", "", "input kcl spec file to pass to kubectl kcl")
	f.StringVarP(&o.OutputPath, "output", "o", "", "output yaml path, default is stdout")
	f.StringVarP(&o.Namespace, "namespace", "n", "default", "kubernetes namespace default is the default namespace ")
	f.StringVarP(&o.Selector, "selector", "l", "", "Selector (label query) to filter on.(e.g. -l key1=value1,key2=value2)")
	f.StringVarP(&o.FieldSelector, "field-selector", "", "", "Selector (field query) to filter on.(e.g. --field-selector key1=value1,key2=value2)")
	return cmd
}
