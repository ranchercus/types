package v3

import (
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type HarborRepositoryLifecycle interface {
	Create(obj *HarborRepository) (runtime.Object, error)
	Remove(obj *HarborRepository) (runtime.Object, error)
	Updated(obj *HarborRepository) (runtime.Object, error)
}

type harborRepositoryLifecycleAdapter struct {
	lifecycle HarborRepositoryLifecycle
}

func (w *harborRepositoryLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *harborRepositoryLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *harborRepositoryLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*HarborRepository))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *harborRepositoryLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*HarborRepository))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *harborRepositoryLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*HarborRepository))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewHarborRepositoryLifecycleAdapter(name string, clusterScoped bool, client HarborRepositoryInterface, l HarborRepositoryLifecycle) HarborRepositoryHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(HarborRepositoryGroupVersionResource)
	}
	adapter := &harborRepositoryLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *HarborRepository) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
