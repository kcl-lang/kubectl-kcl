package options

// RunOptions is the options for the run command
type RunOptions struct {
	// FilePath is the filepath flag
	FilePath string
}

// RunOptions creates a new options for the run command.
func NewRunOptions() *RunOptions {
	return &RunOptions{}
}
