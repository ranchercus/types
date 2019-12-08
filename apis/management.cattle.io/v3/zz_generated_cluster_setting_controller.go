package v3

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	ClusterSettingGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ClusterSetting",
	}
	ClusterSettingResource = metav1.APIResource{
		Name:         "clustersettings",
		SingularName: "clustersetting",
		Namespaced:   false,
		Kind:         ClusterSettingGroupVersionKind.Kind,
	}

	ClusterSettingGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "clustersettings",
	}
)

func init() {
	resource.Put(ClusterSettingGroupVersionResource)
}

func NewClusterSetting(namespace, name string, obj ClusterSetting) *ClusterSetting {
	obj.APIVersion, obj.Kind = ClusterSettingGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type ClusterSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterSetting `json:"items"`
}

type ClusterSettingHandlerFunc func(key string, obj *ClusterSetting) (runtime.Object, error)

type ClusterSettingChangeHandlerFunc func(obj *ClusterSetting) (runtime.Object, error)

type ClusterSettingLister interface {
	List(namespace string, selector labels.Selector) (ret []*ClusterSetting, err error)
	Get(namespace, name string) (*ClusterSetting, error)
}

type ClusterSettingController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() ClusterSettingLister
	AddHandler(ctx context.Context, name string, handler ClusterSettingHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync ClusterSettingHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler ClusterSettingHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler ClusterSettingHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ClusterSettingInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*ClusterSetting) (*ClusterSetting, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ClusterSetting, error)
	Get(name string, opts metav1.GetOptions) (*ClusterSetting, error)
	Update(*ClusterSetting) (*ClusterSetting, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ClusterSettingList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ClusterSettingController
	AddHandler(ctx context.Context, name string, sync ClusterSettingHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync ClusterSettingHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle ClusterSettingLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle ClusterSettingLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ClusterSettingHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync ClusterSettingHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ClusterSettingLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle ClusterSettingLifecycle)
}

type clusterSettingLister struct {
	controller *clusterSettingController
}

func (l *clusterSettingLister) List(namespace string, selector labels.Selector) (ret []*ClusterSetting, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*ClusterSetting))
	})
	return
}

func (l *clusterSettingLister) Get(namespace, name string) (*ClusterSetting, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    ClusterSettingGroupVersionKind.Group,
			Resource: "clusterSetting",
		}, key)
	}
	return obj.(*ClusterSetting), nil
}

type clusterSettingController struct {
	controller.GenericController
}

func (c *clusterSettingController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *clusterSettingController) Lister() ClusterSettingLister {
	return &clusterSettingLister{
		controller: c,
	}
}

func (c *clusterSettingController) AddHandler(ctx context.Context, name string, handler ClusterSettingHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*ClusterSetting); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *clusterSettingController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler ClusterSettingHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*ClusterSetting); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *clusterSettingController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler ClusterSettingHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*ClusterSetting); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *clusterSettingController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler ClusterSettingHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*ClusterSetting); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type clusterSettingFactory struct {
}

func (c clusterSettingFactory) Object() runtime.Object {
	return &ClusterSetting{}
}

func (c clusterSettingFactory) List() runtime.Object {
	return &ClusterSettingList{}
}

func (s *clusterSettingClient) Controller() ClusterSettingController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.clusterSettingControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ClusterSettingGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &clusterSettingController{
		GenericController: genericController,
	}

	s.client.clusterSettingControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type clusterSettingClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ClusterSettingController
}

func (s *clusterSettingClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *clusterSettingClient) Create(o *ClusterSetting) (*ClusterSetting, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ClusterSetting), err
}

func (s *clusterSettingClient) Get(name string, opts metav1.GetOptions) (*ClusterSetting, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ClusterSetting), err
}

func (s *clusterSettingClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ClusterSetting, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*ClusterSetting), err
}

func (s *clusterSettingClient) Update(o *ClusterSetting) (*ClusterSetting, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ClusterSetting), err
}

func (s *clusterSettingClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *clusterSettingClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *clusterSettingClient) List(opts metav1.ListOptions) (*ClusterSettingList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ClusterSettingList), err
}

func (s *clusterSettingClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *clusterSettingClient) Patch(o *ClusterSetting, patchType types.PatchType, data []byte, subresources ...string) (*ClusterSetting, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*ClusterSetting), err
}

func (s *clusterSettingClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *clusterSettingClient) AddHandler(ctx context.Context, name string, sync ClusterSettingHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *clusterSettingClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync ClusterSettingHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *clusterSettingClient) AddLifecycle(ctx context.Context, name string, lifecycle ClusterSettingLifecycle) {
	sync := NewClusterSettingLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *clusterSettingClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle ClusterSettingLifecycle) {
	sync := NewClusterSettingLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *clusterSettingClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ClusterSettingHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *clusterSettingClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync ClusterSettingHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *clusterSettingClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ClusterSettingLifecycle) {
	sync := NewClusterSettingLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *clusterSettingClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle ClusterSettingLifecycle) {
	sync := NewClusterSettingLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

type ClusterSettingIndexer func(obj *ClusterSetting) ([]string, error)

type ClusterSettingClientCache interface {
	Get(namespace, name string) (*ClusterSetting, error)
	List(namespace string, selector labels.Selector) ([]*ClusterSetting, error)

	Index(name string, indexer ClusterSettingIndexer)
	GetIndexed(name, key string) ([]*ClusterSetting, error)
}

type ClusterSettingClient interface {
	Create(*ClusterSetting) (*ClusterSetting, error)
	Get(namespace, name string, opts metav1.GetOptions) (*ClusterSetting, error)
	Update(*ClusterSetting) (*ClusterSetting, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*ClusterSettingList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() ClusterSettingClientCache

	OnCreate(ctx context.Context, name string, sync ClusterSettingChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync ClusterSettingChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync ClusterSettingChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() ClusterSettingInterface
}

type clusterSettingClientCache struct {
	client *clusterSettingClient2
}

type clusterSettingClient2 struct {
	iface      ClusterSettingInterface
	controller ClusterSettingController
}

func (n *clusterSettingClient2) Interface() ClusterSettingInterface {
	return n.iface
}

func (n *clusterSettingClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *clusterSettingClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *clusterSettingClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *clusterSettingClient2) Create(obj *ClusterSetting) (*ClusterSetting, error) {
	return n.iface.Create(obj)
}

func (n *clusterSettingClient2) Get(namespace, name string, opts metav1.GetOptions) (*ClusterSetting, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *clusterSettingClient2) Update(obj *ClusterSetting) (*ClusterSetting, error) {
	return n.iface.Update(obj)
}

func (n *clusterSettingClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *clusterSettingClient2) List(namespace string, opts metav1.ListOptions) (*ClusterSettingList, error) {
	return n.iface.List(opts)
}

func (n *clusterSettingClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *clusterSettingClientCache) Get(namespace, name string) (*ClusterSetting, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *clusterSettingClientCache) List(namespace string, selector labels.Selector) ([]*ClusterSetting, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *clusterSettingClient2) Cache() ClusterSettingClientCache {
	n.loadController()
	return &clusterSettingClientCache{
		client: n,
	}
}

func (n *clusterSettingClient2) OnCreate(ctx context.Context, name string, sync ClusterSettingChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &clusterSettingLifecycleDelegate{create: sync})
}

func (n *clusterSettingClient2) OnChange(ctx context.Context, name string, sync ClusterSettingChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &clusterSettingLifecycleDelegate{update: sync})
}

func (n *clusterSettingClient2) OnRemove(ctx context.Context, name string, sync ClusterSettingChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &clusterSettingLifecycleDelegate{remove: sync})
}

func (n *clusterSettingClientCache) Index(name string, indexer ClusterSettingIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*ClusterSetting); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *clusterSettingClientCache) GetIndexed(name, key string) ([]*ClusterSetting, error) {
	var result []*ClusterSetting
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*ClusterSetting); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *clusterSettingClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type clusterSettingLifecycleDelegate struct {
	create ClusterSettingChangeHandlerFunc
	update ClusterSettingChangeHandlerFunc
	remove ClusterSettingChangeHandlerFunc
}

func (n *clusterSettingLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *clusterSettingLifecycleDelegate) Create(obj *ClusterSetting) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *clusterSettingLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *clusterSettingLifecycleDelegate) Remove(obj *ClusterSetting) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *clusterSettingLifecycleDelegate) Updated(obj *ClusterSetting) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
