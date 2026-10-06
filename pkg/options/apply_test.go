package options

import (
	"testing"
)

func TestApplyOptions_Run(t *testing.T) {
	type fields struct {
		InputPaths []string
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
				InputPaths: []string{"../../examples/kcl-apply.yaml"},
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
					InputPaths: tt.fields.InputPaths,
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

// TestRunOptions_reader_MultipleFiles regresses #10: a single -f flag
// must support more than one input file and concatenate them through
// the kio pipeline.
func TestRunOptions_reader_MultipleFiles(t *testing.T) {
	o := &RunOptions{
		InputPaths: []string{
			"../../examples/kcl-apply.yaml",
			"../../examples/kcl-apply.yaml",
		},
	}
	r, err := o.reader()
	if err != nil {
		t.Fatalf("reader() returned error for two existing files: %v", err)
	}
	if r == nil {
		t.Fatal("reader() returned a nil io.Reader")
	}
	buf := make([]byte, 4096)
	n, _ := r.Read(buf)
	if n == 0 {
		t.Fatal("reader() produced empty stream for two non-empty inputs; io.MultiReader wiring is broken")
	}
}

// TestRunOptions_reader_StdinFallback covers the empty-slice branch which
// leaves the original "no -f at all" UX intact.
func TestRunOptions_reader_StdinFallback(t *testing.T) {
	o := &RunOptions{}
	r, err := o.reader()
	if err != nil {
		t.Fatalf("reader() with no paths should default to stdin, got error: %v", err)
	}
	if r == nil {
		t.Fatal("reader() returned nil for empty InputPaths")
	}
}

// TestRunOptions_reader_MissingPath covers the error path where a
// configured file does not exist.
func TestRunOptions_reader_MissingPath(t *testing.T) {
	o := &RunOptions{
		InputPaths: []string{"../../does-not-exist.yaml"},
	}
	if _, err := o.reader(); err == nil {
		t.Fatal("reader() should error on missing file, got nil")
	}
}
