package client

import (
	"io"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/kubectl/pkg/scheme"
)

// GetGeneralResources get kubernetes general resource like `kubectl get all`.
func GetGeneralResources(flags *genericclioptions.ConfigFlags, w io.Writer) error {
	r, err := fetchResourcesBulk(flags)
	if err != nil {
		return err
	}
	p := printers.NewTypeSetter(scheme.Scheme).ToPrinter(&printers.YAMLPrinter{})
	if err := p.PrintObj(r, w); err != nil {
		return err
	}
	return nil
}

func fetchResourcesBulk(flags resource.RESTClientGetter) (runtime.Object, error) {
	resources := []string{"deployments", "daemonsets", "service", "pod"}
	var ns string
	var selector string
	var fieldSelector string

	request := resource.NewBuilder(flags).
		Unstructured().
		ResourceTypes(resources...).
		NamespaceParam(ns).DefaultNamespace().AllNamespaces(ns == "").
		LabelSelectorParam(selector).FieldSelectorParam(fieldSelector).SelectAllParam(selector == "" && fieldSelector == "").
		Flatten().
		Latest()
	return request.Do().Object()
}
