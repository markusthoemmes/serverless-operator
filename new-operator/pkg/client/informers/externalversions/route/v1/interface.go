// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/openshift-knative/serverless-operator/new-operator/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Routes returns a RouteInformer.
	Routes() RouteInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Routes returns a RouteInformer.
func (v *version) Routes() RouteInformer {
	return &routeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
