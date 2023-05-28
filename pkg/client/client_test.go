package client

import (
	"os"
	"testing"
)

func TestGetGeneralResources(t *testing.T) {
	cli := NewKubeCliRuntime()
	cli.Namespace = "default"
	err := cli.GetGeneralResources(os.Stdout)
	if err != nil {
		t.Fatalf("GetGeneralResources err: %s", err.Error())
	}
}
