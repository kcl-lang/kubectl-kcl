package options

import (
	"bufio"
	"io"
	"os"

	"kcl-lang.io/krm-kcl/pkg/kio"
)

// RunOptions is the options for the run command
type RunOptions struct {
	// InputPaths is the -f flag. Multiple files may be passed as either
	// repeated `-f` flags or as a single comma-separated `-f a,b` value.
	// When empty (or contains only "-") input is read from stdin.
	InputPaths []string
	// OutputPath is the -o flag
	OutputPath string
}

// RunOptions creates a new options for the run command.
func NewRunOptions() *RunOptions {
	return &RunOptions{}
}

// Run the with the run command options.
func (o *RunOptions) Run() error {
	reader, err := o.reader()
	if err != nil {
		return err
	}
	writer, err := o.writer()
	if err != nil {
		return err
	}
	pipeline := kio.NewPipeline(reader, writer, false)
	return pipeline.Execute()
}

// Validate the options.
func (o *RunOptions) Validate() error {
	return nil
}

// reader returns an io.Reader that concatenates all configured input
// sources. An empty InputPaths reads from stdin; a single "-" entry
// also reads from stdin; otherwise each file is opened and streamed in
// order via io.MultiReader so multi-document YAMLs flow into the
// pipeline sequentially (kubectl `kustomize/kyaml` expects `---`
// separators in the combined stream).
func (o *RunOptions) reader() (io.Reader, error) {
	if len(o.InputPaths) == 0 {
		return os.Stdin, nil
	}
	readers := make([]io.Reader, 0, len(o.InputPaths))
	for _, p := range o.InputPaths {
		if p == "-" {
			readers = append(readers, os.Stdin)
			continue
		}
		f, err := os.Open(p)
		if err != nil {
			return nil, err
		}
		readers = append(readers, bufio.NewReader(f))
	}
	return io.MultiReader(readers...), nil
}

func (o *RunOptions) writer() (io.Writer, error) {
	if o.OutputPath == "" {
		return os.Stdout, nil
	}
	file, err := os.OpenFile(o.OutputPath, os.O_CREATE|os.O_RDWR, 0744)
	if err != nil {
		return nil, err
	}
	return bufio.NewWriter(file), nil
}
