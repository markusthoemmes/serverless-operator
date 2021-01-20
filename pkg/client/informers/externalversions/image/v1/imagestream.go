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

// ImageStreamInformer provides access to a shared informer and lister for
// ImageStreams.
type ImageStreamInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ImageStreamLister
}

type imageStreamInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewImageStreamInformer constructs a new informer for ImageStream type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewImageStreamInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredImageStreamInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredImageStreamInformer constructs a new informer for ImageStream type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredImageStreamInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ImageV1().ImageStreams(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ImageV1().ImageStreams(namespace).Watch(context.TODO(), options)
			},
		},
		&imagev1.ImageStream{},
		resyncPeriod,
		indexers,
	)
}

func (f *imageStreamInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredImageStreamInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *imageStreamInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&imagev1.ImageStream{}, f.defaultInformer)
}

func (f *imageStreamInformer) Lister() v1.ImageStreamLister {
	return v1.NewImageStreamLister(f.Informer().GetIndexer())
}
