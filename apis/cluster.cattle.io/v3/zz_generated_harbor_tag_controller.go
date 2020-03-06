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
	HarborTagGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "HarborTag",
	}
	HarborTagResource = metav1.APIResource{
		Name:         "harbortags",
		SingularName: "harbortag",
		Namespaced:   false,
		Kind:         HarborTagGroupVersionKind.Kind,
	}

	HarborTagGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "harbortags",
	}
)

func init() {
	resource.Put(HarborTagGroupVersionResource)
}

func NewHarborTag(namespace, name string, obj HarborTag) *HarborTag {
	obj.APIVersion, obj.Kind = HarborTagGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type HarborTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HarborTag `json:"items"`
}

type HarborTagHandlerFunc func(key string, obj *HarborTag) (runtime.Object, error)

type HarborTagChangeHandlerFunc func(obj *HarborTag) (runtime.Object, error)

type HarborTagLister interface {
	List(namespace string, selector labels.Selector) (ret []*HarborTag, err error)
	Get(namespace, name string) (*HarborTag, error)
}

type HarborTagController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() HarborTagLister
	AddHandler(ctx context.Context, name string, handler HarborTagHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync HarborTagHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler HarborTagHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler HarborTagHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type HarborTagInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*HarborTag) (*HarborTag, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*HarborTag, error)
	Get(name string, opts metav1.GetOptions) (*HarborTag, error)
	Update(*HarborTag) (*HarborTag, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*HarborTagList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() HarborTagController
	AddHandler(ctx context.Context, name string, sync HarborTagHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync HarborTagHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle HarborTagLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle HarborTagLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync HarborTagHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync HarborTagHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle HarborTagLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle HarborTagLifecycle)
}

type harborTagLister struct {
	controller *harborTagController
}

func (l *harborTagLister) List(namespace string, selector labels.Selector) (ret []*HarborTag, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*HarborTag))
	})
	return
}

func (l *harborTagLister) Get(namespace, name string) (*HarborTag, error) {
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
			Group:    HarborTagGroupVersionKind.Group,
			Resource: "harborTag",
		}, key)
	}
	return obj.(*HarborTag), nil
}

type harborTagController struct {
	controller.GenericController
}

func (c *harborTagController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *harborTagController) Lister() HarborTagLister {
	return &harborTagLister{
		controller: c,
	}
}

func (c *harborTagController) AddHandler(ctx context.Context, name string, handler HarborTagHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*HarborTag); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *harborTagController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler HarborTagHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*HarborTag); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *harborTagController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler HarborTagHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*HarborTag); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *harborTagController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler HarborTagHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*HarborTag); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type harborTagFactory struct {
}

func (c harborTagFactory) Object() runtime.Object {
	return &HarborTag{}
}

func (c harborTagFactory) List() runtime.Object {
	return &HarborTagList{}
}

func (s *harborTagClient) Controller() HarborTagController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.harborTagControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(HarborTagGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &harborTagController{
		GenericController: genericController,
	}

	s.client.harborTagControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type harborTagClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   HarborTagController
}

func (s *harborTagClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *harborTagClient) Create(o *HarborTag) (*HarborTag, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*HarborTag), err
}

func (s *harborTagClient) Get(name string, opts metav1.GetOptions) (*HarborTag, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*HarborTag), err
}

func (s *harborTagClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*HarborTag, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*HarborTag), err
}

func (s *harborTagClient) Update(o *HarborTag) (*HarborTag, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*HarborTag), err
}

func (s *harborTagClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *harborTagClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *harborTagClient) List(opts metav1.ListOptions) (*HarborTagList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*HarborTagList), err
}

func (s *harborTagClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *harborTagClient) Patch(o *HarborTag, patchType types.PatchType, data []byte, subresources ...string) (*HarborTag, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*HarborTag), err
}

func (s *harborTagClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *harborTagClient) AddHandler(ctx context.Context, name string, sync HarborTagHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *harborTagClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync HarborTagHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *harborTagClient) AddLifecycle(ctx context.Context, name string, lifecycle HarborTagLifecycle) {
	sync := NewHarborTagLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *harborTagClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle HarborTagLifecycle) {
	sync := NewHarborTagLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *harborTagClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync HarborTagHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *harborTagClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync HarborTagHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *harborTagClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle HarborTagLifecycle) {
	sync := NewHarborTagLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *harborTagClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle HarborTagLifecycle) {
	sync := NewHarborTagLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

type HarborTagIndexer func(obj *HarborTag) ([]string, error)

type HarborTagClientCache interface {
	Get(namespace, name string) (*HarborTag, error)
	List(namespace string, selector labels.Selector) ([]*HarborTag, error)

	Index(name string, indexer HarborTagIndexer)
	GetIndexed(name, key string) ([]*HarborTag, error)
}

type HarborTagClient interface {
	Create(*HarborTag) (*HarborTag, error)
	Get(namespace, name string, opts metav1.GetOptions) (*HarborTag, error)
	Update(*HarborTag) (*HarborTag, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*HarborTagList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() HarborTagClientCache

	OnCreate(ctx context.Context, name string, sync HarborTagChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync HarborTagChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync HarborTagChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() HarborTagInterface
}

type harborTagClientCache struct {
	client *harborTagClient2
}

type harborTagClient2 struct {
	iface      HarborTagInterface
	controller HarborTagController
}

func (n *harborTagClient2) Interface() HarborTagInterface {
	return n.iface
}

func (n *harborTagClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *harborTagClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *harborTagClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *harborTagClient2) Create(obj *HarborTag) (*HarborTag, error) {
	return n.iface.Create(obj)
}

func (n *harborTagClient2) Get(namespace, name string, opts metav1.GetOptions) (*HarborTag, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *harborTagClient2) Update(obj *HarborTag) (*HarborTag, error) {
	return n.iface.Update(obj)
}

func (n *harborTagClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *harborTagClient2) List(namespace string, opts metav1.ListOptions) (*HarborTagList, error) {
	return n.iface.List(opts)
}

func (n *harborTagClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *harborTagClientCache) Get(namespace, name string) (*HarborTag, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *harborTagClientCache) List(namespace string, selector labels.Selector) ([]*HarborTag, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *harborTagClient2) Cache() HarborTagClientCache {
	n.loadController()
	return &harborTagClientCache{
		client: n,
	}
}

func (n *harborTagClient2) OnCreate(ctx context.Context, name string, sync HarborTagChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &harborTagLifecycleDelegate{create: sync})
}

func (n *harborTagClient2) OnChange(ctx context.Context, name string, sync HarborTagChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &harborTagLifecycleDelegate{update: sync})
}

func (n *harborTagClient2) OnRemove(ctx context.Context, name string, sync HarborTagChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &harborTagLifecycleDelegate{remove: sync})
}

func (n *harborTagClientCache) Index(name string, indexer HarborTagIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*HarborTag); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *harborTagClientCache) GetIndexed(name, key string) ([]*HarborTag, error) {
	var result []*HarborTag
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*HarborTag); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *harborTagClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type harborTagLifecycleDelegate struct {
	create HarborTagChangeHandlerFunc
	update HarborTagChangeHandlerFunc
	remove HarborTagChangeHandlerFunc
}

func (n *harborTagLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *harborTagLifecycleDelegate) Create(obj *HarborTag) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *harborTagLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *harborTagLifecycleDelegate) Remove(obj *HarborTag) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *harborTagLifecycleDelegate) Updated(obj *HarborTag) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
