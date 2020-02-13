package v3

import (
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type PipelineTemplateLifecycle interface {
	Create(obj *PipelineTemplate) (runtime.Object, error)
	Remove(obj *PipelineTemplate) (runtime.Object, error)
	Updated(obj *PipelineTemplate) (runtime.Object, error)
}

type pipelineTemplateLifecycleAdapter struct {
	lifecycle PipelineTemplateLifecycle
}

func (w *pipelineTemplateLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *pipelineTemplateLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *pipelineTemplateLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*PipelineTemplate))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *pipelineTemplateLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*PipelineTemplate))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *pipelineTemplateLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*PipelineTemplate))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewPipelineTemplateLifecycleAdapter(name string, clusterScoped bool, client PipelineTemplateInterface, l PipelineTemplateLifecycle) PipelineTemplateHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(PipelineTemplateGroupVersionResource)
	}
	adapter := &pipelineTemplateLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *PipelineTemplate) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
