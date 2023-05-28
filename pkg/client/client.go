package client

import (
	"io"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/kubectl/pkg/scheme"
)

type KubeCliRuntime struct {
	Flags         *genericclioptions.ConfigFlags
	AllNamespaces bool
	Namespace     string
	Selector      string
	FieldSelector string
}

func NewKubeCliRuntime() *KubeCliRuntime {
	return &KubeCliRuntime{
		Flags: genericclioptions.NewConfigFlags(true),
	}
}

// GetGeneralResources get kubernetes general resource like `kubectl get all`.
func (k *KubeCliRuntime) GetGeneralResources(w io.Writer) error {
	r, err := k.fetchResourcesBulk(k.Flags)
	if err != nil {
		return err
	}
	p := printers.NewTypeSetter(scheme.Scheme).ToPrinter(&printers.YAMLPrinter{})
	if err := p.PrintObj(r, w); err != nil {
		return err
	}
	return nil
}

func (k *KubeCliRuntime) fetchResourcesBulk(flags resource.RESTClientGetter) (runtime.Object, error) {
	resources := []string{"deployments", "daemonsets", "service", "pod"}

	request := resource.NewBuilder(flags).
		Unstructured().
		ResourceTypes(resources...).
		NamespaceParam(k.Namespace).DefaultNamespace().AllNamespaces(k.AllNamespaces).
		LabelSelectorParam(k.Selector).FieldSelectorParam(k.FieldSelector).SelectAllParam(k.Selector == "" && k.FieldSelector == "").
		Flatten().
		Latest()
	return request.Do().Object()
}
