/*
Copyright 2022 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package informergen

import (
	"io"
	"text/template"

	"github.com/kcp-dev/code-generator/v2/pkg/parser"
)

type Factory struct {
	// Groups are the groups in this informer factory.
	Groups []parser.Group

	// PackagePath is the package under which these informers will be exposed.
	// e.g. "github.com/kcp-dev/client-go/clients/informers"
	// TODO(skuznets) we should be able to figure this out from the output dir, ideally
	PackagePath string

	// ClientsetPackagePath is the package under which the cluster-aware client-set will be exposed.
	// e.g. "github.com/kcp-dev/client-go/clients/clientset/versioned"
	// TODO(skuznets) we should be able to figure this out from the output dir, ideally
	ClientsetPackagePath string

	// SingleClusterClientPackagePath is the root directory under which single-cluster-aware clients exist.
	// e.g. "k8s.io/client-go/kubernetes"
	SingleClusterClientPackagePath string `marker:""`

	// SingleClusterInformerPackagePath is the package under which the cluster-unaware listers are exposed.
	// e.g. "k8s.io/client-go/informers"
	SingleClusterInformerPackagePath string
}

func (f *Factory) WriteContent(w io.Writer) error {
	templ, err := template.New("factory").Funcs(templateFuncs).Parse(sharedInformerFactoryStruct)
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"groups":                           f.Groups,
		"packagePath":                      f.PackagePath,
		"clientsetPackagePath":             f.ClientsetPackagePath,
		"singleClusterClientPackagePath":   f.SingleClusterClientPackagePath,
		"singleClusterInformerPackagePath": f.SingleClusterInformerPackagePath,
		"useUpstreamInterfaces":            f.SingleClusterInformerPackagePath != "",
	}
	return templ.Execute(w, m)
}

var sharedInformerFactoryStruct = `
// Code generated by kcp code-generator. DO NOT EDIT.

package informers

import (
	"reflect"
	"sync"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"

	clientset "{{.clientsetPackagePath}}"
	{{if not .useUpstreamInterfaces -}}	
	scopedclientset "{{.singleClusterClientPackagePath}}"
	{{end -}}
	{{if .useUpstreamInterfaces -}}
	upstreaminformers "{{.singleClusterInformerPackagePath}}"
	{{end -}}

{{range .groups}}	{{.PackageName}}informers "{{$.packagePath}}/{{.PackageName}}"
{{end -}}

	"{{.packagePath}}/internalinterfaces"
)

// SharedInformerOption defines the functional option type for SharedInformerFactory.
type SharedInformerOption func(*SharedInformerOptions) *SharedInformerOptions

type SharedInformerOptions struct {
	customResync map[reflect.Type]time.Duration
	tweakListOptions internalinterfaces.TweakListOptionsFunc
    transform cache.TransformFunc
	{{if not .useUpstreamInterfaces -}}
	namespace string
	{{end -}}
}

type sharedInformerFactory struct {
	client clientset.ClusterInterface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	lock sync.Mutex
	defaultResync time.Duration
	customResync map[reflect.Type]time.Duration
    transform cache.TransformFunc

	informers map[reflect.Type]kcpcache.ScopeableSharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
	// wg tracks how many goroutines were started.
	wg sync.WaitGroup
	// shuttingDown is true when Shutdown has been called. It may still be running
	// because it needs to wait for goroutines.
	shuttingDown bool
}

// WithCustomResyncConfig sets a custom resync period for the specified informer types.
func WithCustomResyncConfig(resyncConfig map[metav1.Object]time.Duration) SharedInformerOption {
	return func(opts *SharedInformerOptions) *SharedInformerOptions {
		for k, v := range resyncConfig {
			opts.customResync[reflect.TypeOf(k)] = v
		}
		return opts
	}
}

// WithTweakListOptions sets a custom filter on all listers of the configured SharedInformerFactory.
func WithTweakListOptions(tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerOption {
	return func(opts *SharedInformerOptions) *SharedInformerOptions {
		opts.tweakListOptions = tweakListOptions
		return opts
	}
}

// WithTransform sets a transform on all informers.
func WithTransform(transform cache.TransformFunc) SharedInformerOption {
	return func(opts *SharedInformerOptions) *SharedInformerOptions {
	    opts.transform = transform
		return opts
	}
}

// NewSharedInformerFactory constructs a new instance of SharedInformerFactory for all namespaces.
func NewSharedInformerFactory(client clientset.ClusterInterface, defaultResync time.Duration) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync)
}

// NewSharedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedInformerFactoryWithOptions(client clientset.ClusterInterface, defaultResync time.Duration, options ...SharedInformerOption) SharedInformerFactory {
	factory := &sharedInformerFactory{
		client:           client,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]kcpcache.ScopeableSharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
		customResync:     make(map[reflect.Type]time.Duration),
	}

	opts := &SharedInformerOptions{
		customResync:     make(map[reflect.Type]time.Duration),
	}

	// Apply all options
	for _, opt := range options {
		opts = opt(opts)
	}

	// Forward options to the factory
	factory.customResync = opts.customResync
	factory.tweakListOptions = opts.tweakListOptions
    factory.transform = opts.transform

	return factory
}

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.shuttingDown {
		return
	}

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			f.wg.Add(1)
			// We need a new variable in each loop iteration,
			// otherwise the goroutine would use the loop variable
			// and that keeps changing.
			informer := informer
			go func() {
				defer f.wg.Done()
				informer.Run(stopCh)
			}()
			f.startedInformers[informerType] = true
		}
	}
}

func (f *sharedInformerFactory) Shutdown() {
	f.lock.Lock()
	f.shuttingDown = true
	f.lock.Unlock()

	// Will return immediately if there is nothing to wait for.
	f.wg.Wait()
}

// WaitForCacheSync waits for all started informers' cache were synced.
func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	informers := func()map[reflect.Type]kcpcache.ScopeableSharedIndexInformer{
               f.lock.Lock()
               defer f.lock.Unlock()

               informers := map[reflect.Type]kcpcache.ScopeableSharedIndexInformer{}
               for informerType, informer := range f.informers {
                       if f.startedInformers[informerType] {
                               informers[informerType] = informer
                       }
               }
               return informers
       }()

       res := map[reflect.Type]bool{}
       for informType, informer := range informers {
               res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
       }
       return res
}

// InformerFor returns the SharedIndexInformer for obj.
func (f *sharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) kcpcache.ScopeableSharedIndexInformer {
  f.lock.Lock()
  defer f.lock.Unlock()

  informerType := reflect.TypeOf(obj)
  informer, exists := f.informers[informerType]
  if exists {
    return informer
  }

  resyncPeriod, exists := f.customResync[informerType]
  if !exists {
    resyncPeriod = f.defaultResync
  }

  informer = newFunc(f.client, resyncPeriod)
  f.informers[informerType] = informer

  return informer
}

type ScopedDynamicSharedInformerFactory interface {
	// ForResource gives generic access to a shared informer of the matching type.
	ForResource(resource schema.GroupVersionResource) ({{if .useUpstreamInterfaces}}upstreaminformers.{{end}}GenericInformer, error)
	
	// Start initializes all requested informers. They are handled in goroutines
	// which run until the stop channel gets closed.
	Start(stopCh <-chan struct{})
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
//
// It is typically used like this:
//
//	ctx, cancel := context.Background()
//	defer cancel()
//	factory := NewSharedInformerFactoryWithOptions(client, resyncPeriod)
//	defer factory.Shutdown()    // Returns immediately if nothing was started.
//	genericInformer := factory.ForResource(resource)
//	typedInformer := factory.SomeAPIGroup().V1().SomeType()
//	factory.Start(ctx.Done())          // Start processing these informers.
//	synced := factory.WaitForCacheSync(ctx.Done())
//	for v, ok := range synced {
//	    if !ok {
//	        fmt.Fprintf(os.Stderr, "caches failed to sync: %v", v)
//	        return
//	    }
//	}
//
//	// Creating informers can also be created after Start, but then
//	// Start must be called again:
//	anotherGenericInformer := factory.ForResource(resource)
//	factory.Start(ctx.Done())
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory

	Cluster(logicalcluster.Name) ScopedDynamicSharedInformerFactory

	// Start initializes all requested informers. They are handled in goroutines
	// which run until the stop channel gets closed.
	Start(stopCh <-chan struct{})

	// Shutdown marks a factory as shutting down. At that point no new
	// informers can be started anymore and Start will return without
	// doing anything.
	//
	// In addition, Shutdown blocks until all goroutines have terminated. For that
	// to happen, the close channel(s) that they were started with must be closed,
	// either before Shutdown gets called or while it is waiting.
	//
	// Shutdown may be called multiple times, even concurrently. All such calls will
	// block until all goroutines have terminated.
	Shutdown()

	// ForResource gives generic access to a shared informer of the matching type.
	ForResource(resource schema.GroupVersionResource) (GenericClusterInformer, error)

	// WaitForCacheSync blocks until all started informers' caches were synced
	// or the stop channel gets closed.
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

	// InformerFor returns the SharedIndexInformer for obj.
	InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) kcpcache.ScopeableSharedIndexInformer

{{range .groups}}	{{.GroupGoName}}() {{.PackageName}}informers.ClusterInterface
{{end -}}
}

{{range .groups}}
func (f *sharedInformerFactory) {{.GroupGoName}}() {{.PackageName}}informers.ClusterInterface {
  return {{.PackageName}}informers.New(f, f.tweakListOptions)
}
{{end}}

func (f *sharedInformerFactory) Cluster(clusterName logicalcluster.Name) ScopedDynamicSharedInformerFactory {
	return &scopedDynamicSharedInformerFactory{
		sharedInformerFactory: f,
		clusterName: clusterName,
	}
}

type scopedDynamicSharedInformerFactory struct {
	*sharedInformerFactory
	clusterName logicalcluster.Name
}

func (f *scopedDynamicSharedInformerFactory) ForResource(resource schema.GroupVersionResource) ({{if .useUpstreamInterfaces}}upstreaminformers.{{end}}GenericInformer, error) {
	clusterInformer, err := f.sharedInformerFactory.ForResource(resource)
	if err != nil {
		return nil, err
	}
	return clusterInformer.Cluster(f.clusterName), nil 
}

func (f *scopedDynamicSharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.sharedInformerFactory.Start(stopCh)
}

{{if not .useUpstreamInterfaces -}}
// WithNamespace limits the SharedInformerFactory to the specified namespace.
func WithNamespace(namespace string) SharedInformerOption {
	return func(opts *SharedInformerOptions) *SharedInformerOptions {
		opts.namespace = namespace
		return opts
	}
}


type sharedScopedInformerFactory struct {
	client scopedclientset.Interface
	namespace string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	lock sync.Mutex
	defaultResync time.Duration
	customResync map[reflect.Type]time.Duration
    transform cache.TransformFunc

	informers map[reflect.Type]cache.SharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
}

// NewSharedScopedInformerFactory constructs a new instance of SharedInformerFactory for some or all namespaces.
func NewSharedScopedInformerFactory(client scopedclientset.Interface, defaultResync time.Duration, namespace string) SharedScopedInformerFactory {
	return NewSharedScopedInformerFactoryWithOptions(client, defaultResync, WithNamespace(namespace))
}

// NewSharedScopedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedScopedInformerFactoryWithOptions(client scopedclientset.Interface, defaultResync time.Duration, options ...SharedInformerOption) SharedScopedInformerFactory {
	factory := &sharedScopedInformerFactory{
		client:           client,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]cache.SharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
		customResync:     make(map[reflect.Type]time.Duration),
	}

	opts := &SharedInformerOptions{
		customResync:     make(map[reflect.Type]time.Duration),
	}

	// Apply all options
	for _, opt := range options {
		opts = opt(opts)
	}

	// Forward options to the factory
	factory.customResync = opts.customResync
	factory.tweakListOptions = opts.tweakListOptions
	factory.namespace = opts.namespace

	return factory
}

// Start initializes all requested informers.
func (f *sharedScopedInformerFactory) Start(stopCh <-chan struct{}) {
  f.lock.Lock()
  defer f.lock.Unlock()

  for informerType, informer := range f.informers {
    if !f.startedInformers[informerType] {
      go informer.Run(stopCh)
      f.startedInformers[informerType] = true
    }
  }
}

// WaitForCacheSync waits for all started informers' cache were synced.
func (f *sharedScopedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	informers := func()map[reflect.Type]cache.SharedIndexInformer{
               f.lock.Lock()
               defer f.lock.Unlock()

               informers := map[reflect.Type]cache.SharedIndexInformer{}
               for informerType, informer := range f.informers {
                       if f.startedInformers[informerType] {
                               informers[informerType] = informer
                       }
               }
               return informers
       }()

       res := map[reflect.Type]bool{}
       for informType, informer := range informers {
               res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
       }
       return res
}

// InformerFor returns the SharedIndexInformer for obj.
func (f *sharedScopedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewScopedInformerFunc) cache.SharedIndexInformer {
  f.lock.Lock()
  defer f.lock.Unlock()

  informerType := reflect.TypeOf(obj)
  informer, exists := f.informers[informerType]
  if exists {
    return informer
  }

  resyncPeriod, exists := f.customResync[informerType]
  if !exists {
    resyncPeriod = f.defaultResync
  }

  informer = newFunc(f.client, resyncPeriod)
  informer.SetTransform(f.transform)
  f.informers[informerType] = informer

  return informer
}

// SharedScopedInformerFactory provides shared informers for resources in all known
// API group versions, scoped to one workspace.
type SharedScopedInformerFactory interface {
	internalinterfaces.SharedScopedInformerFactory
	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

{{range .groups}}	{{.GroupGoName}}() {{.PackageName}}informers.Interface
{{end -}}
}


{{range .groups}}
func (f *sharedScopedInformerFactory) {{.GroupGoName}}() {{.PackageName}}informers.Interface {
  return {{.PackageName}}informers.NewScoped(f, f.namespace, f.tweakListOptions)
}
{{end}}
{{end}}
`
