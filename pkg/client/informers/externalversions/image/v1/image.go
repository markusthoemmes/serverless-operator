// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	versioned "github.com/openshift-knative/serverless-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openshift-knative/serverless-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift-knative/serverless-operator/pkg/client/listers/image/v1"
	imagev1 "github.com/openshift/api/image/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ImageInformer provides access to a shared informer and lister for
// Images.
type ImageInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ImageLister
}

type imageInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewImageInformer constructs a new informer for Image type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewImageInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredImageInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredImageInformer constructs a new informer for Image type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredImageInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ImageV1().Images().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ImageV1().Images().Watch(context.TODO(), options)
			},
		},
		&imagev1.Image{},
		resyncPeriod,
		indexers,
	)
}

func (f *imageInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredImageInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *imageInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&imagev1.Image{}, f.defaultInformer)
}

func (f *imageInformer) Lister() v1.ImageLister {
	return v1.NewImageLister(f.Informer().GetIndexer())
}
