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
	HarborProjectGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "HarborProject",
	}
	HarborProjectResource = metav1.APIResource{
		Name:         "harborprojects",
		SingularName: "harborproject",
		Namespaced:   false,
		Kind:         HarborProjectGroupVersionKind.Kind,
	}

	HarborProjectGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "harborprojects",
	}
)

func init() {
	resource.Put(HarborProjectGroupVersionResource)
}

func NewHarborProject(namespace, name string, obj HarborProject) *HarborProject {
	obj.APIVersion, obj.Kind = HarborProjectGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type HarborProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HarborProject `json:"items"`
}

type HarborProjectHandlerFunc func(key string, obj *HarborProject) (runtime.Object, error)

type HarborProjectChangeHandlerFunc func(obj *HarborProject) (runtime.Object, error)

type HarborProjectLister interface {
	List(namespace string, selector labels.Selector) (ret []*HarborProject, err error)
	Get(namespace, name string) (*HarborProject, error)
}

type HarborProjectController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() HarborProjectLister
	AddHandler(ctx context.Context, name string, handler HarborProjectHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync HarborProjectHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler HarborProjectHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler HarborProjectHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type HarborProjectInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*HarborProject) (*HarborProject, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*HarborProject, error)
	Get(name string, opts metav1.GetOptions) (*HarborProject, error)
	Update(*HarborProject) (*HarborProject, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*HarborProjectList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() HarborProjectController
	AddHandler(ctx context.Context, name string, sync HarborProjectHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync HarborProjectHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle HarborProjectLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle HarborProjectLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync HarborProjectHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync HarborProjectHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle HarborProjectLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle HarborProjectLifecycle)
}

type harborProjectLister struct {
	controller *harborProjectController
}

func (l *harborProjectLister) List(namespace string, selector labels.Selector) (ret []*HarborProject, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*HarborProject))
	})
	return
}

func (l *harborProjectLister) Get(namespace, name string) (*HarborProject, error) {
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
			Group:    HarborProjectGroupVersionKind.Group,
			Resource: "harborProject",
		}, key)
	}
	return obj.(*HarborProject), nil
}

type harborProjectController struct {
	controller.GenericController
}

func (c *harborProjectController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *harborProjectController) Lister() HarborProjectLister {
	return &harborProjectLister{
		controller: c,
	}
}

func (c *harborProjectController) AddHandler(ctx context.Context, name string, handler HarborProjectHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*HarborProject); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *harborProjectController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler HarborProjectHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*HarborProject); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *harborProjectController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler HarborProjectHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*HarborProject); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *harborProjectController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler HarborProjectHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*HarborProject); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type harborProjectFactory struct {
}

func (c harborProjectFactory) Object() runtime.Object {
	return &HarborProject{}
}

func (c harborProjectFactory) List() runtime.Object {
	return &HarborProjectList{}
}

func (s *harborProjectClient) Controller() HarborProjectController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.harborProjectControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(HarborProjectGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &harborProjectController{
		GenericController: genericController,
	}

	s.client.harborProjectControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type harborProjectClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   HarborProjectController
}

func (s *harborProjectClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *harborProjectClient) Create(o *HarborProject) (*HarborProject, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*HarborProject), err
}

func (s *harborProjectClient) Get(name string, opts metav1.GetOptions) (*HarborProject, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*HarborProject), err
}

func (s *harborProjectClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*HarborProject, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*HarborProject), err
}

func (s *harborProjectClient) Update(o *HarborProject) (*HarborProject, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*HarborProject), err
}

func (s *harborProjectClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *harborProjectClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *harborProjectClient) List(opts metav1.ListOptions) (*HarborProjectList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*HarborProjectList), err
}

func (s *harborProjectClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *harborProjectClient) Patch(o *HarborProject, patchType types.PatchType, data []byte, subresources ...string) (*HarborProject, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*HarborProject), err
}

func (s *harborProjectClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *harborProjectClient) AddHandler(ctx context.Context, name string, sync HarborProjectHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *harborProjectClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync HarborProjectHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *harborProjectClient) AddLifecycle(ctx context.Context, name string, lifecycle HarborProjectLifecycle) {
	sync := NewHarborProjectLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *harborProjectClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle HarborProjectLifecycle) {
	sync := NewHarborProjectLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *harborProjectClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync HarborProjectHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *harborProjectClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync HarborProjectHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *harborProjectClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle HarborProjectLifecycle) {
	sync := NewHarborProjectLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *harborProjectClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle HarborProjectLifecycle) {
	sync := NewHarborProjectLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

type HarborProjectIndexer func(obj *HarborProject) ([]string, error)

type HarborProjectClientCache interface {
	Get(namespace, name string) (*HarborProject, error)
	List(namespace string, selector labels.Selector) ([]*HarborProject, error)

	Index(name string, indexer HarborProjectIndexer)
	GetIndexed(name, key string) ([]*HarborProject, error)
}

type HarborProjectClient interface {
	Create(*HarborProject) (*HarborProject, error)
	Get(namespace, name string, opts metav1.GetOptions) (*HarborProject, error)
	Update(*HarborProject) (*HarborProject, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*HarborProjectList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() HarborProjectClientCache

	OnCreate(ctx context.Context, name string, sync HarborProjectChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync HarborProjectChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync HarborProjectChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() HarborProjectInterface
}

type harborProjectClientCache struct {
	client *harborProjectClient2
}

type harborProjectClient2 struct {
	iface      HarborProjectInterface
	controller HarborProjectController
}

func (n *harborProjectClient2) Interface() HarborProjectInterface {
	return n.iface
}

func (n *harborProjectClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *harborProjectClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *harborProjectClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *harborProjectClient2) Create(obj *HarborProject) (*HarborProject, error) {
	return n.iface.Create(obj)
}

func (n *harborProjectClient2) Get(namespace, name string, opts metav1.GetOptions) (*HarborProject, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *harborProjectClient2) Update(obj *HarborProject) (*HarborProject, error) {
	return n.iface.Update(obj)
}

func (n *harborProjectClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *harborProjectClient2) List(namespace string, opts metav1.ListOptions) (*HarborProjectList, error) {
	return n.iface.List(opts)
}

func (n *harborProjectClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *harborProjectClientCache) Get(namespace, name string) (*HarborProject, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *harborProjectClientCache) List(namespace string, selector labels.Selector) ([]*HarborProject, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *harborProjectClient2) Cache() HarborProjectClientCache {
	n.loadController()
	return &harborProjectClientCache{
		client: n,
	}
}

func (n *harborProjectClient2) OnCreate(ctx context.Context, name string, sync HarborProjectChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &harborProjectLifecycleDelegate{create: sync})
}

func (n *harborProjectClient2) OnChange(ctx context.Context, name string, sync HarborProjectChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &harborProjectLifecycleDelegate{update: sync})
}

func (n *harborProjectClient2) OnRemove(ctx context.Context, name string, sync HarborProjectChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &harborProjectLifecycleDelegate{remove: sync})
}

func (n *harborProjectClientCache) Index(name string, indexer HarborProjectIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*HarborProject); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *harborProjectClientCache) GetIndexed(name, key string) ([]*HarborProject, error) {
	var result []*HarborProject
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*HarborProject); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *harborProjectClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type harborProjectLifecycleDelegate struct {
	create HarborProjectChangeHandlerFunc
	update HarborProjectChangeHandlerFunc
	remove HarborProjectChangeHandlerFunc
}

func (n *harborProjectLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *harborProjectLifecycleDelegate) Create(obj *HarborProject) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *harborProjectLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *harborProjectLifecycleDelegate) Remove(obj *HarborProject) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *harborProjectLifecycleDelegate) Updated(obj *HarborProject) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
