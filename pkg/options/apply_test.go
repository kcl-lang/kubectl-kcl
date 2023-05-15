package options

import (
	"testing"

	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func TestApplyOptions_Run(t *testing.T) {
	type fields struct {
		InputPath     string
		OutputPath    string
		GenericCliCfg *genericclioptions.ConfigFlags
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"test1",
			fields{
				InputPath:     "../../examples/kcl-apply.yaml",
				OutputPath:    "",
				GenericCliCfg: genericclioptions.NewConfigFlags(true),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ApplyOptions{
				InputPath:     tt.fields.InputPath,
				OutputPath:    tt.fields.OutputPath,
				GenericCliCfg: tt.fields.GenericCliCfg,
			}
			if err := o.Run(); (err != nil) != tt.wantErr {
				t.Errorf("ApplyOptions.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
