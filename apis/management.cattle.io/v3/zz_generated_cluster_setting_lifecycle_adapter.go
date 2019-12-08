package v3

import (
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type ClusterSettingLifecycle interface {
	Create(obj *ClusterSetting) (runtime.Object, error)
	Remove(obj *ClusterSetting) (runtime.Object, error)
	Updated(obj *ClusterSetting) (runtime.Object, error)
}

type clusterSettingLifecycleAdapter struct {
	lifecycle ClusterSettingLifecycle
}

func (w *clusterSettingLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *clusterSettingLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *clusterSettingLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ClusterSetting))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterSettingLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ClusterSetting))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterSettingLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ClusterSetting))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewClusterSettingLifecycleAdapter(name string, clusterScoped bool, client ClusterSettingInterface, l ClusterSettingLifecycle) ClusterSettingHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(ClusterSettingGroupVersionResource)
	}
	adapter := &clusterSettingLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ClusterSetting) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
