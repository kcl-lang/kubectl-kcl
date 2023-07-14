package options

import (
	"testing"
)

func TestApplyOptions_Run(t *testing.T) {
	type fields struct {
		InputPath  string
		OutputPath string
		Namespace  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"test1",
			fields{
				InputPath:  "../../examples/kcl-apply.yaml",
				OutputPath: "",
				Namespace:  "",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ApplyOptions{
				RunOptions: RunOptions{
					InputPath:  tt.fields.InputPath,
					OutputPath: tt.fields.OutputPath,
				},
				Namespace: tt.fields.Namespace,
			}
			if err := o.Run(); (err != nil) != tt.wantErr {
				t.Errorf("ApplyOptions.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
