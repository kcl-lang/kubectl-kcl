package options

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"

	"github.com/KusionStack/krm-kcl/pkg/kio"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"kusionstack.io/kclvm-go/pkg/logger"
	"kusionstack.io/kubectl-kcl/pkg/client"
)

// ApplyOptions is the options for the apply sub command.
type ApplyOptions struct {
	// InputPath is the -f flag
	InputPath string
	// OutputPath is the -o flag
	OutputPath string
	// GenericCliCfg is the cli-runtime config
	GenericCliCfg *genericclioptions.ConfigFlags
}

// NewRunOptions() create a new apply options for the apply command.
func NewApplyOptions() *ApplyOptions {
	return &ApplyOptions{
		GenericCliCfg: genericclioptions.NewConfigFlags(true),
	}
}

// Run apply command option.
func (o *ApplyOptions) Run() error {
	var buf bytes.Buffer
	err := client.GetGeneralResources(o.GenericCliCfg, &buf)
	if err != nil {
		logger.GetLogger().Errorf("get general resource err: %s", err.Error())
		return err
	}
	_, bs, err := o.reader()
	if err != nil {
		logger.GetLogger().Errorf("read kcl code err: %s", err.Error())
		return err
	}
	n := append(buf.Bytes(), bs...)
	reader := bytes.NewReader(n)
	writer, err := o.writer()
	if err != nil {
		return err
	}
	pipeline := kio.NewPipeline(reader, writer, true)
	return pipeline.Execute()
}

// Validate the options.
func (o *ApplyOptions) Validate() error {
	return nil
}

func (o *ApplyOptions) reader() (io.Reader, []byte, error) {
	if o.InputPath == "-" {
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			return nil, nil, errors.New("input scan err")
		}
		line := scanner.Bytes()
		return os.Stdin, line, nil
	} else {
		file, err := os.Open(o.InputPath)
		if err != nil {
			return nil, nil, err
		}
		stat, err := file.Stat()
		if err != nil {
			return nil, nil, err
		}
		bs := make([]byte, stat.Size())
		_, err = bufio.NewReader(file).Read(bs)
		if err != nil {
			return nil, nil, err
		}
		return bufio.NewReader(file), bs, nil
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
