package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewApplyCmd returns the run command.
func NewApplyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "apply",
		Short: "Apply manifests defined by KCL.",
		Run: func(*cobra.Command, []string) {
			fmt.Println("Hello kubectl kcl apply")
		},
		SilenceUsage: true,
	}
}
