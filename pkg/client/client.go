package client

import (
	"io"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/kubectl/pkg/scheme"
)

type KubeCliRuntime struct {
	Flags          *genericclioptions.ConfigFlags
	AllNamespaces  bool
	Namespace      string
	Selector       string
	FieldSelector  string
	ForceConflicts bool
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
	// Generates the objects using the resource builder if they have not
	// already been stored by calling "SetObjects()" in the pre-processor.
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
		return utilerrors.NewAggregate(errs)
	}
	return nil
}

func (k *KubeCliRuntime) applyOneObject(info *resource.Info) error {
	helper := resource.NewHelper(info.Client, info.Mapping).WithFieldManager("kubectl-kcl-client-side-apply")
	// Send the full object to be applied on the server side.
	data, err := runtime.Encode(unstructured.UnstructuredJSONScheme, info.Object)
	if err != nil {
		return err
	}
	obj, err := helper.Patch(
		info.Namespace,
		info.Name,
		types.ApplyPatchType,
		data,
		&metav1.PatchOptions{
			Force: &k.ForceConflicts,
		},
	)
	if err != nil {
		return err
	}
	if err := info.Refresh(obj, true); err != nil {
		return err
	}
	// TODO: use kubectl ApplyOptions intead of genericclioptions.
	printer, err := genericclioptions.NewPrintFlags("configured").ToPrinter()
	if err != nil {
		return err
	}

	if err = printer.PrintObj(info.Object, os.Stdout); err != nil {
		return err
	}
	return nil
}
