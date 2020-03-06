package v3

import (
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type HarborProjectLifecycle interface {
	Create(obj *HarborProject) (runtime.Object, error)
	Remove(obj *HarborProject) (runtime.Object, error)
	Updated(obj *HarborProject) (runtime.Object, error)
}

type harborProjectLifecycleAdapter struct {
	lifecycle HarborProjectLifecycle
}

func (w *harborProjectLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *harborProjectLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *harborProjectLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*HarborProject))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *harborProjectLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*HarborProject))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *harborProjectLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*HarborProject))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewHarborProjectLifecycleAdapter(name string, clusterScoped bool, client HarborProjectInterface, l HarborProjectLifecycle) HarborProjectHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(HarborProjectGroupVersionResource)
	}
	adapter := &harborProjectLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *HarborProject) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
