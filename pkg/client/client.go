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

// Get kubernetes resource object info from io reader
func (k *KubeCliRuntime) GetObjects(flags resource.RESTClientGetter, r io.Reader) ([]*resource.Info, error) {
	request := resource.NewBuilder(flags).
		Unstructured().
		ContinueOnError().
		NamespaceParam(k.Namespace).DefaultNamespace().
		Stream(r, "").
		LabelSelectorParam(k.Selector).
		Flatten().
		Do()
	return request.Infos()
}

// Apply yaml file from io reader
func (k *KubeCliRuntime) Apply(flags resource.RESTClientGetter, r io.Reader) error {
	errs := []error{}
	infos, err := k.GetObjects(flags, r)
	if err != nil {
		errs = append(errs, err)
	}
	if len(infos) == 0 && len(errs) == 0 {
		return err
	}
	// Iterate through all objects, applying each one.
	for _, info := range infos {
		if err := k.applyOneObject(info); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) == 1 {
		return errs[0]
	}
	if len(errs) > 1 {
		return errs[0]
	}
	return nil
}

func (k *KubeCliRuntime) applyOneObject(info *resource.Info) error {
	helper := resource.NewHelper(info.Client, info.Mapping).WithFieldManager("kubectl-client-side-apply")
	obj, err := helper.Replace(
		info.Namespace,
		info.Name,
		true,
		info.Object,
	)
	if err != nil {
		return err
	}
	if err := info.Refresh(obj, true); err != nil {
		return err
	}
	return nil
}
