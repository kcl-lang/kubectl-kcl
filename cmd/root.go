package cmd

import (
	"github.com/spf13/cobra"
)

const rootCmdLongUsage = `
The Kubectl KCL Plugin.

* Generate, mutate and validate Kubernetes manifests using the KCL programming language.
`

// New creates a new cobra client
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "kcl",
		Short:        "Generate, mutate and validate Kubernetes manifests using the KCL programming language.",
		Long:         rootCmdLongUsage,
		SilenceUsage: true,
	}

	cmd.AddCommand(NewVersionCmd())
	cmd.AddCommand(NewRunCmd())
	cmd.AddCommand(NewApplyCmd())
	cmd.SetHelpCommand(&cobra.Command{}) // Disable the help command
	return cmd
}
