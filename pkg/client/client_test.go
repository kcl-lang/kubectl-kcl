package client

import (
	"os"
	"testing"

	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func TestGetGeneralResources(t *testing.T) {
	flags := genericclioptions.NewConfigFlags(true)
	err := GetGeneralResources(flags, os.Stdout)
	if err != nil {
		t.Fatalf("GetGeneralResources err: %s", err.Error())
	}
}
