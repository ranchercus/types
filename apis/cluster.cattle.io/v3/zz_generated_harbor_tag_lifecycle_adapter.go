package v3

import (
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type HarborTagLifecycle interface {
	Create(obj *HarborTag) (runtime.Object, error)
	Remove(obj *HarborTag) (runtime.Object, error)
	Updated(obj *HarborTag) (runtime.Object, error)
}

type harborTagLifecycleAdapter struct {
	lifecycle HarborTagLifecycle
}

func (w *harborTagLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *harborTagLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *harborTagLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*HarborTag))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *harborTagLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*HarborTag))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *harborTagLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*HarborTag))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewHarborTagLifecycleAdapter(name string, clusterScoped bool, client HarborTagInterface, l HarborTagLifecycle) HarborTagHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(HarborTagGroupVersionResource)
	}
	adapter := &harborTagLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *HarborTag) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
