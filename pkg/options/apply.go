package options

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"kcl-lang.io/kcl-go/pkg/logger"
	"kcl-lang.io/krm-kcl/pkg/kio"
	"kcl-lang.io/kubectl-kcl/pkg/client"
)

// ApplyOptions is the options for the apply sub command.
type ApplyOptions struct {
	RunOptions

	// Namespace is the -n flag --namespace. It will set kubernetes namespace scope
	Namespace string

	// Selector is the -l flag --selector. It will return by label selector
	Selector string

	// FieldSelector is the --field-selector. Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector
	// key1=value1,key2=value2). The server only supports a limited number of field queries per type
	FieldSelector string
}

// NewRunOptions() create a new apply options for the apply command.
func NewApplyOptions() *ApplyOptions {
	return &ApplyOptions{}
}

// Run apply command option.
func (o *ApplyOptions) Run() error {
	cli, err := o.getCliRuntime()
	if err != nil {
		return err
	}

	reader, err := o.reader()
	if err != nil {
		logger.GetLogger().Errorf("read manifests err: %s", err.Error())
		return err
	}
	// use the io pipe feat to connect io writer to io reader
	pr, pw := io.Pipe()
	go func() {
		pipeline := kio.NewPipeline(reader, pw, true)
		if err := pipeline.Execute(); err != nil {
			logger.GetLogger().Errorf("pipeline execute err: %s", err.Error())
		}
		defer pw.Close()
	}()
	var input bytes.Buffer
	_, err = input.ReadFrom(pr)
	if err != nil {
		logger.GetLogger().Errorf("io reader err: %s", err.Error())
		return err
	}
	if err := cli.Apply(&input); err != nil {
		logger.GetLogger().Errorf("apply err: %s", err.Error())
		return err
	}
	return nil
}

func (o *ApplyOptions) getCliRuntime() (*client.KubeCliRuntime, error) {
	cliRuntime, err := client.NewKubeCliRuntime()
	if err != nil {
		return nil, err
	}
	return cliRuntime, nil
}

// Validate the options.
func (o *ApplyOptions) Validate() error {
	return nil
}

func (o *ApplyOptions) reader() (io.Reader, error) {
	if o.InputPath == "-" {
		return os.Stdin, nil
	} else {
		file, err := os.Open(o.InputPath)
		if err != nil {
			return nil, err
		}
		return bufio.NewReader(file), nil
	}
}

func (o *ApplyOptions) writer() (io.Writer, error) {
	if o.OutputPath == "" {
		return os.Stdout, nil
	} else {
		file, err := os.OpenFile(o.OutputPath, os.O_CREATE|os.O_RDWR, 0744)
		if err != nil {
			return nil, err
		}
		return bufio.NewWriter(file), nil
	}
}
