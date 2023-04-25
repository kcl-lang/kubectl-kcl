package main

import (
	"os"

	_ "github.com/KusionStack/krm-kcl/pkg/config"
	_ "kusionstack.io/kclvm-go"
	"kusionstack.io/kubectl-kcl/cmd"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		os.Exit(1)
	}
}
