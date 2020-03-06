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
	HarborRepositoryGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "HarborRepository",
	}
	HarborRepositoryResource = metav1.APIResource{
		Name:         "harborrepositories",
		SingularName: "harborrepository",
		Namespaced:   false,
		Kind:         HarborRepositoryGroupVersionKind.Kind,
	}

	HarborRepositoryGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "harborrepositories",
	}
)

func init() {
	resource.Put(HarborRepositoryGroupVersionResource)
}

func NewHarborRepository(namespace, name string, obj HarborRepository) *HarborRepository {
	obj.APIVersion, obj.Kind = HarborRepositoryGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type HarborRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HarborRepository `json:"items"`
}

type HarborRepositoryHandlerFunc func(key string, obj *HarborRepository) (runtime.Object, error)

type HarborRepositoryChangeHandlerFunc func(obj *HarborRepository) (runtime.Object, error)

type HarborRepositoryLister interface {
	List(namespace string, selector labels.Selector) (ret []*HarborRepository, err error)
	Get(namespace, name string) (*HarborRepository, error)
}

type HarborRepositoryController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() HarborRepositoryLister
	AddHandler(ctx context.Context, name string, handler HarborRepositoryHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync HarborRepositoryHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler HarborRepositoryHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler HarborRepositoryHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type HarborRepositoryInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*HarborRepository) (*HarborRepository, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*HarborRepository, error)
	Get(name string, opts metav1.GetOptions) (*HarborRepository, error)
	Update(*HarborRepository) (*HarborRepository, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*HarborRepositoryList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() HarborRepositoryController
	AddHandler(ctx context.Context, name string, sync HarborRepositoryHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync HarborRepositoryHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle HarborRepositoryLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle HarborRepositoryLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync HarborRepositoryHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync HarborRepositoryHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle HarborRepositoryLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle HarborRepositoryLifecycle)
}

type harborRepositoryLister struct {
	controller *harborRepositoryController
}

func (l *harborRepositoryLister) List(namespace string, selector labels.Selector) (ret []*HarborRepository, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*HarborRepository))
	})
	return
}

func (l *harborRepositoryLister) Get(namespace, name string) (*HarborRepository, error) {
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
			Group:    HarborRepositoryGroupVersionKind.Group,
			Resource: "harborRepository",
		}, key)
	}
	return obj.(*HarborRepository), nil
}

type harborRepositoryController struct {
	controller.GenericController
}

func (c *harborRepositoryController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *harborRepositoryController) Lister() HarborRepositoryLister {
	return &harborRepositoryLister{
		controller: c,
	}
}

func (c *harborRepositoryController) AddHandler(ctx context.Context, name string, handler HarborRepositoryHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*HarborRepository); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *harborRepositoryController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler HarborRepositoryHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*HarborRepository); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *harborRepositoryController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler HarborRepositoryHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*HarborRepository); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *harborRepositoryController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler HarborRepositoryHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*HarborRepository); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type harborRepositoryFactory struct {
}

func (c harborRepositoryFactory) Object() runtime.Object {
	return &HarborRepository{}
}

func (c harborRepositoryFactory) List() runtime.Object {
	return &HarborRepositoryList{}
}

func (s *harborRepositoryClient) Controller() HarborRepositoryController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.harborRepositoryControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(HarborRepositoryGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &harborRepositoryController{
		GenericController: genericController,
	}

	s.client.harborRepositoryControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type harborRepositoryClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   HarborRepositoryController
}

func (s *harborRepositoryClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *harborRepositoryClient) Create(o *HarborRepository) (*HarborRepository, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*HarborRepository), err
}

func (s *harborRepositoryClient) Get(name string, opts metav1.GetOptions) (*HarborRepository, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*HarborRepository), err
}

func (s *harborRepositoryClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*HarborRepository, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*HarborRepository), err
}

func (s *harborRepositoryClient) Update(o *HarborRepository) (*HarborRepository, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*HarborRepository), err
}

func (s *harborRepositoryClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *harborRepositoryClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *harborRepositoryClient) List(opts metav1.ListOptions) (*HarborRepositoryList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*HarborRepositoryList), err
}

func (s *harborRepositoryClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *harborRepositoryClient) Patch(o *HarborRepository, patchType types.PatchType, data []byte, subresources ...string) (*HarborRepository, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*HarborRepository), err
}

func (s *harborRepositoryClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *harborRepositoryClient) AddHandler(ctx context.Context, name string, sync HarborRepositoryHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *harborRepositoryClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync HarborRepositoryHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *harborRepositoryClient) AddLifecycle(ctx context.Context, name string, lifecycle HarborRepositoryLifecycle) {
	sync := NewHarborRepositoryLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *harborRepositoryClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle HarborRepositoryLifecycle) {
	sync := NewHarborRepositoryLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *harborRepositoryClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync HarborRepositoryHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *harborRepositoryClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync HarborRepositoryHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *harborRepositoryClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle HarborRepositoryLifecycle) {
	sync := NewHarborRepositoryLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *harborRepositoryClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle HarborRepositoryLifecycle) {
	sync := NewHarborRepositoryLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

type HarborRepositoryIndexer func(obj *HarborRepository) ([]string, error)

type HarborRepositoryClientCache interface {
	Get(namespace, name string) (*HarborRepository, error)
	List(namespace string, selector labels.Selector) ([]*HarborRepository, error)

	Index(name string, indexer HarborRepositoryIndexer)
	GetIndexed(name, key string) ([]*HarborRepository, error)
}

type HarborRepositoryClient interface {
	Create(*HarborRepository) (*HarborRepository, error)
	Get(namespace, name string, opts metav1.GetOptions) (*HarborRepository, error)
	Update(*HarborRepository) (*HarborRepository, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*HarborRepositoryList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() HarborRepositoryClientCache

	OnCreate(ctx context.Context, name string, sync HarborRepositoryChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync HarborRepositoryChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync HarborRepositoryChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() HarborRepositoryInterface
}

type harborRepositoryClientCache struct {
	client *harborRepositoryClient2
}

type harborRepositoryClient2 struct {
	iface      HarborRepositoryInterface
	controller HarborRepositoryController
}

func (n *harborRepositoryClient2) Interface() HarborRepositoryInterface {
	return n.iface
}

func (n *harborRepositoryClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *harborRepositoryClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *harborRepositoryClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *harborRepositoryClient2) Create(obj *HarborRepository) (*HarborRepository, error) {
	return n.iface.Create(obj)
}

func (n *harborRepositoryClient2) Get(namespace, name string, opts metav1.GetOptions) (*HarborRepository, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *harborRepositoryClient2) Update(obj *HarborRepository) (*HarborRepository, error) {
	return n.iface.Update(obj)
}

func (n *harborRepositoryClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *harborRepositoryClient2) List(namespace string, opts metav1.ListOptions) (*HarborRepositoryList, error) {
	return n.iface.List(opts)
}

func (n *harborRepositoryClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *harborRepositoryClientCache) Get(namespace, name string) (*HarborRepository, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *harborRepositoryClientCache) List(namespace string, selector labels.Selector) ([]*HarborRepository, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *harborRepositoryClient2) Cache() HarborRepositoryClientCache {
	n.loadController()
	return &harborRepositoryClientCache{
		client: n,
	}
}

func (n *harborRepositoryClient2) OnCreate(ctx context.Context, name string, sync HarborRepositoryChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &harborRepositoryLifecycleDelegate{create: sync})
}

func (n *harborRepositoryClient2) OnChange(ctx context.Context, name string, sync HarborRepositoryChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &harborRepositoryLifecycleDelegate{update: sync})
}

func (n *harborRepositoryClient2) OnRemove(ctx context.Context, name string, sync HarborRepositoryChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &harborRepositoryLifecycleDelegate{remove: sync})
}

func (n *harborRepositoryClientCache) Index(name string, indexer HarborRepositoryIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*HarborRepository); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *harborRepositoryClientCache) GetIndexed(name, key string) ([]*HarborRepository, error) {
	var result []*HarborRepository
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*HarborRepository); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *harborRepositoryClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type harborRepositoryLifecycleDelegate struct {
	create HarborRepositoryChangeHandlerFunc
	update HarborRepositoryChangeHandlerFunc
	remove HarborRepositoryChangeHandlerFunc
}

func (n *harborRepositoryLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *harborRepositoryLifecycleDelegate) Create(obj *HarborRepository) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *harborRepositoryLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *harborRepositoryLifecycleDelegate) Remove(obj *HarborRepository) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *harborRepositoryLifecycleDelegate) Updated(obj *HarborRepository) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
