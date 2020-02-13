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
	PipelineTemplateGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "PipelineTemplate",
	}
	PipelineTemplateResource = metav1.APIResource{
		Name:         "pipelinetemplates",
		SingularName: "pipelinetemplate",
		Namespaced:   true,

		Kind: PipelineTemplateGroupVersionKind.Kind,
	}

	PipelineTemplateGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "pipelinetemplates",
	}
)

func init() {
	resource.Put(PipelineTemplateGroupVersionResource)
}

func NewPipelineTemplate(namespace, name string, obj PipelineTemplate) *PipelineTemplate {
	obj.APIVersion, obj.Kind = PipelineTemplateGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type PipelineTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PipelineTemplate `json:"items"`
}

type PipelineTemplateHandlerFunc func(key string, obj *PipelineTemplate) (runtime.Object, error)

type PipelineTemplateChangeHandlerFunc func(obj *PipelineTemplate) (runtime.Object, error)

type PipelineTemplateLister interface {
	List(namespace string, selector labels.Selector) (ret []*PipelineTemplate, err error)
	Get(namespace, name string) (*PipelineTemplate, error)
}

type PipelineTemplateController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() PipelineTemplateLister
	AddHandler(ctx context.Context, name string, handler PipelineTemplateHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync PipelineTemplateHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler PipelineTemplateHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler PipelineTemplateHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type PipelineTemplateInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*PipelineTemplate) (*PipelineTemplate, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*PipelineTemplate, error)
	Get(name string, opts metav1.GetOptions) (*PipelineTemplate, error)
	Update(*PipelineTemplate) (*PipelineTemplate, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*PipelineTemplateList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() PipelineTemplateController
	AddHandler(ctx context.Context, name string, sync PipelineTemplateHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync PipelineTemplateHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle PipelineTemplateLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle PipelineTemplateLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync PipelineTemplateHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync PipelineTemplateHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle PipelineTemplateLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle PipelineTemplateLifecycle)
}

type pipelineTemplateLister struct {
	controller *pipelineTemplateController
}

func (l *pipelineTemplateLister) List(namespace string, selector labels.Selector) (ret []*PipelineTemplate, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*PipelineTemplate))
	})
	return
}

func (l *pipelineTemplateLister) Get(namespace, name string) (*PipelineTemplate, error) {
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
			Group:    PipelineTemplateGroupVersionKind.Group,
			Resource: "pipelineTemplate",
		}, key)
	}
	return obj.(*PipelineTemplate), nil
}

type pipelineTemplateController struct {
	controller.GenericController
}

func (c *pipelineTemplateController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *pipelineTemplateController) Lister() PipelineTemplateLister {
	return &pipelineTemplateLister{
		controller: c,
	}
}

func (c *pipelineTemplateController) AddHandler(ctx context.Context, name string, handler PipelineTemplateHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*PipelineTemplate); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *pipelineTemplateController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler PipelineTemplateHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*PipelineTemplate); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *pipelineTemplateController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler PipelineTemplateHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*PipelineTemplate); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *pipelineTemplateController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler PipelineTemplateHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*PipelineTemplate); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type pipelineTemplateFactory struct {
}

func (c pipelineTemplateFactory) Object() runtime.Object {
	return &PipelineTemplate{}
}

func (c pipelineTemplateFactory) List() runtime.Object {
	return &PipelineTemplateList{}
}

func (s *pipelineTemplateClient) Controller() PipelineTemplateController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.pipelineTemplateControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(PipelineTemplateGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &pipelineTemplateController{
		GenericController: genericController,
	}

	s.client.pipelineTemplateControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type pipelineTemplateClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   PipelineTemplateController
}

func (s *pipelineTemplateClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *pipelineTemplateClient) Create(o *PipelineTemplate) (*PipelineTemplate, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*PipelineTemplate), err
}

func (s *pipelineTemplateClient) Get(name string, opts metav1.GetOptions) (*PipelineTemplate, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*PipelineTemplate), err
}

func (s *pipelineTemplateClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*PipelineTemplate, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*PipelineTemplate), err
}

func (s *pipelineTemplateClient) Update(o *PipelineTemplate) (*PipelineTemplate, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*PipelineTemplate), err
}

func (s *pipelineTemplateClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *pipelineTemplateClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *pipelineTemplateClient) List(opts metav1.ListOptions) (*PipelineTemplateList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*PipelineTemplateList), err
}

func (s *pipelineTemplateClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *pipelineTemplateClient) Patch(o *PipelineTemplate, patchType types.PatchType, data []byte, subresources ...string) (*PipelineTemplate, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*PipelineTemplate), err
}

func (s *pipelineTemplateClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *pipelineTemplateClient) AddHandler(ctx context.Context, name string, sync PipelineTemplateHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *pipelineTemplateClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync PipelineTemplateHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *pipelineTemplateClient) AddLifecycle(ctx context.Context, name string, lifecycle PipelineTemplateLifecycle) {
	sync := NewPipelineTemplateLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *pipelineTemplateClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle PipelineTemplateLifecycle) {
	sync := NewPipelineTemplateLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *pipelineTemplateClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync PipelineTemplateHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *pipelineTemplateClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync PipelineTemplateHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *pipelineTemplateClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle PipelineTemplateLifecycle) {
	sync := NewPipelineTemplateLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *pipelineTemplateClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle PipelineTemplateLifecycle) {
	sync := NewPipelineTemplateLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

type PipelineTemplateIndexer func(obj *PipelineTemplate) ([]string, error)

type PipelineTemplateClientCache interface {
	Get(namespace, name string) (*PipelineTemplate, error)
	List(namespace string, selector labels.Selector) ([]*PipelineTemplate, error)

	Index(name string, indexer PipelineTemplateIndexer)
	GetIndexed(name, key string) ([]*PipelineTemplate, error)
}

type PipelineTemplateClient interface {
	Create(*PipelineTemplate) (*PipelineTemplate, error)
	Get(namespace, name string, opts metav1.GetOptions) (*PipelineTemplate, error)
	Update(*PipelineTemplate) (*PipelineTemplate, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*PipelineTemplateList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() PipelineTemplateClientCache

	OnCreate(ctx context.Context, name string, sync PipelineTemplateChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync PipelineTemplateChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync PipelineTemplateChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() PipelineTemplateInterface
}

type pipelineTemplateClientCache struct {
	client *pipelineTemplateClient2
}

type pipelineTemplateClient2 struct {
	iface      PipelineTemplateInterface
	controller PipelineTemplateController
}

func (n *pipelineTemplateClient2) Interface() PipelineTemplateInterface {
	return n.iface
}

func (n *pipelineTemplateClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *pipelineTemplateClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *pipelineTemplateClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *pipelineTemplateClient2) Create(obj *PipelineTemplate) (*PipelineTemplate, error) {
	return n.iface.Create(obj)
}

func (n *pipelineTemplateClient2) Get(namespace, name string, opts metav1.GetOptions) (*PipelineTemplate, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *pipelineTemplateClient2) Update(obj *PipelineTemplate) (*PipelineTemplate, error) {
	return n.iface.Update(obj)
}

func (n *pipelineTemplateClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *pipelineTemplateClient2) List(namespace string, opts metav1.ListOptions) (*PipelineTemplateList, error) {
	return n.iface.List(opts)
}

func (n *pipelineTemplateClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *pipelineTemplateClientCache) Get(namespace, name string) (*PipelineTemplate, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *pipelineTemplateClientCache) List(namespace string, selector labels.Selector) ([]*PipelineTemplate, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *pipelineTemplateClient2) Cache() PipelineTemplateClientCache {
	n.loadController()
	return &pipelineTemplateClientCache{
		client: n,
	}
}

func (n *pipelineTemplateClient2) OnCreate(ctx context.Context, name string, sync PipelineTemplateChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &pipelineTemplateLifecycleDelegate{create: sync})
}

func (n *pipelineTemplateClient2) OnChange(ctx context.Context, name string, sync PipelineTemplateChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &pipelineTemplateLifecycleDelegate{update: sync})
}

func (n *pipelineTemplateClient2) OnRemove(ctx context.Context, name string, sync PipelineTemplateChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &pipelineTemplateLifecycleDelegate{remove: sync})
}

func (n *pipelineTemplateClientCache) Index(name string, indexer PipelineTemplateIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*PipelineTemplate); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *pipelineTemplateClientCache) GetIndexed(name, key string) ([]*PipelineTemplate, error) {
	var result []*PipelineTemplate
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*PipelineTemplate); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *pipelineTemplateClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type pipelineTemplateLifecycleDelegate struct {
	create PipelineTemplateChangeHandlerFunc
	update PipelineTemplateChangeHandlerFunc
	remove PipelineTemplateChangeHandlerFunc
}

func (n *pipelineTemplateLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *pipelineTemplateLifecycleDelegate) Create(obj *PipelineTemplate) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *pipelineTemplateLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *pipelineTemplateLifecycleDelegate) Remove(obj *PipelineTemplate) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *pipelineTemplateLifecycleDelegate) Updated(obj *PipelineTemplate) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
