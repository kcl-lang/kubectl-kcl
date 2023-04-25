package cmd

import (
	"github.com/spf13/cobra"
)

const rootCmdLongUsage = `
The Kubectl KCL Plugin.

* Edit, transformer, validate Helm charts using the KCL programming language.
`

// New creates a new cobra client
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "kcl",
		Short:        "Edit, transformer, validate Helm charts using the KCL programming language.",
		Long:         rootCmdLongUsage,
		SilenceUsage: true,
	}

	cmd.AddCommand(NewVersionCmd())
	cmd.AddCommand(NewRunCmd())
	cmd.SetHelpCommand(&cobra.Command{}) // Disable the help command
	return cmd
}
