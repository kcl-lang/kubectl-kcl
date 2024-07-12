package client

import (
	"context"
	"fmt"
	"io"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	krmkube "kcl-lang.io/krm-kcl/pkg/kube"
	"kcl-lang.io/krm-kcl/pkg/kube/engine"
	"kcl-lang.io/kubectl-kcl/pkg/kube"
)

type KubeCliRuntime struct {
	engine *engine.Engine
}

func NewKubeCliRuntime() (*KubeCliRuntime, error) {
	e, err := engine.NewFromClientGetter(kube.KubeConfigFlags)
	if err != nil {
		return nil, err
	}
	return &KubeCliRuntime{
		engine: e,
	}, nil
}

// Apply yaml file from io reader
func (k *KubeCliRuntime) Apply(r io.Reader) error {
	// Generates the objects using the resource builder if they have not
	// already been stored by calling "SetObjects()" in the pre-processor.
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	objects, err := krmkube.YamlStreamByteToUnstructuredList(data)
	if err != nil {
		return err
	}
	// Normalize KRM KCL Pipeline output YAML Byte to JSON Byte in `unstructured.Unstructured`.
	applyToObjects := make([]*unstructured.Unstructured, len(objects))
	for i, o := range objects {
		j, err := o.MarshalJSON()
		if err != nil {
			return err
		}
		jObj, err := krmkube.JsonByteToUnstructured(j)
		if err != nil {
			return err
		}
		applyToObjects[i] = jObj
	}
	status, err := k.engine.ApplyAll(context.Background(), applyToObjects, &engine.ApplyOptions{})
	for _, e := range status.Entries {
		fmt.Printf("%s %s\n", e.ObjMetadata.ID(), string(e.Status))
	}
	if err != nil {
		return err
	}
	return nil
}
